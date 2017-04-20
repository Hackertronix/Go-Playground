package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer listener.Close()

	for {

		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		io.WriteString(conn, "\n Hello from the other side\n")
		fmt.Fprintln(conn, "Atleast I can say that I tried")
		fmt.Fprintf(conn, "%v", "To tell you Im sorry")

		conn.Close()

	}
}
