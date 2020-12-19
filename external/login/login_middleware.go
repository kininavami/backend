package login

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/vmware/vending/external/middleware"
	"github.com/vmware/vending/internal/constants"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == constants.Login{
			next.ServeHTTP(w, r)
		} else {
			w.Header().Set("Content-Type", "application/json")
			tokenString := r.Header.Get("Authorization")
			_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method")
				}
				return []byte("secret"), nil
			})
			if err != nil {
				middleware.RespondError(w, http.StatusUnauthorized, errors.New("UnAuthorized"))
				return
			}
			next.ServeHTTP(w, r)
		}
	})
}