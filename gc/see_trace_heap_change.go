package main

import (
	"fmt"
	"net/http"
	"sync"
)

func SeeTraceAndHeapChange() {
	go func() {
		http.ListenAndServe("127.0.0.1:6060", nil)
	}()

	http.HandleFunc("/ex", func(w http.ResponseWriter, r *http.Request) {
		b := bufPool.Get().([]byte)
		//b := bufPool
		for idx := range b {
			b[idx] = 1
		}
		fmt.Fprintf(w, "done, %v", r.URL.Path[1:])
	})

	http.ListenAndServe(":8080", nil)
}

var bufPool = sync.Pool{New: func() interface{} {
	return make([]byte, 10<<20)
}}

//var bufPool = make([]byte, 10<<20)
