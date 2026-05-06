package routers

import (
	"net/http"
	"restapi/internal/api/handlers"
)

func studentRouter() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.GetStudentsHandler)
	mux.HandleFunc("GET /{student_id}", handlers.GetStudentsHandler)
	mux.HandleFunc("POST /", handlers.AddStudentsHandler)
	mux.HandleFunc("PATCH /{student_id}", handlers.PatchOneStudentsHandler)
	mux.HandleFunc("PATCH /", handlers.PatchStudentsHandler)
	mux.HandleFunc("PUT /{student_id}", handlers.UpdateStudentsHandler)
	mux.HandleFunc("DELETE /", handlers.DeleteStudentsHandler)
	mux.HandleFunc("DELETE /{student_id}", handlers.DeleteOneStudentsHandler)

	return mux
}
