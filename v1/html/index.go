package html

import (
	"codetube.cn/interface-web/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
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
		title:              "",
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
		//渲染模板
		c.HTML(http.StatusOK, i.template, i.data)
	})
}
