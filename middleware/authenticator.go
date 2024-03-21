package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
)

func Authenticate(next func(w http.ResponseWriter, r *http.Request)) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["authorization"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("not authorized 1"))
				}
				return SECRET, nil
			})
			if err != nil {
				fmt.Fprintf(w, err.Error())
				w.Write([]byte("not authorized 2"))
			}
			if token.Valid {
				next(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("not authorized 3"))
		}

	})
}
