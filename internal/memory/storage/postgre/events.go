package storage

import (
	"database/sql"

	"github.com/SashaMelva/web-service-gin/internal/entity"
)

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

func (s *Storage) GetEventById(id int) (*entity.Event, error) {
	var event entity.Event
	query := `select id, title, date_time_start, date_time_end, description from events where id = $1`
	row := s.ConnectionDB.QueryRow(query, id)

	err := row.Scan(
		&event.Id,
		&event.Title,
		&event.DateTimeStart,
		&event.DateTimeEnd,
		&event.Description,
	)

	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return &event, err
	}

	return &event, nil
}

func (s *Storage) GetEvents() ([]entity.Event, error) {
	var events []entity.Event
	query := `select * from events`
	rows, err := s.ConnectionDB.QueryContext(s.Ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		event := entity.Event{}

		if err := rows.Scan(event); err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
