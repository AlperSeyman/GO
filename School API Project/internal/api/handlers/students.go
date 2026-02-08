package handlers

import "net/http"

func StudentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Student Page"))
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("GET Method on Students Route Page"))
	case http.MethodPost:
		w.Write([]byte("POST Method on Students Route Page"))
	case http.MethodPut:
		w.Write([]byte("PUT Method on Students Route Page"))
	case http.MethodDelete:
		w.Write([]byte("DELETE Method on Students Route Page"))
	default:
		w.Write([]byte("Students Route Page"))
	}
}
