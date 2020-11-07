package auth

import (
	"account/crypto"
	"net/http"
	"time"
)

func LoginUser(uuid string, support, rememberMe bool, res http.ResponseWriter) {
	var dur time.Duration
	var maxAge int
	if rememberMe {
		dur = longSession
		maxAge = 0
	} else {
		dur = shortSession
		maxAge = int(dur.Seconds())
	}
	token, err := crypto.SessionToken(uuid, signingSecret, support, dur)
	if err != nil {
		panic(err)
	}
	cookie := http.Cookie{
		Name:   cookieName,
		Value:  token,
		Path:   "/",
		MaxAge: maxAge,
		Domain: "." + "joywork.local",
	}
	http.SetCookie(res, &cookie)
}

func getSession(req *http.Request) (uuid string, support bool, err error) {
	cookie, err := req.Cookie(cookieName)
	if err != nil {
		return
	}
	uuid, support, err = crypto.RetrieveSessionInformation(cookie.Value, signingSecret)
	return
}

func Logout(res http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   cookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
		Domain: "." + "joywork.local",
	}
	http.SetCookie(res, cookie)
}
