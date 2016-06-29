package main

import (
	"log"
	"net/http"

	"github.com/abhigupta912/learngo/proxyserver/proxy"
)

func main() {
	server := proxy.NewProxyServer(nil)
	log.Fatalln(http.ListenAndServe(":9000", server))
}
