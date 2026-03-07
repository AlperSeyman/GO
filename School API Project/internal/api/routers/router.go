package routers

import (
	"net/http"
)

func MainRouter() *http.ServeMux {

	mainMux := http.NewServeMux()

	mainMux.Handle("/teachers/", teacherRouter())
	mainMux.Handle("/students/", studentRouter())
	mainMux.Handle("/execs/", execsRouter())
	return mainMux
}
