// Package main -----------------------------
// @author   : deng zhi
// @time     : 2024/9/11 11:40
//
// -------------------------------------------
package main

import (
	"github.com/onepiece-dz/drpc-server/registry"
	"net"
	"net/http"
)

func main() {
	startRegistry()
}

func startRegistry() {
	l, _ := net.Listen("tcp", registry.RegistryPort)
	registry.HandleHTTP()
	_ = http.Serve(l, nil)
}
