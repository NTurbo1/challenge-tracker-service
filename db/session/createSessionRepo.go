package session

import (
	"fmt"
	"os"
	"bufio"
	"time"
	"io"
	"encoding/binary"
	"strconv"

	"github.com/nturbo1/challenge-tracker-service/log"
	"github.com/nturbo1/challenge-tracker-service/util"
)

func CreateSessionRepo(filepath string) (*SessionRepo, error) {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Error("Failed to open session storage file: ", filepath)
		return nil, err
	}

	sessRepo, err := parseSessionsCSV(file)
	if err != nil {
		log.Error("Failed to parse the sessions csv file")
		return nil, err
	}

	return sessRepo, nil
}

func parseSessionsCSV(file *os.File) (*SessionRepo, error) {
	log.Info("Parsing sessions csv file...")
	if file == nil {
		return nil, fmt.Errorf("Can't parse nil file.")
	}

	var err error
	if err != nil {
		return nil, err
	}

	r := bufio.NewReaderSize(file, sessionCSVRowSize)
	validSessionsMap := make(map[string]SessionInfo)
	var rowIndex int64 = 0
	for {
		rowBytes, err := r.ReadBytes(util.AsciiNewLine)
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}

		id, sessInfoPtr, err := parseRowBytes(rowBytes, rowIndex)
		if err != nil {
			return nil, err
		}
		if sessInfoPtr.Valid {
			validSessionsMap[id] = *sessInfoPtr
		}
		rowIndex++
	}
	totalNumRows := rowIndex

	return &SessionRepo{
		file,
		totalNumRows,
		validSessionsMap,
	}, nil
}

// Function parseSessionInfo parses a given csv line bytes and returns a session id value and
// session info.
// The returned error is not nil, if there's a parsing error.
func parseRowBytes(rowBytes []byte, rowIndex int64) (string, *SessionInfo, error) {
	log.Debug("Parsing sessions csv row %d...", rowIndex + 1)
	rowBytesSize := len(rowBytes)
	if rowBytesSize < sessionCSVRowSize {
		return "", nil, fmt.Errorf(
			"Expected a sessions csv row to have length %d, but received %d bytes",
			sessionCSVRowSize,
			rowBytesSize,
		)
	}

	id := string(rowBytes[columnOffsetId:valueSizeId])
	log.Debug("Parsed id: %s", id)

	userIdBytes := rowBytes[columnOffsetUserId:(columnOffsetUserId + valueSizeUserId)]
	userId := int64(binary.BigEndian.Uint64(userIdBytes))
	log.Debug("Parsed userId: %d", userId)

	createdAtBytes := rowBytes[columnOffsetCreatedAt:(columnOffsetCreatedAt + valueSizeCreatedAt)]
	createdAt, err := time.Parse(timeLayout, string(createdAtBytes))
	if err != nil {
		log.Error(
			"Failed to parse 'createdAt' bytes string '%s' as %s format",
			createdAtBytes,
			timeLayout,
		)
		return "", nil, err
	}
	log.Debug("Parsed createdAt: %s", createdAt.Format(timeLayout))

	expiresAtBytes := rowBytes[columnOffsetExpiresAt:(columnOffsetExpiresAt + valueSizeExpiresAt)]
	expiresAt, err := time.Parse(timeLayout, string(expiresAtBytes))
	if err != nil {
		log.Error(
			"Failed to parse 'expiresAt' bytes string '%s' as %s format",
			expiresAtBytes,
			timeLayout,
		)
		return "", nil, err
	}
	log.Debug("Parsed expiresAt: %s", expiresAt.Format(timeLayout))

	validBytes := rowBytes[columnOffsetValid:(columnOffsetValid + valueSizeValid)]
	valid, err := util.BytesSliceToBool(validBytes)
	if err != nil {
		log.Error("Failed to parse 'valid' bytes slice %#v to a boolean", validBytes)
		return "", nil, err
	}
	log.Debug("Parsed valid: %s", strconv.FormatBool(valid))

	var offset int64 = rowIndex * sessionCSVRowSize
	log.Debug("Calulated offset: %d", offset)

	sessInfo := SessionInfo{
		UserId: userId,
		CreatedAt: createdAt,
		ExpiresAt: expiresAt,
		Valid: valid,
		Offset: offset,
	}

	return id, &sessInfo, nil
}
