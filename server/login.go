package server

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"

	"github.com/nturbo1/challenge-tracker-service/log"
)

type LoginPayload struct {
	Username string
	Password string
}
func (lp *LoginPayload) String() string {
	return "{username: " + lp.Username + ", password: " + lp.Password + "}"
}

func handleLogin(rw http.ResponseWriter, req *http.Request) {
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
	var err error
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
	log.Error("IMPLEMENT AUTHENTICATION LOGIC IN HANDLE LOGIN!!!")
}
