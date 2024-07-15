package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
)

type Message struct {
	ContentLength int
	Content       string
}

func EncodeMessage(msg any) (string, error) {
	content, err := json.Marshal(msg)
	if err != nil {
		log.Println("unable to encode message: ", err)
		return "", err
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n", len(content)) + string(content), nil
}

func DecodeMessage(msg []byte) (int, error) {
	header, _, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return -1, errors.New("did not find the seperator")
	}
	//Content-Length = <number>
	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Content-length = ", contentLength)
	return contentLength, nil
}
