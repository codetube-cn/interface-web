package html

import (
	"codetube.cn/core/codes"
	"codetube.cn/core/libraries"
	"codetube.cn/core/service"
	"codetube.cn/interface-web/interfaces"
	"codetube.cn/proto/service_course"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

type Category struct {
	template    string
	title       string
	keywords    string
	description string
	data        map[string]interface{}
	*interfaces.HtmlInterfaceTrait

	request  *http.Request
	response *http.Response
}

func PageCategory() *Category {
	return &Category{
		HtmlInterfaceTrait: interfaces.NewHtmlInterfaceTrait(),
		template:           "v1/category",
		title:              "",
		keywords:           "",
		description:        "",
		data:               make(map[string]interface{}),
	}
}

func (i *Category) Handler() {
	i.WithHandler(func(c *gin.Context) {
		//分类标识，无则直接 404
		name := c.Param("name")
		p := c.Param("page")
		if name == "" {
			i.Html.AbortWithStatus(http.StatusNotFound)
		}

		//课程分类
		courseServiceClient, err := service.Client.Course()
		if err != nil {
			log.Println("课程服务出错["+strconv.Itoa(codes.ServiceConnectedFail)+"]", err)
			fmt.Println("课程服务出错["+strconv.Itoa(codes.ServiceConnectedFail)+"]", err)
			i.Html.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		//这个是公共的，要想办法抽出来
		categories, err := courseServiceClient.GetCategoryTree(c, &service_course.CategoriesTreeRequest{ParentId: 0})
		if err != nil {
			log.Println("课程服务出错["+strconv.Itoa(codes.ServiceRequestFail)+"]", err)
			fmt.Println("课程服务出错["+strconv.Itoa(codes.ServiceRequestFail)+"]", err)
			i.Html.AbortWithStatus(http.StatusInternalServerError)
		} else {
			i.data["categories"] = categories.GetData()
		}

		//判断分类标识，获取分类信息
		category := &service_course.Category{}
		if libraries.IsDigit(name) {
			id, _ := strconv.Atoi(name)
			category, err = courseServiceClient.GetCategoryById(c, &service_course.GetCategoryByIdRequest{Id: int64(id)})
		} else {
			category, err = courseServiceClient.GetCategoryByUrlName(c, &service_course.GetCategoryByUrlNameRequest{UrlName: name})
		}

		if err != nil {
			fmt.Println("错误不等于 not found", reflect.TypeOf(err))
			log.Println("课程服务出错["+strconv.Itoa(codes.ServiceRequestFail)+"]", err)
			fmt.Println("课程服务出错["+strconv.Itoa(codes.ServiceRequestFail)+"]", err)
			i.Html.AbortWithStatus(http.StatusNotFound)
			return
		} else {
			i.data["category"] = category
			i.title = category.GetName()
			i.keywords = category.GetKeywords()
			i.description = category.GetDescription()
		}

		var page int64
		page = 1
		if libraries.IsDigit(p) {
			pp, _ := strconv.Atoi(p)
			if pp > 0 {
				page = int64(pp)
			}
		}

		//课程分页列表
		courses, err := courseServiceClient.GetCoursesPagination(c, &service_course.GetCoursesRequest{
			Page:       page,
			CategoryId: category.GetId(),
		})
		if err != nil {
			log.Println("课程服务出错["+strconv.Itoa(codes.ServiceRequestFail)+"]", err)
			fmt.Println("课程服务出错["+strconv.Itoa(codes.ServiceRequestFail)+"]", err)
			i.Html.AbortWithStatus(http.StatusInternalServerError)
		} else {
			i.data["pagination"] = courses.GetPagination()
			i.data["courses"] = courses.GetItems()
		}

		//渲染模板
		i.data["meta_title"] = i.title
		i.data["meta_keywords"] = i.keywords
		i.data["meta_description"] = i.description
		c.HTML(http.StatusOK, i.template, i.data)
	})
}
