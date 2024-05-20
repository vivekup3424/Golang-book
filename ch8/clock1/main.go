package main

import (
	"io"
	"net"
	"time"
)

func main() {
	listener, _ := net.Listen("tcp", "localhost:8080")
	for {
		conn, _ := listener.Accept()
		handleConn(conn)
	}
}
func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		io.WriteString(conn, time.Now().String()+"\n")
		//fmt.Printf("length of output:= %d", n)
		time.Sleep(1 * time.Second)

	}
}
