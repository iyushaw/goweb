package main

import (
	"fmt"
	"net/http"
)

func res(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to H2</h1>")
}

func main() {
	http.HandleFunc("/", res)

	fmt.Println("Server Starting on http://localhost:4000")
	err := http.ListenAndServe(":4000", nil)

	if err != nil {
		panic(err)
	}

}
