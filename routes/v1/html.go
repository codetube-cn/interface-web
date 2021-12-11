package v1

import (
	"codetube.cn/interface-web/v1/html"
	"github.com/gin-gonic/gin"
)

func HtmlRegister(group *gin.RouterGroup) {
	htmlGroup := ApiRouter.Group(group, "/html")
	{
		ApiRouter.Get("/", html.PageIndex(), htmlGroup)
	}
}