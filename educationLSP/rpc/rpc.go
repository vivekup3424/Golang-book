package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type Message struct {
	ContentLength int
	Content       string
}

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		log.Println("unable to encode message: ", err)
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n", len(content)) + string(content)
}

func DecodeMessage(msg []byte) error {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return errors.New("Did not find the seperator")
	}
	//Content-Length = <number>
	contentLengthBytes := header["Content-Length: "]
	return nil
}
