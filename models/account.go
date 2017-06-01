package models

import (
	"time"
)

type Account struct {
	ID        int       `storm:"id,increment"`
	UserName  string    `storm:"index,unique"`
	Email     string    `storm:"unique"`
	Password  string    `json:"password,omitempty"`
	WebSite   string    `json:"webSite"`
	CreatedAt time.Time `storm:"index"`
}
