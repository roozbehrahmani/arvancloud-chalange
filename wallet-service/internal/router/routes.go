package router

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) Routes() *gin.Engine {
	router := s.router

	router.GET("/wallet/:phone", s.app.GetWallet())
	router.POST("/wallet/charge", s.app.ChargeWallet())

	return router
}
