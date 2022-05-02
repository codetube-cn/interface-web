package template_funcs

import "html/template"

var FuncMap = template.FuncMap{
	"chunkStart": ChunkStart,
	"chunkEnd":   ChunkEnd,

	"categoryUrl": CategoryUrl,

	"courseUrl": CourseUrl,
}
