package helpers

import (
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ListCrudResponseGen(data interface{}, count int64, limit, offset int) gin.H {
	return gin.H{
		"data": data,
		"meta": gin.H{
			"total_page": math.Ceil(float64(count) / float64(limit)),
			"next_offset": func() interface{} {
				if int64(limit+offset) < count {
					return limit + offset
				}
				return nil
			}(),
		},
	}
}

func ServerError(c *gin.Context, err error) {
	logrus.Error(err)
	c.JSON(http.StatusInternalServerError, "internal server error")
}
