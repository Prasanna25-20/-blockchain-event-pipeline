package main

import (
	"blockchain-event-platform/db" // import the db package
	"log"
)

func main() {
	db.Init() // initialize DB connections

	// Example write
	_, err := db.WriteDB.Exec("INSERT INTO events(contract_address, block_number, data) VALUES($1,$2,$3)",
		"0x123", 101, "some_event")
	if err != nil {
		log.Println("Write error:", err)
	}

	// Example read
	rows, err := db.ReadDB.Query("SELECT * FROM events WHERE contract_address=$1", "0x123")
	if err != nil {
		log.Println("Read error:", err)
	}
	defer rows.Close()
}
