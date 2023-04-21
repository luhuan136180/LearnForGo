package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	Params map[string]string //对路由参数进行访问，将解析后的参数存储在Params中，
	// response info
	StatusCode int
}

//私有函数：初始化一个自定义context
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method, //(Get,Post,Put.....)
	}
}

//
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

//返回响应体中的表单等类型中的参数
func (c *Context) PostForm(key string) string {
	//FormValue函数是用来获取POST请求中的表单数据的，
	//取POST请求的raw的内容，需要通过解析request body来获取。
	return c.Req.FormValue(key)
}

//返回请求相应的url中携带的参数中的key对应参数值
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

//设置response的head中信息
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	//写入response
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

//配置response
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
