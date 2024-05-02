package httphandler

import (
	"net/http"
	"time"

	"github.com/SashaMelva/web-service-gin/internal/entity"
	"github.com/gin-gonic/gin"
)

func (s *Service) GetEventsBySendingHandler(ctx *gin.Context) {
	var events *entity.EventsList
	var err error
	var startDate time.Time

	period := ctx.Params.ByName("period")
	if period == "" {
		s.log.Debug("get all events")
		events, err = s.app.GetEvents()

	} else {
		startDate = time.Now()
		s.log.Debug("get events for sending by period ", period)
		s.log.Debug("StartDate ", startDate)
		events, err = s.app.GetEventsSendingByPeriodConst(period, &startDate)
	}

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}
	s.log.Debug(events)
	ctx.JSON(http.StatusOK, events)
}
