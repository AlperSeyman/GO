package middlewares

import (
	"fmt"
	"net/http"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		fmt.Println(origin)

		r.Header.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		r.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
		r.Header.Set("Access-Control-Credentials", "true")
		next.ServeHTTP(w, r)
	})
}

/*
Basic CORS Middleware Skeleon
func CORSMiddleware(next http.HandlerFunc) http.Handler{
	return  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		next.ServeHTTP(w, r)
	})
}
*/
