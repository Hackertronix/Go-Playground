package main

import (
	"net"
)

func main() {

	listener, err := net.Listener("tcp", "localhost:8080")
	if err != nil{
		log.Fatal(err)

}

for{
 conn,err := listener.Accept()
 if err != nil{
	log.Print(err)
	continue
	

}

}
