package handlers

import "net/http"

func RootHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Main Page")
	w.Write([]byte("Main Page"))
}
