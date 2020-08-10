package models

import "time"


type UserModel struct {
	ID           int32 `sql:"primary_key"`
	Name         *string
	LoginID      string
	Password     string
	Role         string
	CreatedBy    string
	CreatedDate  time.Time
	ModifiedBy   *string
	ModifiedDate time.Time
}