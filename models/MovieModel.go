package models

import "time"



type MovieModel struct {
	ID          int64 `sql:"primary_key"`
	Title       string
	Description string
	Rating int32
	CreatedBy   string
	CreatedDate time.Time
}