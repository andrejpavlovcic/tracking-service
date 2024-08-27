package db

import "time"

type AccountEvent struct {
	AccountID int64     `json:"account_id"`
	Data      string    `json:"data"`
	Timestamp time.Time `json:"timestamp"`
}
