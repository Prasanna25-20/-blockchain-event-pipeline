package handlers

import (
	"encoding/json"

	"blockchain-event-platform/cache"
	"blockchain-event-platform/database"
	"blockchain-event-platform/models"

	"github.com/gofiber/fiber/v2"
)

func GetEvents(c *fiber.Ctx) error {

	contract := c.Query("contract")
	event := c.Query("event")

	cacheKey := "events:" + contract + ":" + event

	cached, err := cache.RDB.Get(cache.Ctx, cacheKey).Result()

	if err == nil {
		var events []models.Event
		json.Unmarshal([]byte(cached), &events)
		return c.JSON(events)
	}

	rows, err := database.DB.Query(
		`SELECT block_number, tx_hash, contract_address, event_name, data, created_at
		 FROM events
		 WHERE contract_address=$1 AND event_name=$2`,
		contract, event,
	)

	if err != nil {
		return err
	}

	defer rows.Close()

	events := []models.Event{}

	for rows.Next() {
		var e models.Event

		err := rows.Scan(
			&e.BlockNumber,
			&e.TxHash,
			&e.Contract,
			&e.EventName,
			&e.Payload,
			&e.Timestamp,
		)

		if err != nil {
			return err
		}

		events = append(events, e)
	}

	data, _ := json.Marshal(events)
	cache.RDB.Set(cache.Ctx, cacheKey, data, 0)

	return c.JSON(events)
}

func GetEventByTx(c *fiber.Ctx) error {

	txHash := c.Params("txHash")

	row := database.DB.QueryRow(
		`SELECT block_number, tx_hash, contract_address, event_name, data, created_at
		 FROM events
		 WHERE tx_hash=$1`,
		txHash,
	)

	var e models.Event

	err := row.Scan(
		&e.BlockNumber,
		&e.TxHash,
		&e.Contract,
		&e.EventName,
		&e.Payload,
		&e.Timestamp,
	)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "event not found",
		})
	}

	return c.JSON(e)
}
