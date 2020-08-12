package main

import (
	"log"
	"github.com/udooz/golog"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}