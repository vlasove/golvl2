package models

import (
	"errors"
	"time"
)

var (
	errInvalidUserID = errors.New("user_id should be positive integer")
	errInvalidDate   = errors.New("date should be in YYYY-MM-DD format")
	errInvalidInfo   = errors.New("info field is required and should has min 3 symbols")
)

// EventRequest ...
type EventRequest struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Date   string `json:"date"`
	Info   string `json:"info"`
}

//Validate ...
func (e *EventRequest) Validate() error {
	if e.UserID <= 0 {
		return errInvalidUserID
	}

	if _, err := time.Parse("2006-01-02", e.Date); err != nil {
		return errInvalidDate
	}
	if len(e.Info) == 0 {
		return errInvalidInfo
	}
	return nil
}

// NewEventFromRequest ...
func NewEventFromRequest(e *EventRequest) *Event {
	return &Event{
		ID:     e.ID,
		UserID: e.UserID,
		Date:   e.Date,
		Info:   e.Info,
	}
}
