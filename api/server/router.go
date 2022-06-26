package server

import "net/http"

func Register() {
	http.HandleFunc("/test", TestHandler)
}
