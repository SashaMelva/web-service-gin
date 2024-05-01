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

func (a *App) GetEvents() ([]entity.Event, error) {
	var events []entity.Event

	events, err := a.storage.GetEvents()

	if err != nil {
		a.log.Error(err)
		return nil, err
	}

	return events, nil
}
