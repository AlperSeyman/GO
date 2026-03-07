package routers

import (
	"net/http"
	"restapi/internal/api/handlers"
)

func Router() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.RootHandler)

	// teacher's routes
	mux.HandleFunc("GET /teachers", handlers.GetTeachersHandler)
	mux.HandleFunc("GET /teachers/{teacher_id}", handlers.GetTeachersHandler)
	mux.HandleFunc("POST /teachers", handlers.AddTeachersHandler)
	mux.HandleFunc("PATCH /teachers/{teacher_id}", handlers.PatchOneTeachersHandler)
	mux.HandleFunc("PATCH /teachers", handlers.PatchTeachersHandler)
	mux.HandleFunc("PUT /teachers/{teacher_id}", handlers.UpdateTeachersHandler)
	mux.HandleFunc("DELETE /teachers", handlers.DeleteTeachersHandler)
	mux.HandleFunc("DELETE /teachers/{teacher_id}", handlers.DeleteOneTeachersHandler)
	mux.HandleFunc("GET /teachers/{teacher_id}/students", handlers.GetStudentsByTeacherID)
	mux.HandleFunc("GET /teachers/{teacher_id}/studentCount", handlers.GetStudentCountByTeacherId)

	// student's routes
	mux.HandleFunc("GET /students", handlers.GetStudentsHandler)
	mux.HandleFunc("GET /stutents/{student_id}", handlers.GetStudentsHandler)
	mux.HandleFunc("POST /students", handlers.AddStudentsHandler)
	mux.HandleFunc("PATCH /students/{student_id}", handlers.PatchOneStudentsHandler)
	mux.HandleFunc("PATCH /students", handlers.PatchStudentsHandler)
	mux.HandleFunc("PUT /students/{student_id}", handlers.UpdateStudentsHandler)
	mux.HandleFunc("DELETE /students", handlers.DeleteStudentsHandler)
	mux.HandleFunc("DELETE /students/{student_id}", handlers.DeleteOneStudentsHandler)

	// exec's routes
	// mux.HandleFunc("/execs/", handlers.ExecsHandler)

	return mux
}
