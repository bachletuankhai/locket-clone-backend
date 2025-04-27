package main

import "net/http"

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi"))
}

func main() {
	mux := http.NewServeMux()

	handler := homeHandler{}
	mux.Handle("/", &handler)

	http.ListenAndServe("127.0.0.1:8080", mux)
}