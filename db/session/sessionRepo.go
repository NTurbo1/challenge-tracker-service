package session

import (
	"time"
	"os"
	"strconv"
	"fmt"
	"encoding/binary"

	"github.com/nturbo1/challenge-tracker-service/customErrors"
	"github.com/nturbo1/challenge-tracker-service/log"
	"github.com/nturbo1/challenge-tracker-service/util"
)

type SessionInfo struct { // id is used as a key in the sessions map managed by SessionRepo
	UserId int64
	CreatedAt time.Time
	ExpiresAt time.Time
	Valid bool
	Offset int64
}
func (si *SessionInfo) String() string {
	return strconv.Itoa(int(si.UserId)) + "," + si.CreatedAt.Format(timeLayout) + "," + 
	si.ExpiresAt.Format(timeLayout) + "," + strconv.FormatBool(si.Valid) + "," + 
	strconv.Itoa(int(si.Offset))
}

type SessionRepo struct {
	File *os.File
	ValidSessionsMap map[string]SessionInfo // keys used in the map are session ids.
}

// Function findSession searches for a valid session with a given id and checks if the found session
// has expired. If the found session's expired, then it invalidates it and returns
// SessionExpiredError instance. The function also returns any error occurred during the session
// invalidation process. If no session is found by a given id, then it simply returns nil with 
// no error.
func (sp *SessionRepo) FindValidSession(id string) (*SessionInfo, error) {
	if sessionInfo, exists := sp.ValidSessionsMap[id]; exists {
		currTime := time.Now()
		if currTime.Compare(sessionInfo.ExpiresAt) >= 0 {
			err := sp.InvalidateSession(id)
			if err != nil {
				return nil, err
			}
			return nil, &customErrors.SessionExpiredError{id, sessionInfo.ExpiresAt}
		}

		return &sessionInfo, nil
	}
	
	return nil, nil
}

// Function AddSession adds a new session to the sessions storage.
// A returned nil error value means success.
// Returns SessionExistsError instance if there's a valid/invalid session with a given id already stored.
func (sp *SessionRepo) AddSession(id string, sessInfo *SessionInfo) error {
	if err := sp.verifySessionNotExists(id); err != nil {
		return err
	}

	err := writeSessionRowAt(sp.File, sessInfo, id)
	if err != nil {
		return err
	}
	sp.ValidSessionsMap[id] = *sessInfo

	return nil
}

func (sp *SessionRepo) InvalidateSession(id string) error {
	log.Info("Invalidating session with id %s...", id)
	sessionInfo, exists := sp.ValidSessionsMap[id]
	if !exists {
		return fmt.Errorf("Session with id %s is already invalid or doesn't exist", id)
	}

	var err error
	sessionInfo.Valid = false
	err = writeSessionRowAt(sp.File, &sessionInfo, id)
	if err != nil {
		log.Error("Failed to invalidate session with id %s", id)
		return err
	}
	delete(sp.ValidSessionsMap, id)
	log.Info("Invalidated the session with id %s", id)

	return nil
}

func (sp *SessionRepo) FlushAllData() error {
	log.Error("FlushAllData method should probably be removed, buddy")

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

// Function verifySessionNotExists checks if there's a session with a given id that
// is already stored.
// The returned error is nil if there's no session with a given id.
// The function returns SessionExistsError instance if there's a session with a given id.
func (sp *SessionRepo) verifySessionNotExists(id string) error {
	log.Warn(
		"Only VALID sessions are looked up while verifying that a session with a certain id doesn't exist!",
	)
	_, exists := sp.ValidSessionsMap[id]
	if exists {
		return &customErrors.SessionExistsError{id}
	}

	return nil
}

func FormatToRow(sessInfo *SessionInfo, id string) ([]byte, error) {
	rowBuf := make([]byte, sessionCSVRowSize)

	idBytes := []byte(id)

	if len(id) > valueSizeId {
		return nil, fmt.Errorf(
			"Size of session id '%s' exceeds %d bytes! Can't format it to a csv row.",
			id, 
			valueSizeId,
		)
	}

	userIdBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(userIdBytes, uint64(sessInfo.UserId))

	createdAtBytes := []byte(sessInfo.CreatedAt.Format(timeLayout))
	expiresAtBytes := []byte(sessInfo.ExpiresAt.Format(timeLayout))
	validBytes := util.BoolToByteSlice(sessInfo.Valid)

	copy(rowBuf[columnOffsetId:], idBytes)
	rowBuf[columnOffsetId + valueSizeId] = util.AsciiComma

	copy(rowBuf[columnOffsetUserId:], userIdBytes)
	rowBuf[columnOffsetUserId + valueSizeUserId] = util.AsciiComma

	copy(rowBuf[columnOffsetCreatedAt:], createdAtBytes)
	rowBuf[columnOffsetCreatedAt + valueSizeCreatedAt] = util.AsciiComma

	copy(rowBuf[columnOffsetExpiresAt:], expiresAtBytes)
	rowBuf[columnOffsetExpiresAt + valueSizeExpiresAt] = util.AsciiComma

	copy(rowBuf[columnOffsetValid:], validBytes)
	rowBuf[columnOffsetValid + valueSizeValid] = util.AsciiNewLine

	return rowBuf, nil
}

func writeSessionRowAt(file *os.File, sessInfo *SessionInfo, id string) error {
	var err error
	rowBytes, err := FormatToRow(sessInfo, id)
	if err != nil {
		log.Error("Failed to format session info '%s' with id %s to a csv row.", sessInfo, id)
		return err
	}

	_, err = file.WriteAt(rowBytes, sessInfo.Offset) // TODO: Handle incomplete write!
	log.Warn(
		"Incomplete write case to the session csv file is not handled... " + 
		"Just saying, buddy... Juuust saying...",
	)
	if err != nil {
		log.Error(
			"Failed to write session info with id %s to the session csv file at offset %d",
			id,
			sessInfo.Offset,
		)
		return err
	}

	log.Info(
		"Successfully wrote session info '%s' with id %s to the session csv file at offset %d",
		sessInfo,
		id,
		sessInfo.Offset,
	)
	return nil
}
