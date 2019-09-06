package main

import (
	_ "dogego-mini/conf"
	"dogego-mini/protocol/http"
	"dogego-mini/routers"
	"log"
	"os"
)

func main() {
	router := routers.NewRouter()
	server := http.NewHttpProtocol(router)

	err := server.Start(os.Getenv("ADDR"))

	if err != nil {
		log.Fatal(err)
	}
}
