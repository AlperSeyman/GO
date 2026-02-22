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
	mux.HandleFunc("PATCH /teachers/{teacher_id}", handlers.PatchOneTeachersHandler)
	mux.HandleFunc("PATCH /teachers", handlers.PatchTeachersHandler)
	mux.HandleFunc("PUT /teachers/{teacher_id}", handlers.UpdateTeachersHandler)
	mux.HandleFunc("DELETE /teachers/{teacher_id}", handlers.DeleteOneTeachersHandler)

	// student's routes
	mux.HandleFunc("/students/", handlers.StudentsHandler)

	// exec's routes
	mux.HandleFunc("/execs/", handlers.ExecsHandler)

	return mux
}
