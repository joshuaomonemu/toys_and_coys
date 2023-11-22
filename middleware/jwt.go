package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

var SECRET = []byte("super-secret-key")
var api_key = "1234"

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["ExpiresAt"] = time.Now().Add(100 * time.Minute)

	tokenString, err := token.SignedString(SECRET)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(next func(w http.ResponseWriter, r *http.Request)) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
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

func GetJwt(w http.ResponseWriter, r *http.Request) {
	if r.Header["Access"] != nil {
		if r.Header["Access"][0] != api_key {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("not authorized"))
			return
		} else {
			token, err := GenerateJWT()
			if err != nil {
				fmt.Fprintf(w, err.Error())
				return
			}
			fmt.Fprintf(w, token)
		}
	}
}
