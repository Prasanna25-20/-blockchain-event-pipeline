# 🚀 Scalable Blockchain Event Processing Platform

Production-style **event-driven system** that ingests blockchain events using Kafka, processes them with Go, stores in PostgreSQL, and visualizes via a React dashboard with Prometheus & Grafana monitoring.

---

## 💡 Overview
- Real-time event ingestion (Kafka)
- Concurrent processing using Go workers
- Reliable storage (PostgreSQL)
- REST API for data access
- Monitoring with Prometheus + Grafana
- Frontend dashboard (React)

---

## 🧠 Architecture
Producer → Kafka → Go Workers → PostgreSQL → API → React UI
↓
Prometheus → Grafana


---

## ⚙️ Tech Stack
- Go, Kafka, PostgreSQL  
- REST (net/http)  
- React.js  
- Prometheus, Grafana  
- Docker  

---

## 🚀 Key Features
- Event-driven architecture  
- Concurrent processing with Go routines  
- Fault-tolerant & scalable design  
- Real-time monitoring & metrics  
- Full-stack implementation  

---

## 📊 Performance
- ~500–1000 events/sec (simulated)  
- < 50ms API latency  
- Supports horizontal scaling  

---

## ▶️ Run Locally
```bash
git clone https://github.com/Prasanna25-20/scalable-blockchain-event-pipeline-.git
cd scalable-blockchain-event-pipeline-
docker-compose up --build

Access:

API → http://localhost:8080/events
Prometheus → http://localhost:9090
Grafana → http://localhost:3000

Use Cases
Blockchain indexing
Real-time analytics pipelines
Event-driven backend systems
👤 Author

Prasanna
https://github.com/Prasanna25-20
