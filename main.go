package main

import (
	"fmt"
	"net"
	// "time"
	"strings"
)

func read_udp_connection(address string, recv_chan chan []byte) {
	// defer close(recv_chan)
	addr, addr_er := net.ResolveUDPAddr("udp", address)
	if addr_er != nil {
		fmt.Println("Address Error:", addr_er)
	}

	conn, conn_er := net.ListenUDP("udp", addr)
	if conn_er != nil {
		fmt.Println("Connection Error:", conn_er)
	}

	buffer := make([]byte, 1024)
	for {
		fmt.Println("Reading from:", address)
		len, incomming_addr, read_er := conn.ReadFromUDP(buffer)

		if read_er != nil {
			fmt.Println("Read error:", read_er)
			break
		}
		fmt.Println("Incoming data from:", incomming_addr)

		recv_chan <- buffer[:len]
	}
}

func read_tcp_connection(address string, recv_chan chan []byte) {
	
}

func main() {

	addr_array := []string{
		"udp:127.0.0.1:5760",
		"udp:127.0.0.1:5770",
	}

	recv:= make(chan []byte)

	for i:=0;i<len(addr_array);i++ {
		addr_info := strings.SplitN(addr_array[i], ":", 2)

		switch addr_info[0] {
			case "udp":
				fmt.Println("IT IS A UDP")
				go read_udp_connection(addr_info[1], recv)
			case "tcp":
				fmt.Println("Another TCP")
			default:
				fmt.Println("Unknown type! Type:", addr_info[0])
		}
	}

	fmt.Println("Reading input!")

	for {
		data, open := <- recv

		if open {
			fmt.Println("Data out:", string(data[:]))
		} else {
			fmt.Println("Channel closed!")
			break
		}
	}

}
