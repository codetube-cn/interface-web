package html

import (
	"codetube.cn/core/codes"
	"codetube.cn/core/service"
	"codetube.cn/interface-web/interfaces"
	"codetube.cn/proto/service_course"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Index struct {
	template    string
	title       string
	keywords    string
	description string
	data        map[string]interface{}
	*interfaces.HtmlInterfaceTrait

	request  *http.Request
	response *http.Response
}

func PageIndex() *Index {
	return &Index{
		HtmlInterfaceTrait: interfaces.NewHtmlInterfaceTrait(),
		template:           "v1/index",
		title:              "CodeTube 在线教育官方网站",
		keywords:           "",
		description:        "",
		data:               make(map[string]interface{}),
	}
}

func (i *Index) Handler() {
	i.WithHandler(func(c *gin.Context) {
		i.request = c.Request
		userid := i.request.Header.Get("CodeTube-User-ID")
		i.data["userid"] = userid
		//课程分类
		courseServiceClient, err := service.Client.Course()
		if err != nil {
			log.Println("课程服务出错["+strconv.Itoa(codes.ServiceConnectedFail)+"]", err)
			i.Html.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		categories, err := courseServiceClient.GetCategoryTree(c, &service_course.CategoriesTreeRequest{ParentId: 0})
		if err != nil {
			log.Println("课程服务出错["+strconv.Itoa(codes.ServiceRequestFail)+"]", err)
			i.Html.AbortWithStatus(http.StatusInternalServerError)
		} else {
			i.data["categories"] = categories.GetData()
		}
		//精选推荐课程
		recommendedCourses, err := courseServiceClient.GetCourses(c, &service_course.GetCoursesRequest{
			Page:              1,
			PageSize:          12,
			IsRecommended:     true,
			IsRecommendedFill: false,
			CategoryId:        0,
		})
		if err != nil {
			log.Println("课程服务出错["+strconv.Itoa(codes.ServiceRequestFail)+"]", err)
			i.Html.AbortWithStatus(http.StatusInternalServerError)
		} else {
			i.data["recommended_courses"] = recommendedCourses.GetCourses()
		}

		//渲染模板
		i.data["meta_title"] = i.title
		i.data["meta_keywords"] = i.keywords
		i.data["meta_description"] = i.description
		c.HTML(http.StatusOK, i.template, i.data)
	})
}
