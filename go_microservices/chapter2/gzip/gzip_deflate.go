package gzip

import (
	"compress/gzip"
	"net/http"
	"strings"
)

type GzipResponseWriter struct {
	gw *gzip.Writer
	w  http.ResponseWriter
}
type GzipHandler struct {
	next http.Handler
}

// Parameter:
//
//	b []byte - the data to be written
//
// Return:
//
//	int - the number of bytes written
//	error - any error that occurred during writing
func (grw *GzipResponseWriter) Write(b []byte) (int, error) {
	if _, ok := grw.w.Header()["Content-Type"]; !ok {
		//if the content type is not set, infer it from the uncompressed body
		grw.w.Header().Set("Content-Type", http.DetectContentType(b))
	}
	return grw.gw.Write(b)
}

// implementing the serveHttp method for the GzipHandler
func (gh *GzipHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		grw := &GzipResponseWriter{gw: gzip.NewWriter(w), w: w}
		defer grw.gw.Close()
		grw.w.Header().Add("Content-Encoding", "gzip")
		gh.next.ServeHTTP(grw.w, r)
	}
}
