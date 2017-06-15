package models

import (
	"time"
)

type Account struct {
	ID          int    `storm:"id,increment"`
	UserName    string `storm:"index"`
	Email       string
	Password    string    `json:"password,omitempty"`
	WebSite     string    `json:"webSite"`
	CreatedAt   time.Time `storm:"index"`
	CurrentUser int
	Users       []int
}
