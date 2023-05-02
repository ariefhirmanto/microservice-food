package models

import "time"

type Restaurant struct {
	ID        int64
	Name      string
	Location  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
