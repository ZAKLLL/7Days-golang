package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//封装上下文

type H map[string]interface{}

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	Params map[string]string
	// response info
	StatusCode int

	// middleware
	handlers []HandlerFunc
	index    int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index:  -1,
	}
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	//不是所有的handler都会调用 Next()。
	//手工调用 Next()，一般用于在请求前后各实现一些行为。如果中间件只作用于请求前，可以省略调用Next()，此种写法可以兼容 不调用Next的写法
	//并且使用c.index<s;c.index++ 也能保证中间件只执行一次
	//当中间件不调用 next函数时,通过此循环保证中间件执行顺序
	//当中间件调用next 函数时,且当中间件执行完毕,对应c.index也已经到达指定index
	//前面存在的for循环因为是通过c.index<s 做循环判断,则不会重复执行已经执行过的中间件
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Fail(statusCode int, msg string) {
	c.Status(statusCode)
	c.Writer.Write([]byte(msg))
}
