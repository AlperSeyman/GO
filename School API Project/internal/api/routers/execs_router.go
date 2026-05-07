package routers

import (
	"net/http"
	"restapi/internal/api/handlers"
)

func execsRouter() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.GetExecsHandler)
	mux.HandleFunc("GET /{exec_id}", handlers.GetExecsHandler)

	mux.HandleFunc("POST /", handlers.AddExecsHandler)
	mux.HandleFunc("POST /login", handlers.LoginExecsHandler)
	mux.HandleFunc("POST /logout", handlers.GetExecsHandler)
	mux.HandleFunc("POST /forgotpassword", handlers.GetExecsHandler)
	mux.HandleFunc("POST /updatepassword/reset/{resetcode}", handlers.GetExecsHandler)
	mux.HandleFunc("POST /{exec_id}/updatepassword", handlers.GetExecsHandler)

	mux.HandleFunc("PATCH /", handlers.PatchExecsHandler)
	mux.HandleFunc("PATCH /{exec_id}", handlers.PatchOneExecsHandler)

	mux.HandleFunc("DELETE /", handlers.DeleteExecsHandler)
	mux.HandleFunc("DELETE /{exec_id}", handlers.DeleteOneExecsHandler)

	return mux

}
