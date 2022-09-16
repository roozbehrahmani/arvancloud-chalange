package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func CheckHasErrorBadRequest(c *gin.Context, err error) {

	if err != nil {
		logrus.Error(err)
		c.AbortWithError(http.StatusOK, err)
		return
	}
}
