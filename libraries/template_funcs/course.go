package template_funcs

import (
	"strconv"
)

//---------------------------------------------------------
//              所有模板公共自定义函数
//---------------------------------------------------------

// CourseUrl 获取课程 URL
func CourseUrl(urlName string, id int64) string {
	if urlName != "" {
		return "/course/" + urlName
	} else {
		return "/course/" + strconv.FormatInt(id, 10)
	}
}
