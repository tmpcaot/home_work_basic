package main

import (
	"flag"

	"github.com/tmpcaot/HOME_WORK_BASIC/hw13_http/pkg/server"
)

func main() {
	addr := flag.String("addr", ":8080", "Server address and port")
	flag.Parse()

	server.RunServer(*addr)
}
