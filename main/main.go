// Package main -----------------------------
// @author   : deng zhi
// @time     : 2024/9/5 18:00
//
// -------------------------------------------
package main

import (
	"fmt"
	drpcserver "github.com/onepiece-dz/drpc-server"
	"github.com/onepiece-dz/drpc-server/registry"
	"github.com/onepiece-dz/drpc-server/services"
	"log"
	"net"
	"net/http"
)

func registerRpcService() {
	var foo services.Foo
	if err := drpcserver.Register(&foo); err != nil {
		log.Fatal("register foo error:", err)
	}
	var bar services.Bar
	if err := drpcserver.Register(&bar); err != nil {
		log.Fatal("register bar error:", err)
	}
}
func startServer() {
	// pick a free port
	l, err := net.Listen("tcp", "[::]:8888")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	drpcserver.Accept(l)
}

func startTwoServer() {
	go func() {
		// pick a free port
		l, err := net.Listen("tcp", "[::]:8001")
		if err != nil {
			log.Fatal("network error:", err)
		}
		log.Println("start rpc server on", l.Addr())
		drpcserver.Accept(l)
	}()

	go func() {
		// pick a free port
		l, err := net.Listen("tcp", "[::]:8002")
		if err != nil {
			log.Fatal("network error:", err)
		}
		log.Println("start rpc server on", l.Addr())
		drpcserver.Accept(l)
	}()
}

func startTwoServerWithRegistry(registryAddr string) {
	go startServerWithRegistry(registryAddr, "[::]:8001")

	go startServerWithRegistry(registryAddr, "[::]:8002")
}

func startServerWithRegistry(registryAddr string, addr string) {
	l, _ := net.Listen("tcp", addr)
	server := drpcserver.NewServer()
	registry.Heartbeat(registryAddr, "tcp@"+l.Addr().String(), 0)
	log.Println("start rpc server on", l.Addr())
	server.Accept(l)
}

func startHttpServer() {
	l, _ := net.Listen("tcp", ":9999")
	drpcserver.HandleHTTP()
	_ = http.Serve(l, nil)
}

func main() {
	registerRpcService()
	//go startServer()
	//startTwoServer()
	startTwoServerWithRegistry(fmt.Sprintf("http://localhost%s%s", registry.RegistryPort, registry.RegistryDefaultPath))
	startHttpServer()
}
