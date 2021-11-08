package cookies

import (
	"fmt"
	"net/http"
)

type Jar struct {
	Secure bool
}

func (j Jar) Clear(w http.ResponseWriter, key string) {
	http.SetCookie(w, &http.Cookie{
		Name:     key,
		Value:    "",
		Path:     "/",
		Secure:   j.Secure,
		HttpOnly: true,
		MaxAge:   -1,
	})
}

func (j Jar) Get(r *http.Request, key string) (string, error) {
	cookie, err := r.Cookie(key)
	if err != nil {
		return "", fmt.Errorf("get cookie failed: %w", err)
	}
	return cookie.Value, nil
}

func (j Jar) Set(w http.ResponseWriter, key, value string) {
	http.SetCookie(w, &http.Cookie{
		Name:     key,
		Value:    value,
		Path:     "/",
		Secure:   j.Secure,
		HttpOnly: true,
	})
}
