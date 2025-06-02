package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	fmt.Println("Listening on port :6379")

	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		return
	}

	defer conn.Close()

	for {
		buf := make([]byte, 1024)

		_, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading from connection:", err.Error())
			os.Exit(1)

		}

		conn.Write([]byte("+OK\r\n"))
	}
}
