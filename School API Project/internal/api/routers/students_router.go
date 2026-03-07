package routers

import (
	"net/http"
	"restapi/internal/api/handlers"
)

func studentRouter() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /students", handlers.GetStudentsHandler)
	mux.HandleFunc("GET /stutents/{student_id}", handlers.GetStudentsHandler)
	mux.HandleFunc("POST /students", handlers.AddStudentsHandler)
	mux.HandleFunc("PATCH /students/{student_id}", handlers.PatchOneStudentsHandler)
	mux.HandleFunc("PATCH /students", handlers.PatchStudentsHandler)
	mux.HandleFunc("PUT /students/{student_id}", handlers.UpdateStudentsHandler)
	mux.HandleFunc("DELETE /students", handlers.DeleteStudentsHandler)
	mux.HandleFunc("DELETE /students/{student_id}", handlers.DeleteOneStudentsHandler)

	return mux
}
