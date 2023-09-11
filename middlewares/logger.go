package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc{
	return gin.LoggerWithFormatter(func (param gin.LogFormatterParams) string  {
		return fmt.Sprintf("%s %d %s",
			param.ClientIP,
			param.StatusCode,
			param.TimeStamp,
		)
	})
}