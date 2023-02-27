package main

import (
	"example/gee"
	"fmt"
	"net/http"
)

func main() {
	server := gee.New()

	server.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})
	server.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s\n", c.Query("name"))
	})

	server.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	fmt.Println("Start listening on port 8080...")
	server.Run(":8080")
}
