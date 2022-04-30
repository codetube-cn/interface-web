package template_funcs

//---------------------------------------------------------
//              所有模板公共自定义函数
//---------------------------------------------------------

// ChunkStart 切块：当前元素是否为开头
func ChunkStart(size int, key int) bool {
	return key%size == 0
}

// ChunkEnd 切块：当前元素是否为结尾
func ChunkEnd(size int, key int, len int) bool {
	return key >= len-1 || key%size == size-1
}
