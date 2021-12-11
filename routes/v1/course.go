package v1

import (
	"codetube.cn/interface-web/v1/course"
	"github.com/gin-gonic/gin"
)

func CourseRegister(group *gin.RouterGroup) {
	courseGroup := ApiRouter.Group(group, "/course")
	{
		ApiRouter.Get("", course.ApiGetCourses(), courseGroup)
	}
}