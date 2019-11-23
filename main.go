package main

import (
	"os"

	"github.com/leobrines/pfinance-api/webserver"
)

func main() {
	os.Setenv("PORT", "8080")

	server := webserver.New()
	server.Start()
}
