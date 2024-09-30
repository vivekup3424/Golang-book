package ch03

import (
	"net"
	"testing"
)
func TestListener(t *testing.T){
  listener,err := net.Listen("tcp", "127.0.0.1:8080")
  if err != nil{
    t.Fatal(err)
  }
  defer listener.Close()
  t.Logf("bound to %v", listener.Addr())
}
func main(){
  
}