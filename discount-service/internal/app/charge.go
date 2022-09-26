package application

import (
	"github.com/gin-gonic/gin"
	"github.com/roozbehrahmani/abrarvan_test/internal/models"
	charge_code_job "github.com/roozbehrahmani/abrarvan_test/internal/queue/jobs/chargeCode"
	"github.com/sirupsen/logrus"
	"gopkg.in/validator.v2"
	"net/http"
)

func (a *Application) ChargeWallet() func(c *gin.Context) {
	return func(c *gin.Context) {
		var ChargeRequest models.DiscountServiceChargeWithCodeAndPhoneRequest

		err := c.ShouldBindJSON(&ChargeRequest)
		if err != nil {
			logrus.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := validator.Validate(ChargeRequest); err != nil {
			logrus.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, chargeCode, err := a.ChargeCodeService.Validation(ChargeRequest.Code, ChargeRequest.Phone)
		if err != nil {
			logrus.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		a.dispatcher.JobQueue <- charge_code_job.ChargeWalletJob{Name: "charge wallet", User: user, ChargeCode: chargeCode, ChargeCodeService: a.ChargeCodeService}

		c.JSON(http.StatusCreated, "received")

	}
}
