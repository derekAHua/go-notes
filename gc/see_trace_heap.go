package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func SeeTraceAndHeap() {
	go func() {
		http.ListenAndServe("127.0.0.1:6060", nil)
	}()

	http.HandleFunc("/ex", func(w http.ResponseWriter, r *http.Request) {
		b := newBuf()
		for idx := range b {
			b[idx] = 1
		}
		fmt.Fprintf(w, "done, %v", r.URL.Path[1:])
	})

	http.ListenAndServe(":8080", nil)
}

func newBuf() []byte {
	return make([]byte, 10<<20)
}
