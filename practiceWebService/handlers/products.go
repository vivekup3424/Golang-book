package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	wdata "web-service/data"
)

type Products struct {
	l    *log.Logger
	data []*wdata.Product
}

func NewProducts(l *log.Logger) *Products {
	return &Products{
		l:    l,
		data: wdata.ProductList,
	}
}

func (p *Products) getProducts() ([]byte, error) {
	return json.Marshal(p.data)
}
func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b, err := p.getProducts()
		if err != nil {
			http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
