package viewModels

import "time"


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
	Rating      int32
	UserID       int64
	Title    string
	Comment    string
	MovieID      int64
}