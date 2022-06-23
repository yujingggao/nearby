package main

import (
	"fmt"
	"log"
	"net/http"

	"nearby/backend"
	"nearby/handler"
)

func main() {
	fmt.Println("started-service")
	backend.InitElasticsearchBackend()
	backend.InitGCSBackend()
	log.Fatal(http.ListenAndServe(":8080", handler.InitRouter()))
}
