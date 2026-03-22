package models

import "time"

type Event struct {
	BlockNumber int64     `json:"block_number" db:"block_number"`
	TxHash      string    `json:"tx_hash" db:"tx_hash"`
	Contract    string    `json:"contract" db:"contract"`
	EventName   string    `json:"event_name" db:"event_name"`
	Payload     string    `json:"payload" db:"payload"`
	Timestamp   time.Time `json:"timestamp" db:"timestamp"`
}
