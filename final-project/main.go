package main

import (
	"log"
	"net/http"
)

func main() {
	handler := NewPlayerServer(NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5000", handler))
}
