package main

import (
	"example/gee"
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := gee.New()

	server.GET("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("hello world"))
		if err != nil {
			log.Println(err)
		}
	})

	server.GET("/home", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("home page"))
		if err != nil {
			log.Println(err)
		}
	})

	fmt.Println("Start listening on port 8080...")
	server.Run(":8080")
}
