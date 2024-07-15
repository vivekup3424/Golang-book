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
	expected := 16
	msg := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	contentLength, err := rpc.DecodeMessage([]byte(msg))
	if err != nil {
		log.Fatal(err)
	}
	if expected != contentLength {
		t.Errorf("Decode Error:-\n expected = %v, got = %v\n", expected, contentLength)
	}
}
