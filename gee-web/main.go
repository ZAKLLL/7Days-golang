package main

import (
	"log"
	"net/http"

	"gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hi/:user/do", func(c *gee.Context) {
		log.Printf("/hi/:user/do")
		c.String(http.StatusOK, "hello %s, you're at \n", c.Query("user"))
	})

	r.GET("/hi/*/v1", func(c *gee.Context) {
		log.Printf("/hi/*/v1")
		c.String(http.StatusOK, "hello %s, you're at \n", c.Query("user"))
	})

	r.GET("/hi/*", func(c *gee.Context) {
		log.Printf("/hi/*")
		c.String(http.StatusOK, "hello %s, you're at \n", c.Query("user"))
	})

	r.GET("/hi/*/:lang/zzz", func(c *gee.Context) {
		log.Printf("/hi/*/:lang/zzz")
		c.String(http.StatusOK, "hello %s, you're at \n", c.Query("user"))
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
