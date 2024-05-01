package app

import (
	"fmt"

	"github.com/SashaMelva/web-service-gin/internal/entity"
)

func (a *App) CreateEvent(event *entity.Event) int {
	id, err := a.storage.CreateEvent(event)

	if err != nil {
		a.log.Error(err)
	} else {
		a.log.Info(fmt.Sprintf("Create event whith id = %v", id))
	}

	return id
}
