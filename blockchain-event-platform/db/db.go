package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var WriteDB *sql.DB // primary DB (for writes)
var ReadDB *sql.DB  // replica DB (for reads)

func Init() {
	var err error

	// Connect to primary DB (write)
	WriteDB, err = sql.Open("postgres", "postgres://myuser:mypassword@localhost:5432/blockchain?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect primary DB:", err)
	}

	// Test the connection
	err = WriteDB.Ping()
	if err != nil {
		log.Fatal("Cannot ping primary DB:", err)
	}

	// Connect to replica DB (read)
	ReadDB, err = sql.Open("postgres", "postgres://myuser:mypassword@localhost:5432/blockchain_read?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect replica DB:", err)
	}

	// Test the connection
	err = ReadDB.Ping()
	if err != nil {
		log.Fatal("Cannot ping replica DB:", err)
	}

	log.Println("Database connections initialized successfully")
}
