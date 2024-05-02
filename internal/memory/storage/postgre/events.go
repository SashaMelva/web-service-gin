package storage

import (
	"database/sql"
	"time"

	"github.com/SashaMelva/web-service-gin/internal/entity"
)

func (s *Storage) CreateEvent(event *entity.Event) (int, error) {
	var eventId int
	query := `insert into events(title, description, date_time_start, date_time_end, date_time_send) values($1, $2, $3, $4, $5) RETURNING id`
	result := s.ConnectionDB.QueryRow(query, event.Title, event.Description, event.DateTimeStart, event.DateTimeEnd, event.DataTimeSend)
	err := result.Scan(&eventId)

	if err != nil {
		return 0, err
	}

	return eventId, nil
}

func (s *Storage) GetEventById(id int) (*entity.Event, error) {
	var event entity.Event
	query := `select id, title, date_time_start, date_time_end, description, date_time_send from events where id = $1`
	row := s.ConnectionDB.QueryRow(query, id)

	err := row.Scan(
		&event.Id,
		&event.Title,
		&event.DateTimeStart,
		&event.DateTimeEnd,
		&event.Description,
		&event.DataTimeSend,
	)

	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return &event, err
	}

	return &event, nil
}

func (s *Storage) GetEvents() (*entity.EventsList, error) {
	var events entity.EventsList
	query := `select id, title, date_time_start, date_time_end, description, date_time_send from events`
	rows, err := s.ConnectionDB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		event := entity.Event{}

		if err := rows.Scan(
			&event.Id,
			&event.Title,
			&event.DateTimeStart,
			&event.DateTimeEnd,
			&event.Description,
			&event.DataTimeSend,
		); err != nil {
			return nil, err
		}

		events.Events = append(events.Events, &event)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &events, nil
}

func (s *Storage) DeleteEventById(id int) error {
	query := `delete from events where id = $1`
	_, err := s.ConnectionDB.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) UpdateEvent(event *entity.Event) error {
	query := `update events set title=$1, description=$2, date_time_start=$3, date_time_end=$4, date_time_send=$5 where id=$6`
	_, err := s.ConnectionDB.Exec(query, event.Title, event.Description, event.DateTimeStart, event.DateTimeEnd, event.DataTimeSend, event.Id)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetEventsByPeriod(dateStart, dateEnd *time.Time) (*entity.EventsList, error) {
	var events entity.EventsList
	query := `select id, title, date_time_start, date_time_end, description, date_time_send from events where date_time_start >= $1::timestamp and date_time_end < $2::timestamp`
	rows, err := s.ConnectionDB.Query(query, dateStart, dateEnd)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		event := entity.Event{}

		if err := rows.Scan(
			&event.Id,
			&event.Title,
			&event.DateTimeStart,
			&event.DateTimeEnd,
			&event.Description,
			&event.DataTimeSend); err != nil {
			return nil, err
		}

		events.Events = append(events.Events, &event)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &events, nil
}

func (s *Storage) GetEventsSendingByPeriod(dateStart, dateEnd *time.Time) (*entity.EventsList, error) {
	var events entity.EventsList
	query := `select id, title, date_time_start, date_time_end, description, date_time_send from events where date_time_send >= $1::timestamp and date_time_send < $2::timestamp`
	rows, err := s.ConnectionDB.Query(query, dateStart.Format("2006-01-02 15:04:05"), dateEnd.Format("2006-01-02 15:04:05"))

	s.log.Debug(query, dateStart.Format("2006-01-02 15:04:05"), dateEnd.Format("2006-01-02 15:04:05"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		event := entity.Event{}

		if err := rows.Scan(
			&event.Id,
			&event.Title,
			&event.DateTimeStart,
			&event.DateTimeEnd,
			&event.Description,
			&event.DataTimeSend); err != nil {
			return nil, err
		}

		events.Events = append(events.Events, &event)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &events, nil
}

func (s *Storage) GetEventsWithNotNullDateSendig() (*entity.EventsList, error) {
	var events entity.EventsList
	query := `select id, title, date_time_start, date_time_end, description, date_time_send from events where date_time_send is not null`
	rows, err := s.ConnectionDB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		event := entity.Event{}

		if err := rows.Scan(
			&event.Id,
			&event.Title,
			&event.DateTimeStart,
			&event.DateTimeEnd,
			&event.Description,
			&event.DataTimeSend,
		); err != nil {
			return nil, err
		}

		events.Events = append(events.Events, &event)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &events, nil
}
