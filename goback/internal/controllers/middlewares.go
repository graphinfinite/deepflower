package controllers

import (
	"context"
	"net/http"
	"strings"
)

type ContextKey string

const ContextUserIdKey ContextKey = "userId"

func (auth *AuthController) JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")
		if !(len(bearer) > 7) || !(strings.ToUpper(bearer[0:6]) == "BEARER") {
			JSON(w, STATUS_ERROR, "bearer token unrecognize<-header.Authorization")
			return
		}
		token := bearer[7:]
		ok, claims, err := auth.Uc.ValidateJwtToken(r.Context(), token)
		if err != nil || !ok {
			JSON(w, STATUS_ERROR, "token invalid")
			return
		}
		userId, err := claims.GetSubject()
		if err != nil {
			JSON(w, STATUS_ERROR, "no subject in sub")
		}
		ctx := context.WithValue(r.Context(), ContextUserIdKey, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
		//next.ServeHTTP(w, r)
	})

}
