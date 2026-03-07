package routers

import (
	"net/http"
	"restapi/internal/api/handlers"
)

func execsRouter() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /execs", handlers.ExecsHandler)
	mux.HandleFunc("GET /execs/{execs_id}", handlers.ExecsHandler)

	mux.HandleFunc("POST /execs", handlers.ExecsHandler)
	mux.HandleFunc("POST /execs/login", handlers.ExecsHandler)
	mux.HandleFunc("POST /execs/logout", handlers.ExecsHandler)
	mux.HandleFunc("POST /execs/forgotpassword", handlers.ExecsHandler)
	mux.HandleFunc("POST /execs/updatepassword/reset/{resetcode}", handlers.ExecsHandler)
	mux.HandleFunc("POST /execs/{execs_id}/updatepassword", handlers.ExecsHandler)

	mux.HandleFunc("PATCH /execs", handlers.ExecsHandler)
	mux.HandleFunc("PATCH /execs/{execs_id}", handlers.ExecsHandler)

	mux.HandleFunc("DELETE /execs/{execs_id}", handlers.ExecsHandler)

	return mux

}
