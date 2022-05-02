package v1

import (
	"codetube.cn/interface-web/v1/html"
	"github.com/gin-gonic/gin"
)

func HtmlRegister(group *gin.RouterGroup) {
	htmlGroup := ApiRouter.Group(group, "/")
	{
		//首页
		ApiRouter.Get("/", html.PageIndex(), htmlGroup)
		//课程分类页
		ApiRouter.Get("/category/:name", html.PageCategory(), htmlGroup)
		//课程分类页 - 带页码
		ApiRouter.Get("/category/:name/:page", html.PageCategory(), htmlGroup)
	}
}
