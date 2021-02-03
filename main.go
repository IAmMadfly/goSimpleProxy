package main

import (
	"net"
	"fmt"
)

func main() {
	_, _ = net.Dial("tcp", "127.0.0.1:8080")

	fmt.Println("New network!")
}