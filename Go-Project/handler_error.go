package main

import "net/http"

func HandlerError(w http.ResponseWriter, _ *http.Request) {
	RespondWithError(w, 500, "Something went wrong")
}
