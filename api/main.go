package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Event struct {
	ID   int    `json:"id"`
	Data string `json:"data"`
}

var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_requests_total",
			Help: "Total API Requests",
		},
		[]string{"method", "endpoint"},
	)

	requestLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "api_latency_seconds",
			Help: "API Latency",
		},
		[]string{"endpoint"},
	)
)

func initLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func initMetrics() {
	prometheus.MustRegister(requestCount, requestLatency)
}

func main() {

	initLogger()
	initMetrics()

	log.Info(" API starting...")

	// 🔹 DB CONNECT WITH RETRY
	var db *sql.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres",
			"postgres://myuser:mypass@postgres:5432/mydb?sslmode=disable")

		if err == nil && db.Ping() == nil {
			break
		}

		log.Warn(" Waiting for DB...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.WithError(err).Fatal(" DB not reachable")
	}
	defer db.Close()

	log.Info(" Connected to Postgres")

	mux := http.NewServeMux()

	//  HEALTH CHECK
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	//  METRICS
	mux.Handle("/metrics", promhttp.Handler())

	//  EVENTS API WITH PAGINATION
	mux.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

		if page <= 0 {
			page = 1
		}
		if limit <= 0 {
			limit = 10
		}

		offset := (page - 1) * limit

		rows, err := db.Query(
			"SELECT id, data FROM events ORDER BY id LIMIT $1 OFFSET $2",
			limit, offset,
		)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer rows.Close()

		events := []Event{}

		for rows.Next() {
			var e Event
			if err := rows.Scan(&e.ID, &e.Data); err != nil {
				continue
			}
			events = append(events, e)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(events)

		requestCount.WithLabelValues(r.Method, "/events").Inc()
		requestLatency.WithLabelValues("/events").
			Observe(time.Since(start).Seconds())

		log.WithFields(log.Fields{
			"endpoint": "/events",
			"page":     page,
			"limit":    limit,
		}).Info(" Request handled")
	})

	log.Info(" API running on :8080")

	http.ListenAndServe(":8080", mux)
}
