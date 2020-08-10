package models

import (
	"time"
)


type UserRatingCommentsModel struct {
	ID           int32 `sql:"primary_key"`
	MovieID      int64
	UserID       int64
	Rating 		int
	Comment      *string
	CreatedBy    string
	CreatedDate  time.Time
	ModifiedBy   *string
	ModifiedDate *time.Time
}

type UserRatingCommentsRespModel struct {
	Rating      int
	UserID       int64
	Comment    string
	MovieID      int64
}