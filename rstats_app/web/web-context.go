package web

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"

	"dev.maizy.ru/rstats/rstats_app"
)

var md5Version string

func WithCommonVars(c *gin.Context, vars gin.H) gin.H {
	version := rstats_app.GetVersion()
	if md5Version == "" {
		hash := md5.New()
		_, _ = io.WriteString(hash, version)
		md5Version = fmt.Sprintf("%x", hash.Sum(nil))
	}
	result := gin.H{
		"lang":       "en",
		"version":    version,
		"staticHash": md5Version,
	}

	for key, value := range vars {
		result[key] = value
	}
	return result
}
