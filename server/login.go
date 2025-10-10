package server

import (
	"net/http"
	"encoding/json"
	"io"
	"time"

	"github.com/nturbo1/challenge-tracker-service/log"
	"github.com/nturbo1/challenge-tracker-service/db"
	"github.com/nturbo1/challenge-tracker-service/db/session"
)

const (
	sessionCookieName = "SESS_jhkqjerqwqwe_ID"
	loginSessionCookieMaxAge = 7 * 24 * 60 * 60 // in seconds
)

type LoginPayload struct {
	Username string
	Password string
}
func (lp *LoginPayload) String() string {
	return "{username: " + lp.Username + ", password: " + lp.Password + "}"
}

func handleLogin(rw http.ResponseWriter, req *http.Request) {
	var err error
	log.HttpRequest(req)

	if req.ContentLength <= 0 {
		log.Error(
			"Request body content length is %d. Can't deserialize or parse the content " + 
			"if I don't know the length :(... Or can I? o.O", 
			req.ContentLength,
		)
		rw.WriteHeader(400)
		return
	}

	buf := make([]byte, req.ContentLength)
	n, err := req.Body.Read(buf)

	if err != nil && err != io.EOF {
		log.Error("%s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	if int64(n) < req.ContentLength { 
		// TODO: Try multiple times instead of giving up immediately!
		log.Error("Read %d bytes, while the content length is %d", n, req.ContentLength)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	var payload LoginPayload
	err = json.Unmarshal(buf, &payload)

	if err != nil {
		log.Error("%s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Info("Payload: %s", &payload)

	user := db.UserRepository.FindByUsername(payload.Username)
	if user == nil {
		log.Error("User with username %s wasn't found.", payload.Username)
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	if user.Password != payload.Password {
		log.Error("Password didn't match for the user with username %s", user.Username)
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	sessId := generateSessionId()
	err = saveSession(sessId, user.Id)

	if err != nil {
		log.Error("Failed to save a session with id = %s", sessId)
		log.Error("%s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	sessCookie := generateLoginSessionCookie(sessId, "localhost")
	rw.Header().Set("Set-Cookie", sessCookie.String())
}

func generateSessionId() string {
	log.Error("IMPLEMENT SESSION ID GENERATOR!!!")
	return "atqwhkl02h320asd8wq98ufasd"
}

func generateLoginSessionCookie(sessionId string, domain string) *http.Cookie {
	return &http.Cookie{
		Name: sessionCookieName,
		Value: sessionId,
		MaxAge: loginSessionCookieMaxAge, // REMINDER: should be in seconds!!!
		Domain: domain,
		Secure: true, // IMPORTANT: Set it to true on production environment!!!
		HttpOnly: true,
		Path: "/",
	}
}

func saveSession(sessId string, userId int) error {
	currTime := time.Now()
	var duration time.Duration = loginSessionCookieMaxAge * 1000 * 1000 * 1000 // converted to nano seconds
	expirationTime := currTime.Add(duration)
	newSess := &session.SessionInfo{
		UserId: userId,
		CreatedAt: currTime,
		ExpiresAt: expirationTime,
	}

	log.Info("Saving a new session: %s", newSess)
	return db.SessionRepository.AddSession(sessId, newSess)
}
