package customErrors

import (
	"time"
)

type SessionExpiredError struct {
	SessionId string
	ExpiredAt time.Time
}
func (see *SessionExpiredError) Error() string {
	return "Session with id '" + see.SessionId + "' expired at " + see.ExpiredAt.String()
}

type SessionInvalidError struct {
	SessionId string
}
func (sie *SessionInvalidError) Error() string {
	return "Session with id '" + sie.SessionId + "' is invalid."
}

type SessionExistsError struct {
	SessionId string
}
func (see *SessionExistsError) Error() string {
	return "Session with id '" + see.SessionId + "' already exists."
}
