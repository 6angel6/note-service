package mw

import (
	"Zametki-go/pkg/jwt"
	"fmt"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid token")
			return
		}
		tokenString = tokenString[len("Bearer "):]

		_, err := jwt.ValidateAccessToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid token")
		}
		w.WriteHeader(http.StatusOK)

		next.ServeHTTP(w, r)
	})

}
