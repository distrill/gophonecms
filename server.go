package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"

	"github.com/distrill/gophonecms/controllers"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func main() {
	port := getPort()
	router := httprouter.New()

	// normal route handlers
	router.GET("/", controllers.IndexHandler)
	router.GET("/gallery", controllers.GalleryHandler)
	router.GET("/gallery/:name", controllers.SubgalleryHandler)
	router.POST("/message", controllers.MessageHandler)

	// static content handler
	router.ServeFiles("/public/*filepath", http.Dir("public"))

	fmt.Printf("gophonecms is serving from http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
