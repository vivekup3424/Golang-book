package server

import (
	"log"
	"net"
	"testing"
)

func TestListener(t *testing.T){
    listener, err := net.Listen("tcp", "localhost:8080")
    if err != nil{
        log.Fatal(err)
    }
    defer func(){
        _ := listener.Close()
    }
    t.Logf("bound to %v", listener.Addr())
}
