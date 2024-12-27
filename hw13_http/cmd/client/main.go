package main

import (
	"flag"
	"log"
	"os"

	"github.com/tmpcaot/HOME_WORK_BASIC/hw13_http/pkg/client"
)

func main() {
	serverURL := flag.String("url", "http://localhost:8080", "Server URL")
	resourcePath := flag.String("path", "/api/example", "Resource path")
	method := flag.String("method", "GET", "HTTP Method (GET or POST)")
	flag.Parse()
	if *serverURL == "" || *resourcePath == "" {
		log.Println("Usage: client -url <server_url> -path <resource_path> -method <get|post>")
		os.Exit(1)
	}
	client.RunClient(*serverURL, *resourcePath, *method)
}
