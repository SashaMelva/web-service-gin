package app

import (
	"errors"
	"fmt"
	"time"

	"github.com/SashaMelva/web-service-gin/internal/entity"
)

func (a *App) CreateEvent(event *entity.Event) (int, error) {
	id, err := a.storage.CreateEvent(event)

	if err != nil {
		a.log.Error(err)
	} else {
		a.log.Info(fmt.Sprintf("Create event whith id = %v", id))
	}

	return id, err
}

func (a *App) GetEvent(id int) (*entity.Event, error) {
	var event *entity.Event
	event, err := a.storage.GetEventById(id)

	if err != nil {
		a.log.Error(err)
		return nil, err
	}

	return event, nil
}

func (a *App) GetEvents() (*entity.EventsList, error) {
	var events *entity.EventsList

	events, err := a.storage.GetEvents()

	if err != nil {
		a.log.Error(err)
		return nil, err
	}

	return events, nil
}

func (a *App) DeleteEvent(id int) error {
	err := a.storage.DeleteEventById(id)

	if err != nil {
		a.log.Error(err)
		return err
	}

	return nil
}

func (a *App) UpdateEvent(event *entity.Event) error {
	err := a.storage.UpdateEvent(event)

	if err != nil {
		a.log.Error(err)
		return err
	}

	return nil
}

func (a *App) GetEventsSendingByPeriodConst(period string, startDate *time.Time) (*entity.EventsList, error) {
	var events *entity.EventsList
	var err error

	if a.period[period] == "" {
		return nil, errors.New("Invalid period name")
	}

	if a.period[period] == "none" {
		events, err = a.storage.GetEventsWithNotNullDateSendig()
	} else {
		events, err = a.storage.GetEventsSendingByPeriod(a.getDatesByPeriod(period, startDate))
	}

	if err != nil {
		a.log.Error(err)
		return nil, err
	}

	return events, nil
}

func (a *App) GetEventsByPeriodConst(period string, startDate *time.Time) (*entity.EventsList, error) {
	var events *entity.EventsList
	var err error

	if a.period[period] == "" {
		return nil, errors.New("Invalid period name")
	}
	if a.period[period] == "none" {
		events, err = a.storage.GetEvents()
	} else {
		events, err = a.storage.GetEventsByPeriod(a.getDatesByPeriod(period, startDate))
	}

	if err != nil {
		a.log.Error(err)
		return nil, err
	}

	return events, nil
}

func (a *App) getDatesByPeriod(period string, startDate *time.Time) (*time.Time, *time.Time) {
	l, _ := time.LoadLocation("Europe/Moscow")

	startDateTime := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, l)
	endDateTime := startDateTime.Add(24 * time.Hour)

	switch period {
	case "Week":
		switch startDateTime.Weekday() {
		case time.Monday:
			endDateTime = startDateTime.Add(7 * 24 * time.Hour)
		case time.Tuesday:
			endDateTime = startDateTime.Add(6 * 24 * time.Hour)
		case time.Wednesday:
			endDateTime = startDateTime.Add(5 * 24 * time.Hour)
		case time.Thursday:
			endDateTime = startDateTime.Add(4 * 24 * time.Hour)
		case time.Friday:
			endDateTime = startDateTime.Add(3 * 24 * time.Hour)
		case time.Saturday:
			endDateTime = startDateTime.Add(2 * 24 * time.Hour)
		case time.Sunday:
			endDateTime = startDateTime.Add(1 * 24 * time.Hour)
		}
	case "Mounth":
		endDateTime = time.Date(startDate.Year(), startDateTime.Month()+1, startDate.Day(), 0, 0, 0, 0, l)
	}

	a.log.Debug(startDateTime, endDateTime)
	return &startDateTime, &endDateTime
}
