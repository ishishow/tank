package main

import (
	"flag"
	"os"

	"20dojo-online/pkg/server"
)

var (
	// Listenするアドレス+ポート
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":"+os.Getenv("PORT"), "tcp host:port to connect")
	flag.Parse()
}

func main() {
	server.Serve(addr)
}
