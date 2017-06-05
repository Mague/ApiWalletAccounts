package models

import (
	"time"
)

type User struct {
	ID        int64     `storm:"id,increment"`
	UserName  string    `storm:"index,unique"`
	Email     string    `storm:"unique"`
	Password  string    `json:"Password,omitempty"`
	CreatedAt time.Time `storm:"index"`
}
