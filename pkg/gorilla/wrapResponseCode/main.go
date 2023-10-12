package main

import (
	"github.com/felixge/httpsnoop"
	"log"
	"net/http"
)

func main() {
	// myH is your app's http handler, perhaps a http.ServeMux or similar.
	var myH = &MyHandler{}
	// wrappedH wraps myH in order to log every request.
	wrappedH := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(myH, w, r)
		log.Printf(
			"%s %s (code=%d dt=%s written=%d)",
			r.Method,
			r.URL,
			m.Code,
			m.Duration,
			m.Written,
		)
	})
	http.ListenAndServe(":8080", wrappedH)
}

type MyHandler struct {
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hi"))
}
