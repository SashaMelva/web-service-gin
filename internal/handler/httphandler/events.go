package httphandler

import (
	"net/http"
	"strconv"

	"github.com/SashaMelva/web-service-gin/internal/entity"
	"github.com/gin-gonic/gin"
)

func (s *Service) CreateEventHandler(ctx *gin.Context) {
	var event entity.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	s.log.Debug(event)

	id, err := s.app.CreateEvent(&event)

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Id": id})
}

func (s *Service) GetAllEventsHandler(ctx *gin.Context) {
	var events *entity.EventsList
	var err error

	// period := ctx.Query("period")
	// if period == "" {

	s.log.Debug("events")
	events, err = s.app.GetEvents()

	s.log.Debug("events2")
	// } else {
	// 	// split := strings.Split(period, ":")
	// 	// if len(split) == 2 {
	// 	// 	events, err = s.app.GetEventsByPeriod(period)
	// 	// } else {

	// 	// }
	// 	s.log.Debug("Period ", period)
	// 	events, err = s.app.GetEventsByPeriodConst(period)
	// }

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}
	s.log.Debug(events)
	ctx.JSON(http.StatusOK, events)
}

func (s *Service) GetEventByIdHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	event, err := s.app.GetEvent(id)
	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func (s *Service) DeleteEventHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	err = s.app.DeleteEvent(id)
	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Id": id})
}

func (s *Service) UpdateEventHandler(ctx *gin.Context) {
	var event entity.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	s.log.Debug(event)

	err := s.app.UpdateEvent(&event)

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
