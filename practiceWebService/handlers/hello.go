package handlers

import (
	"io"
	"log"
	"net/http"
)

// creats an http handler
type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l: l}
}
func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	n, err := io.Copy(w, r.Body)
	defer r.Body.Close()
	if err != nil {
		h.l.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	h.l.Println(r.Header)
	h.l.Printf("Served %d bytes\n", n)
}
