package models

import "time"

type LanguageModel struct {
	ID          int32 `sql:"primary_key"`
	Language    string
	CreatedBy   string
	CreatedDate time.Time
}