package rpc_test

import (
	"educational/lsp/rpc"
	"log"
	"testing"
)

type EncodingExample struct {
	Testing bool
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual, err := rpc.EncodeMessage(EncodingExample{Testing: true})
	if err != nil {
		log.Fatal(err)
	}
	if expected != actual {
		t.Errorf("expected %s,\n got %s\n", expected, actual)
	}
}

func TestDedcode(t *testing.T) {
<<<<<<< HEAD
	expected := 16
	msg := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	contentLength, err := rpc.DecodeMessage([]byte(msg))
	if err != nil {
		log.Fatal(err)
	}
=======
	expected := len("{\"Method\":\"hi\"}")
	msg := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, contentLength, err := rpc.DecodeMessage([]byte(msg))
	if err != nil {
		log.Fatal(err)
	}
	if method != "hi" {
		t.Errorf("Decode Error:-\n expected = %v, got = %v\n", "my friendo", method)
	}
>>>>>>> 0973cc0 (history all f*cked)
	if expected != contentLength {
		t.Errorf("Decode Error:-\n expected = %v, got = %v\n", expected, contentLength)
	}
}
