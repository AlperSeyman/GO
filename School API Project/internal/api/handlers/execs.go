package handlers

import "net/http"

func ExecsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Execs Page"))
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("GET Method on Execs Route Page"))
	case http.MethodPost:
		w.Write([]byte("POST Method on Execs Route Page"))
	case http.MethodPut:
		w.Write([]byte("PUT Method on Execs Route Page"))
	case http.MethodDelete:
		w.Write([]byte("DELETE Method on Execs Route Page"))
	default:
		w.Write([]byte("Execs Route Page"))
	}
}
