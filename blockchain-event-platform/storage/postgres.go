package storage

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

/*
Event model
Represents a decoded blockchain event
*/
type Event struct {
	BlockNumber     int64
	TxHash          string
	ContractAddress string
	EventName       string
	Payload         string
	Timestamp       time.Time
}

/*
Create DB connection
*/
func NewDB() *sql.DB {

	connStr := "user=postgres password=postgres dbname=blockchain_indexer sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("DB not reachable:", err)
	}

	log.Println("Connected to PostgreSQL")

	return db
}

/*
Insert single event
*/
func InsertEvent(db *sql.DB, e Event) error {

	query := `
	INSERT INTO events
	(block_number, tx_hash, contract_address, event_name, payload, timestamp)
	VALUES ($1,$2,$3,$4,$5,$6)
	`

	_, err := db.Exec(
		query,
		e.BlockNumber,
		e.TxHash,
		e.ContractAddress,
		e.EventName,
		e.Payload,
		e.Timestamp,
	)

	return err
}

/*
Batch insert (better performance)
*/
func InsertBatch(db *sql.DB, events []Event) error {

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`
	INSERT INTO events
	(block_number, tx_hash, contract_address, event_name, payload, timestamp)
	VALUES ($1,$2,$3,$4,$5,$6)
	`)

	if err != nil {
		return err
	}

	defer stmt.Close()

	for _, e := range events {

		_, err := stmt.Exec(
			e.BlockNumber,
			e.TxHash,
			e.ContractAddress,
			e.EventName,
			e.Payload,
			e.Timestamp,
		)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
