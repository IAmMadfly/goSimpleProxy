package main

import (
	"fmt"
	"net"
)

func read_connection(conn net.Conn, recv_chan chan []byte) chan []byte {
	bytes := make([]byte, 1024)
	send_chan := make(chan []byte)

	go func() {
		for {
			len, er := conn.Read(bytes)

			if er != nil {
				fmt.Printf("Failed to read connection. Error: %s", er)
				break
			}
			fmt.Println("Sending data")

			send_chan <- bytes[:len]
		}
	}()

	return send_chan
}

func main() {
	conn1, er1 := net.Dial("tcp", "127.0.0.1:5810")
	conn2, er2 := net.Dial("tcp", "127.0.0.1:5820")

	recv_chan := make(chan []byte)

	if er1 != nil && er2 != nil {
		fmt.Printf("Failed to create Connection! Error: %s, %s", er1, er2)
		return
	}

	go read_connection(conn1, recv_chan)
	go read_connection(conn2, recv_chan)

	fmt.Println("All connections made!")

	for {

		fmt.Printf("Recv Data: %d", <-recv_chan)
	}

}
