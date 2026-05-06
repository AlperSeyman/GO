package routers

import (
	"context"
	"net/http"
)

func mount(mux *http.ServeMux, prefix string, sub *http.ServeMux) {
	mux.Handle(prefix+"/", http.StripPrefix(prefix, sub))
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		r2 := r.Clone(context.Background())
		r2.URL.Path = "/"
		sub.ServeHTTP(w, r2)
	})
}

func MainRouter() *http.ServeMux {

	mux := http.NewServeMux()

	mount(mux, "/execs", execsRouter())
	mount(mux, "/students", studentRouter())
	mount(mux, "/teachers", teacherRouter())

	return mux

}
