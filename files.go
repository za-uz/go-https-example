package main

import (
	"fmt"
	"net/http"
)

func FileHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, fmt.Sprintf("./client/%s", r.URL.Path))
}
