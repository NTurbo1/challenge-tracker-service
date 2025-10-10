package session

import (
	"time"
	"os"
	"strconv"

	"github.com/nturbo1/challenge-tracker-service/customErrors"
	"github.com/nturbo1/challenge-tracker-service/log"
)

type SessionInfo struct {
	UserId int
	CreatedAt time.Time
	ExpiresAt time.Time
}
func (si *SessionInfo) String() string {
	return strconv.Itoa(si.UserId) + "," + si.CreatedAt.Format(timeLayout) + "," + si.ExpiresAt.Format(timeLayout)
}

type SessionRepo struct {
	File *os.File
	SessionsMap map[string]SessionInfo // keys used in the map are session ids.
}

// Function findSession returns a session info by a given session id.
// The returned session info is nil if there's no session by that id.
// The function returns SessionExpiredError instance if the found session's expired
// SessionInvalidError if the found session's invalid.
func (sp *SessionRepo) FindSession(id string) (*SessionInfo, error) {
	if sessionInfo, exists := sp.SessionsMap[id]; exists {
		currTime := time.Now()
		if currTime.Compare(sessionInfo.ExpiresAt) >= 0 {
			sp.DeleteSession(id)
			return nil, &customErrors.SessionExpiredError{id, sessionInfo.ExpiresAt}
		}

		return &sessionInfo, nil
	}
	
	return nil, nil
}

// Function AddSession adds a new session to the sessions storage.
// A returned nil error value means success.
// Returns SessionExistsError instance if there's a session with a given id already stored.
func (sp *SessionRepo) AddSession(id string, sessionInfo *SessionInfo) error {
	if err := sp.verifySessionNotExists(id); err != nil {
		return err
	}

	sp.SessionsMap[id] = *sessionInfo
	return nil
}

func (sp *SessionRepo) DeleteSession(id string) {
	delete(sp.SessionsMap, id)
}

func (sp *SessionRepo) FlushAllData() error {
	writeData := sessionCSVHeader + "\n"
	for id, session := range sp.SessionsMap {
		row := id + "," + session.String() + "\n"
		writeData += row
	}
	_, err := sp.File.Write([]byte(writeData))
	if err != nil { // TODO: You can implement multiple tries instead of giving up immediately!
		return err
	}

	return nil
}

// Function verifySessionNotExists checks if there's a session with a given id that
// is already stored.
// The returned error is nil if there's no session with a given id.
func (sp *SessionRepo) verifySessionNotExists(id string) error {
	foundSession, err := sp.FindSession(id)
	if err != nil {
		return err
	}
	if foundSession != nil {
		return &customErrors.SessionExistsError{id}
	}

	return nil
}

func (sp *SessionRepo) Close() error {
	log.Info("Closing the session repo...")
	err := sp.File.Close()
	if err != nil {
		log.Error("While closing the session repo file: ", err)
	}

	return err
}
