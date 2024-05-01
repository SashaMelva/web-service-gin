package entity

import "time"

type Event struct {
	ID            int        `json:"id" db:"id"`
	Title         string     `json:"title" db:"title"`
	DateTimeStart *time.Time `json:"date_time_start" db:"date_time_start"`
	DateTimeEnd   *time.Time `json:"date_time_end" db:"date_time_end"`
	Description   string     `json:"description" db:"description"`
	DataTimeSend  *time.Time `json:"date_time_send" db:"date_time_send"`
}
