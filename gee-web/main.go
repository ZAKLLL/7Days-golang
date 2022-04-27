package main

import (
	"fmt"
	"gee"
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

	r.Static("/static", "D://")
	r.Run(":9999")
}
