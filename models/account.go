package models

type Account struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	WebSite  string `json:"webSite"`
}
