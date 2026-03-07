package routers

import (
	"net/http"
	"restapi/internal/api/handlers"
)

func teacherRouter() *http.ServeMux {

	mux := http.NewServeMux()

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

	return mux
}
