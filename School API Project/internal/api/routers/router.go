package routers

import (
	"net/http"
	"restapi/internal/api/handlers"
)

func Router() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.RootHandler)

	// teacher's routes
	mux.HandleFunc("GET /teachers", handlers.GetTeachersHandler)
	mux.HandleFunc("GET /teachers/{teacher_id}", handlers.GetTeachersHandler)
	mux.HandleFunc("POST /teachers", handlers.AddTeachersHandler)
	mux.HandleFunc("PUT /teachers/{teacher_id}", handlers.UpdateTeachersHandler)
	mux.HandleFunc("PATCH /teachers/{teacher_id}", handlers.PatchUpdateTeachersHandler)
	mux.HandleFunc("DELETE /teachers/{teacher_id}", handlers.DeleteTeachersHandler)

	mux.HandleFunc("/students/", handlers.StudentsHandler)
	mux.HandleFunc("/execs/", handlers.ExecsHandler)

	return mux
}
