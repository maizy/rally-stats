package web

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"

	"github.com/gin-gonic/gin"

	"dev.maizy.ru/rstats/internal/u"
	"dev.maizy.ru/rstats/rstats_app"
)

type NavBarItem struct {
	Text     string
	Link     string
	Active   bool
	WithLink bool
}

var baseNavBarItems = []NavBarItem{
	{"Stats by day", "/by-day", false, true},
	{"Stats by track", "/by-track", false, true},
}

func BuildNavBar(currentPagePath string) []NavBarItem {
	navBarItems := u.CopySlice(baseNavBarItems)
	for i := range navBarItems {
		item := &navBarItems[i]
		if strings.HasPrefix(currentPagePath, item.Link) {
			item.Active = true
		}
		if currentPagePath == item.Link {
			item.WithLink = false
		}
	}
	return navBarItems
}

var md5Version string

func WithCommonVars(c *gin.Context, vars gin.H) gin.H {
	requestPath := c.Request.URL.Path
	navBarItems := BuildNavBar(requestPath)

	version := rstats_app.GetVersion()
	if md5Version == "" {
		hash := md5.New()
		_, _ = io.WriteString(hash, version)
		md5Version = fmt.Sprintf("%x", hash.Sum(nil))
	}
	result := gin.H{
		"lang":        "en",
		"version":     version,
		"staticHash":  md5Version,
		"navBarItems": navBarItems,
	}

	for key, value := range vars {
		result[key] = value
	}
	return result
}
