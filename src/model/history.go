package model

import (
	"time"
)

type History struct {
	ID          int
	Date        time.Time
	Description string
	Success     bool
}
