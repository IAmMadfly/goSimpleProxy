package main

import (
	"fmt"
	"net"
	"strings"
	// "net"
)

func read_connection(addr string, output chan []byte) {

	split_addr := strings.SplitN(addr, ":", 2)

	fmt.Println("Address information:", split_addr)

	conn, err := net.Dial(split_addr[0], split_addr[1])
	fmt.Println("Connection:", conn, "Error:", err)

	if err != nil {
		return
	}

	buffer := make([]byte, 1024)

	
	for {
		fmt.Println("Reading data on:", addr)
		length, read_err := conn.Read(buffer)
		fmt.Println("Data length:", length)
		if read_err != nil {
			fmt.Println("Data read error on:", addr)
		}

		output <- buffer[:length]
	}
	
} 

func main() {

	server_addr, _ := net.ResolveUDPAddr("udp", ":5760")
	
	server_conn, _ := net.ListenUDP("udp", server_addr)

	defer server_conn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := server_conn.ReadFromUDP(buf)

		if err != nil {
			fmt.Println("Got error:", err)
		}

		fmt.Println("Data:", string(buf[0:n]), "from", addr)
	}


	// var addr_array []string

	// addr_array = append(addr_array, "udp:127.0.0.1:5760", "udp:127.0.0.1:5770")
	// output_chan := make(chan []byte)

	// read_connection("udp:127.0.0.1:5760", output_chan)

	// for i := 0; i<len(addr_array); i++ {
	// 	fmt.Println("Addr:", addr_array[i])
	// 	go read_connection(addr_array[i], output_chan)
	// }

	// for {
	// 	output, open := <- output_chan

	// 	if open {
	// 		fmt.Println("Incoming data:", output)
	// 	} else {
	// 		fmt.Println("Main channel closed! Program closing!")
	// 		break
	// 	}
	// }
}
