package router

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) Routes() *gin.Engine {
	router := s.router

	router.POST("/charge", s.app.ChargeWallet())

	return router
}
