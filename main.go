package main

import (
	"fmt"
	"minusblog/minus"
	_ "minusblog/router"
	"time"
)

func main() {
	fmt.Print("main")

	config := &minus.Config{
		HttpAddr:     "localhost",
		HttpPort:     9090,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	minus.Run(config)
}
