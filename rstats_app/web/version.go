package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"dev.maizy.ru/rstats/rstats_app"
)

func BuildVersionHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"version": rstats_app.GetVersion(),
		})
	}
}
