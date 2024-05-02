package http

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/SashaMelva/web-service-gin/internal/app"
	"github.com/SashaMelva/web-service-gin/internal/config"
	"github.com/SashaMelva/web-service-gin/internal/handler/httphandler"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	srv *http.Server
	log *zap.SugaredLogger
}

func NewServer(log *zap.SugaredLogger, app *app.App, config *config.ConfigHttpServer) *Server {
	log.Debug("URL: " + config.Host + ":" + config.Port)
	router := gin.Default()
	handler := httphandler.NewHendler(log, app)

	router.GET("/", func(ctx *gin.Context) {
		fmt.Println("Hellow world)")
	})

	router.POST("/event/", handler.CreateEventHandler)
	router.GET("/events/:period", handler.GetAllEventsHandler)
	router.GET("/event/:id", handler.GetEventByIdHandler)
	router.DELETE("/event/:id", handler.DeleteEventHandler)
	router.PUT("/event/", handler.UpdateEventHandler)

	router.GET("/events-sending/:period", handler.GetEventsBySendingHandler)

	return &Server{
		srv: &http.Server{
			Addr:    config.Host + ":" + config.Port,
			Handler: router,
		},
		log: log,
	}
}

func (s *Server) Start(ctx context.Context) {
	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.log.Fatalf("listen: %s\n", err)
	}
}

func (s *Server) Stop(ctx context.Context) {
	if err := s.srv.Shutdown(ctx); err != nil {
		s.log.Fatal("Server forced to shutdown: ", err)
	}

	os.Exit(1)
}
