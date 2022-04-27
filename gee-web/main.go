package main

import (
	"fmt"
	"gee"
	"net/http"
)

func m0() gee.HandlerFunc {
	return func(context *gee.Context) {
		fmt.Println("m0--------")
	}
}

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
func m4() gee.HandlerFunc {
	return func(context *gee.Context) {
		fmt.Println("m4--------")
	}
}
func m5() gee.HandlerFunc {
	return func(context *gee.Context) {
		fmt.Println("m5--------")
	}
}

func main() {
	r := gee.New()
	r.Use(gee.Logger()) // global midlleware

	v1 := r.Group("/:name")
	v2 := v1.Group("/v2")
	v2.GET("/hi", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	v3 := r.Group("/v1")
	v4 := v3.Group("/v2")
	v4.GET("/hi", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	v1.Use(m0())
	v2.Use(m1()) // v2 group middleware
	v2.Use(m2()) // v2 group middleware
	v2.Use(m3()) // v2 group middleware

	v3.Use(m4())
	v4.Use(m5()) // v2 group middleware

	//v1.Use(m5())
	//{
	//	v2.GET("/hello/:name", func(c *gee.Context) {
	//		// expect /hello/geektutu
	//		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	//	})
	//}

	r.Run(":9999")
}
