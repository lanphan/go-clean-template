package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ironsail/whydah-go-clean-template/pkg/logger"
)

type response struct {
	Error string `json:"error" example:"message"`
}

func errorResponse(c *gin.Context, code int, msg string) {
	logger.Error(fmt.Sprintf("error response: %d - %s", code, msg))
	c.AbortWithStatusJSON(code, response{msg})
}
