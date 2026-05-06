package routers

import (
	"net/http"
	"restapi/internal/api/handlers"
)

func teacherRouter() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.GetTeachersHandler)
	mux.HandleFunc("GET /{teacher_id}", handlers.GetTeachersHandler)
	mux.HandleFunc("POST /", handlers.AddTeachersHandler)
	mux.HandleFunc("PATCH /{teacher_id}", handlers.PatchOneTeachersHandler)
	mux.HandleFunc("PATCH /", handlers.PatchTeachersHandler)
	mux.HandleFunc("PUT /{teacher_id}", handlers.UpdateTeachersHandler)
	mux.HandleFunc("DELETE /", handlers.DeleteTeachersHandler)
	mux.HandleFunc("DELETE /{teacher_id}", handlers.DeleteOneTeachersHandler)
	mux.HandleFunc("GET /{teacher_id}/students", handlers.GetStudentsByTeacherID)
	mux.HandleFunc("GET /{teacher_id}/studentCount", handlers.GetStudentCountByTeacherId)

	return mux
}
