// Package main1 -----------------------------
// @author   : deng zhi
// @time     : 2024/9/10 15:40
//
// -------------------------------------------
package main

import (
	"log"
	"time"
)

func main() {
	called := make(chan struct{}, 1)
	go func() {
		called <- struct{}{} // 阻塞

		log.Println("1111")
	}()
	time.Sleep(5 * time.Second)
	a := <-called
	log.Println(a)
	time.Sleep(3 * time.Second)
}
