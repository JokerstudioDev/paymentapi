package model

import (
	"time"
)

type History struct {
	ID            int
	AccountNumber string
	Date          time.Time
	Description   string
	Success       bool
}
