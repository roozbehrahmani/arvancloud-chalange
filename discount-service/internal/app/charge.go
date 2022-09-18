package application

import (
	"github.com/gin-gonic/gin"
	"github.com/roozbehrahmani/abrarvan_test/internal/models"
	"github.com/sirupsen/logrus"
	"net/http"
)

type channelStruct struct {
	response map[string]interface{}
	error    error
}

func (a *Application) ChargeWallet() func(c *gin.Context) {
	return func(c *gin.Context) {

		var ChargeRequest models.DiscountServiceChargeWithCodeAndPhoneRequest

		err := c.ShouldBindJSON(&ChargeRequest)
		if err != nil {
			logrus.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		channel := make(chan channelStruct)
		go func() {
			response, err := a.ChargeCodeService.ChargeWithCodeAndPhone(ChargeRequest.Code, ChargeRequest.Phone)
			channel <- channelStruct{
				response: response,
				error:    err,
			}
		}()
		out := <-channel

		//response, err := a.ChargeCodeService.ChargeWithCodeAndPhone(ChargeRequest.Code, ChargeRequest.Phone)
		if out.error != nil {
			logrus.Error(out.error)
			c.JSON(http.StatusBadRequest, gin.H{"error": out.error.Error()})
			return
		}

		c.JSON(http.StatusCreated, out.response)

	}
}
