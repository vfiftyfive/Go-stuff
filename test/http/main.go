package main

import "fmt"
import "net/http"

type myHandler struct {
	greeting string
}

func (H myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(fmt.Sprintf("%v world", H.greeting)))
}

func main() {

	http.Handle("/", myHandler{greeting: "hello"})
	http.ListenAndServe(":8000", nil)
}
