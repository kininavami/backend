package session

import (
	"encoding/base64"
	"errors"
	"github.com/vmware/vending/internal/user"
	"net/http"
	"strings"
)

func authenticate(w http.ResponseWriter, r *http.Request) error {
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return errors.New("authorization failed")
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)
		u := user.User{
			Username: pair[0],
			Password: pair[1],
		}
		if len(pair) != 2 || !u.Authenticate() {
			return errors.New("authorization failed")
		}
		return nil
}
