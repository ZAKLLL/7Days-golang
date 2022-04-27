package main

import (
	"fmt"
	"gee"
	"net/http"
)

func m1() gee.HandlerFunc {
	return func(context *gee.Context) {
		fmt.Println("m1--------")
	}
}
func m2() gee.HandlerFunc {
	return func(context *gee.Context) {
		fmt.Println("m2--------start")
		context.Next()
		fmt.Println("m2--------end")
	}
}
func m3() gee.HandlerFunc {
	return func(context *gee.Context) {
		fmt.Println("m3--------")
	}
}

func main() {
	r := gee.New()
	r.Use(gee.Logger()) // global midlleware

	v2 := r.Group("/v2")
	v2.Use(m1()) // v2 group middleware
	v2.Use(m2()) // v2 group middleware
	v2.Use(m3()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
