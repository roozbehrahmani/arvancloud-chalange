package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/roozbehrahmani/abrarvan_test/config"
	application "github.com/roozbehrahmani/abrarvan_test/internal/app"
	"log"
)

type Server struct {
	router *gin.Engine
	app    *application.Application
	cnf    *config.Config
}

func Initialize(ctx context.Context, cnf *config.Config, app *application.Application) *Server {
	router := gin.Default()
	return &Server{
		router: router,
		app:    app,
		cnf:    cnf,
	}
}

func (s *Server) Run() error {
	r := s.Routes()

	err := r.Run("127.0.0.1:8081")

	if err != nil {
		log.Printf("there was an error when calling routes: %v", err)
		return err
	}

	return nil
}
