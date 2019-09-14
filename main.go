package main

import (
	_ "dogego-mini/conf"
	"dogego-mini/protocol/http"
	"dogego-mini/protocol/http2"
	"dogego-mini/routers"
	"log"
	"os"
)

func main() {
	router := routers.NewRouter()
	httpServer := http.NewHttpProtocol(router)

	go func() {
		err := httpServer.Start(os.Getenv("ADDR_HTTP"))

		if err != nil {
			log.Fatal(err)
		}
	}()

	if os.Getenv("ENABLE_JRPC") == "enable" {
		http2Server := http2.NewHttp2Protocol(router)
		err := http2Server.Start(os.Getenv("ADDR_HTTP2"))

		if err != nil {
			log.Fatal(err)
		}
	}
}
