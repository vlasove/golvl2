package models

//Event ...
type Event struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Date   string `json:"date"`
	Info   string `json:"info"`
}
