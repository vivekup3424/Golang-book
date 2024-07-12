package handlers

import (
	"io"
	"log"
	"net/http"
)

// creats an http handler
type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}
func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bye"))
	n, err := io.Copy(w, r.Body)
	defer r.Body.Close()
	if err != nil {
		g.l.Println(err)
		gttp.Error(w, "internal server error", http.StatusInternalServerError)
	}
	g.l.Println(r.Header)
	g.l.Printf("Served %d bytes\n", n)
}
