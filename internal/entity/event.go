package entity

import "time"

type Event struct {
	Id            int        `json:"id,omitempty" db:"id"`
	Title         string     `json:"title" db:"title" binding:"required"`
	DateTimeStart *time.Time `json:"date_time_start" db:"date_time_start" binding:"required"`
	DateTimeEnd   *time.Time `json:"date_time_end" db:"date_time_end"`
	Description   string     `json:"description" db:"description"`
	DataTimeSend  *time.Time `json:"date_time_send" db:"date_time_send"`
}

type EventsList struct {
	Events []*Event `json:"events"`
}

type Period map[string]string
