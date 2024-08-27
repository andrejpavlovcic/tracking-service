package entities

import "time"

type Event struct {
	AccountID int64     `json:"account_id"`
	Data      string    `json:"data"`
	Timestamp time.Time `json:"timestamp"`
}
