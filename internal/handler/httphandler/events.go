package httphandler

import (
	"net/http"

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
