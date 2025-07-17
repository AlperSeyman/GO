package main

import "net/http"

func HandlerReadiness(w http.ResponseWriter, _ *http.Request) {
	RespondWithJson(w, 200, struct{}{})
}
