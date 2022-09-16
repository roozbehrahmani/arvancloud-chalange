package application

import (
	"github.com/gin-gonic/gin"
	"github.com/roozbehrahmani/abrarvan_test/internal/models"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (a *Application) ChargeWallet() func(c *gin.Context) {
	return func(c *gin.Context) {

		var ChargeRequest models.ChargeCodeRequest

		err := c.ShouldBindJSON(&ChargeRequest)
		if err != nil {
			logrus.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response, err := a.ChargeCodeService.ChargeWithCodeAndPhone(ChargeRequest.Code, ChargeRequest.Phone)
		if err != nil {
			logrus.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, response)

	}
}
