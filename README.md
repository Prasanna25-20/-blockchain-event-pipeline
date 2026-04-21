#  Scalable Blockchain Event Processing Platform

A production-style full-stack system that ingests blockchain events in real time using Kafka, processes them efficiently, stores them in PostgreSQL, and visualizes them through a responsive UI dashboard with full observability using Prometheus and Grafana.

---

##  Overview

This project demonstrates a **real-world event-driven architecture** used in scalable backend systems such as blockchain indexing, analytics pipelines, and streaming platforms.

It includes:

* Real-time event ingestion (Kafka)
* Asynchronous processing (Go services)
* Persistent storage (PostgreSQL)
* API layer (REST)
* Monitoring (Prometheus + Grafana)
* Frontend dashboard for visualization

---

##  Architecture

```
Producer → Kafka → Processor → PostgreSQL → API → Frontend (React)
                               ↓
                        Prometheus → Grafana
```

---

##  Tech Stack

| Layer      | Technology |
| ---------- | ---------- |
| Language   | Go         |
| Messaging  | Kafka      |
| Database   | PostgreSQL |
| API        | net/http   |
| Frontend   | React.js   |
| Monitoring | Prometheus |
| Dashboard  | Grafana    |
| Container  | Docker     |

---

##  Key Features

*  Real-time blockchain event ingestion using Kafka
*  Asynchronous event processing pipeline
*  Reliable data storage with PostgreSQL
*  High-performance REST API
*  Live metrics collection using Prometheus
*  Visualization with Grafana dashboards
*  Interactive frontend dashboard for event data
*  Fully containerized system using Docker

---

##  Frontend Dashboard

A responsive UI to visualize and monitor blockchain event processing.

### Features

*  Displays processed events from API
*  Real-time or near real-time updates
*  Clean dashboard interface
*  Easy data inspection

### Tech Stack

* React.js
* Tailwind CSS (if used)
* Fetch / Axios

---

### Screenshots

![Dashboard](./assets/dashboard.png)
![Events](./assets/events.png)

---

## Getting Started

### 1. Clone Repository

```bash
git clone https://github.com/Prasanna25-20/scalable-blockchain-event-pipeline-.git
cd scalable-blockchain-event-pipeline-
```

---

### 2. Start All Services

```bash
docker-compose up --build
```

---

### 3. Create Kafka Topic

```bash
docker exec -it project2-kafka-1 bash

kafka-topics.sh --create \
  --topic events \
  --bootstrap-server kafka:9092 \
  --partitions 1 \
  --replication-factor 1
```

---

### 4. Send Test Events

```bash
kafka-console-producer.sh \
  --broker-list kafka:9092 \
  --topic events
```

---

### 5. Access API

```bash
curl http://localhost:8080/events
```

---

### 6. Run Frontend

```bash
cd frontend
npm install
npm start
```

---

##  Monitoring & Observability

| Tool       | URL                   |
| ---------- | --------------------- |
| Prometheus | http://localhost:9090 |
| Grafana    | http://localhost:3000 |

### Example Metrics

* API request count
* API latency
* Event processing rate
* System performance

---

## Project Structure

```
project/
├── api/                # REST API service
├── processor/          # Kafka event processor
├── frontend/           # React UI dashboard
├── assets/             # Screenshots & diagrams
├── docker-compose.yml
├── prometheus.yml
├── init.sql
└── README.md
```

---

##  Highlights

* Event-driven architecture using Kafka
* Full-stack implementation (Backend + UI)
* Scalable and modular system design
* Observability-first engineering approach
* Production-style development setup

---

##  Future Improvements

* Kubernetes deployment
* Load balancing (NGINX)
* Authentication & rate limiting
* Distributed tracing (OpenTelemetry)
* WebSocket-based real-time UI updates

---

##  Author

**Prasanna**
GitHub: https://github.com/Prasanna25-20

---

##  Support

If you found this project useful, consider giving it a star ⭐
