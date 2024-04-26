package main

import (
	"github.com/clarechu/go-client/example"
	"time"
)

func main() {
	err := example.Socket("localhost:8080", "", "")
	if err != nil {
		panic(err)
	}
	time.Sleep(100 * time.Second)
}
