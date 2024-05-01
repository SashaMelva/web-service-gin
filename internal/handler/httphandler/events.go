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
	}

	id := s.app.CreateEvent(&event)
	ctx.JSON(http.StatusOK, gin.H{"Id": id})
}

func (s *Service) GetAllEventsHandler(ctx *gin.Context) {
	var events []entity.Event
	events, err := s.app.GetEvents()

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

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
