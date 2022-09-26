package application

import (
	"github.com/gin-gonic/gin"
	"github.com/roozbehrahmani/abrarvan_test/internal/app/helpers"
	"github.com/roozbehrahmani/abrarvan_test/internal/models"
	"github.com/sirupsen/logrus"
	"gopkg.in/validator.v2"
	"net/http"
)

type ChargeWalletPostRequest struct {
	Phone  string
	Amount float64
	Secret string
}

func (a *Application) GetWallet() func(c *gin.Context) {
	return func(c *gin.Context) {
		phone := c.Param("phone")
		wallet, err := a.WalletService.GetWalletWithPhone(phone)
		if err != nil {
			logrus.Error(err)
			if err.Error() == noRecordError {
				c.Status(http.StatusNotFound)
			} else {
				c.Status(http.StatusInternalServerError)
			}
			return
		}
		c.JSON(http.StatusOK, wallet)

	}
}
func (a *Application) ChargeWallet() func(c *gin.Context) {
	return func(c *gin.Context) {
		var ChargeWalletRequest models.ChargeWalletRequest
		err := c.ShouldBindJSON(&ChargeWalletRequest)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := validator.Validate(ChargeWalletRequest); err != nil {
			logrus.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if ChargeWalletRequest.Secret != a.Config.Secret {
			c.JSON(http.StatusUnauthorized, gin.H{"error": helpers.UnauthorizedRequestWalletChargingError{}})
			return
		}
		wallet, err := a.WalletService.GetWalletWithPhone(ChargeWalletRequest.Phone)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		wallet, err = a.WalletService.WalletCharging(wallet, ChargeWalletRequest.Amount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"wallet": wallet})

	}
}
