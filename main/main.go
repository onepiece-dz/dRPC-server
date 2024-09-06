// Package main -----------------------------
// @author   : deng zhi
// @time     : 2024/9/5 18:00
//
// -------------------------------------------
package main

import (
	drpcserver "github.com/onepiece-dz/drpc-server"
	"log"
	"net"
)

func startServer() {
	// pick a free port
	l, err := net.Listen("tcp", "[::]:8888")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	drpcserver.Accept(l)
}

func main() {
	startServer()
}
