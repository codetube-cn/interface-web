package template_funcs

import (
	"strconv"
)

//---------------------------------------------------------
//              所有模板公共自定义函数
//---------------------------------------------------------

// CategoryUrl 获取课程分类 URL
func CategoryUrl(id int64, urlName string, page int64) string {
	url := ""
	if urlName != "" {
		url = "/category/" + urlName
	} else {
		url = "/category/" + strconv.FormatInt(id, 10)
	}

	if page > 1 {
		url = url + "/" + strconv.FormatInt(page, 10)
	}

	return url
}

// CourseUrl 获取课程 URL
func CourseUrl(urlName string, id int64) string {
	if urlName != "" {
		return "/course/" + urlName
	} else {
		return "/course/" + strconv.FormatInt(id, 10)
	}
}
