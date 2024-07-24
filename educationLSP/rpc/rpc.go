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
<<<<<<< HEAD
=======
type BaseMessage struct {
	Method string `json:"method"`
}
>>>>>>> 0973cc0 (history all f*cked)

func EncodeMessage(msg any) (string, error) {
	content, err := json.Marshal(msg)
	if err != nil {
		log.Println("unable to encode message: ", err)
		return "", err
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n", len(content)) + string(content), nil
}

<<<<<<< HEAD
func DecodeMessage(msg []byte) (int, error) {
	header, _, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return -1, errors.New("did not find the seperator")
=======
func DecodeMessage(msg []byte) (string, int, error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return "", -1, errors.New("did not find the seperator")
>>>>>>> 0973cc0 (history all f*cked)
	}
	//Content-Length = <number>
	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		log.Fatal(err)
	}
<<<<<<< HEAD
	fmt.Println("Content-length = ", contentLength)
	return contentLength, nil
=======
	fmt.Println("Content-length error= ", contentLength)
	var baseMessage BaseMessage
	err = json.Unmarshal(content[:contentLength], &baseMessage)
	if err != nil {
		log.Fatal(err)
	}
	return baseMessage.Method, contentLength, nil
>>>>>>> 0973cc0 (history all f*cked)
}
