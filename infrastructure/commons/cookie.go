package commons

import (
	"net/http"
	"time"
)

func CreateCookie(token string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	return cookie
}

func DeleteCookie() *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-24 * time.Hour)
	cookie.HttpOnly = true
	return cookie
}
