package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
#代码最开头，给map[string]interface{}起了一个别名gee.H，构建JSON数据时，显得更简洁。
#Context目前只包含了http.ResponseWriter和*http.Request，另外提供了对 Method 和 Path 这两个常用属性的直接访问。
#提供了访问Query和PostForm参数的方法。
#提供了快速构造String/Data/JSON/HTML响应的方法。
*/

//这个就是gee.H   ,类比于gin.H
type H map[string]interface{}

type Context struct {
	//origin objects :来源对象
	Writer http.ResponseWriter
	Req    *http.Request

	//	request info :请求信息
	Path   string
	Method string //(Get,Post,Put.....)

	//response info :相应信息
	StatusCode int
}

//私有函数：初始化一个自定义context
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,

		//从req中直接读取，便于路由函数中直接查询
		Path:   req.URL.Path,
		Method: req.Method, //(Get,Post,Put.....)
	}
}

//返回响应体中的表单等类型中的参数
func (c *Context) PostForm(key string) string {
	//FormValue函数是用来获取POST请求中的表单数据的，
	//取POST请求的raw的内容，需要通过解析request body来获取。
	return c.Req.FormValue(key) //从请求中读取表单数据
}

//返回请求相应的url中携带的参数中的key对应参数值
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key) //从请求的url中读取url数据（?xxx=iii&yyy=aaa）
}

//一下为对相应的操作

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

//设置response的head中信息
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

//以string形式返回相应数据
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	//写入response
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

//以json格式返回相应
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	// Encode将传入参数的JSON编码写入流，
	//后面跟着换行符。
	err := encoder.Encode(obj)
	if err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}

}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content_Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
