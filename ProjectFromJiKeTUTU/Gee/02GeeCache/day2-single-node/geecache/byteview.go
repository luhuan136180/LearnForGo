package geecache

type ByteView struct {
	//b 将会存储真实的缓存值。选择 byte 类型是为了能够支持任意的数据类型的存储，
	//例如字符串、图片等。
	b []byte
}

func (v ByteView) Len() int {
	return len(v.b) //内建函数len返回 v 的长度，这取决于具体类型：
}

//b 是只读的，使用 ByteSlice() 方法返回一个拷贝，防止缓存值被外部程序修改。
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
