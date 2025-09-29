package entity

import (
	"time"
)

type Log struct {
	ID        string
	CreatedAt time.Time
	Level     string
	Message   any
}
