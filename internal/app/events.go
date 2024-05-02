package app

import (
	"fmt"

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

func (a *App) GetEventsByPeriodConst(period string) (*entity.EventsList, error) {
	var events *entity.EventsList
	var err error

	a.log.Debug(period, entity.Period(period))
	// if entity.Period(period) != "" {
	// 	events, err = a.storage.GetEventsByPeriod(entity.Period(period))
	// }

	if err != nil {
		a.log.Error(err)
		return nil, err
	}

	return events, nil
}
