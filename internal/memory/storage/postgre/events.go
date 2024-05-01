package storage

import "github.com/SashaMelva/web-service-gin/internal/entity"

func (s *Storage) CreateEvent(event *entity.Event) (int, error) {
	var eventId int
	query := `insert into events(title, description, date_time_start, date_time_end) values($1, $2, $3, $4) RETURNING id`
	result := s.ConnectionDB.QueryRow(query, event.Title, event.Description, event.DateTimeStart, event.DateTimeEnd)
	err := result.Scan(&eventId)

	if err != nil {
		return 0, err
	}

	return eventId, nil
}
