package entity

import (
	"time"
)

type Note struct {
	ID        string
	CreatedAt time.Time
	Message   any
}
