package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	fmt.Println("Listening on port :6379")

	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Println("Failed to listen:", err)
		return
	}

	aof, err := NewAof("database.aof")
	if err != nil {
		log.Println("Failed to open AOF:", err)
		return
	}
	defer aof.Close()

	aof.Read(func(value Value) {
		command := strings.ToUpper(value.array[0].bulk)
		args := value.array[1:]

		handler, ok := Handlers[command]
		if !ok {
			log.Println("Invalid command in AOF:", command)
			return
		}

		handler(args)
	})

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Failed to accept connection:", err)
			continue
		}

		go handleConnection(conn, aof)
	}
}

func handleConnection(conn net.Conn, aof *Aof) {
	defer conn.Close()

	for {
		resp := NewResp(conn)
		value, err := resp.Read()
		if err != nil {
			log.Println("RESP Read error:", err)
			return
		}

		if value.typ != "array" {
			log.Println("Invalid request: expected array")
			continue
		}

		if len(value.array) == 0 {
			log.Println("Invalid request: empty array")
			continue
		}

		command := strings.ToUpper(value.array[0].bulk)
		args := value.array[1:]

		writer := NewWriter(conn)

		handler, ok := Handlers[command]
		if !ok {
			log.Println("Invalid command:", command)
			writer.Write(Value{typ: "error", str: "ERR unknown command"})
			continue
		}

		if command == "SET" || command == "HSET" {
			if err := aof.Write(value); err != nil {
				log.Println("AOF Write error:", err)
			}
		}

		var result Value
		func() {
			defer func() {
				if r := recover(); r != nil {
					log.Println("Recovered from panic:", r)
					result = Value{typ: "error", str: "ERR internal server error"}
				}
			}()
			result = handler(args)
		}()

		writer.Write(result)
	}
}
