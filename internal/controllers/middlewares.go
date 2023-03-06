package controllers

import (
	"context"
	"net/http"
	"strings"
)

func (auth *AuthController) JWT(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// create new context from `r` request context, and assign key `"user"`
		// to value of `"123"`
		ctx := context.WithValue(r.Context(), "user", "123")

		bearer := r.Header.Get("Authorization")
		if !(len(bearer) > 7) || !(strings.ToUpper(bearer[0:6]) == "BEARER") {
			JSON(w, STATUS_ERROR, "bearer token unrecognize<-header.Authorization")
			return
		}

		//token := bearer[7:]

		next.ServeHTTP(w, r.WithContext(ctx))
	})

}
