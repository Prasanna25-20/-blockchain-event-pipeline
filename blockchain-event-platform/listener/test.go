package main

import (
	"blockchain-event-platform/storage"
	"fmt"
	"time"
)

func main() {

	// connect database
	db := storage.NewDB()

	fmt.Println("Testing Event Listener...")

	// simulate blockchain event
	event := storage.Event{
		BlockNumber:     999999,
		TxHash:          "0xtesthash123",
		ContractAddress: "0xtestcontract",
		EventName:       "Transfer",
		Payload:         `{"from":"0xabc","to":"0xdef","value":"500"}`,
		Timestamp:       time.Now(),
	}

	// insert event
	err := storage.InsertEvent(db, event)

	if err != nil {
		fmt.Println("Insert failed:", err)
		return
	}

	fmt.Println("Test event inserted successfully")
}
