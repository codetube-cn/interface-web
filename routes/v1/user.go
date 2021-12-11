package v1

import (
	"codetube.cn/interface-web/v1/user"
	"github.com/gin-gonic/gin"
)

func UserRegister(group *gin.RouterGroup) {
	userGroup := ApiRouter.Group(group, "/user")
	{
		ApiRouter.Get("", user.ApiGetUser(), userGroup)
	}
}
