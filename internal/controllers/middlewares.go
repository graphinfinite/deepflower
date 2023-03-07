package controllers

import (
	"net/http"
	"strings"
)

func (auth *AuthController) JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")
		if !(len(bearer) > 7) || !(strings.ToUpper(bearer[0:6]) == "BEARER") {
			JSON(w, STATUS_ERROR, "bearer token unrecognize<-header.Authorization")
			return
		}
		token := bearer[7:]
		ok, _, err := auth.Uc.ValidateJwtToken(token)
		if err != nil || !ok {
			JSON(w, STATUS_ERROR, "token invalid")
			return
		}

		//ctx := context.WithValue(r.Context(), some, claims)
		//next.ServeHTTP(w, r.WithContext(ctx))
		next.ServeHTTP(w, r)
	})

}
