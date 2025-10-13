package session

import (
	"fmt"
	"os"
	"bufio"
	"time"
	"io"
	"encoding/binary"

	"github.com/nturbo1/challenge-tracker-service/log"
	"github.com/nturbo1/challenge-tracker-service/util"
)

func CreateSessionRepo(filepath string) (*SessionRepo, error) {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Error("Failed to open session storage file: ", filepath)
		return nil, err
	}

	validSessionsMap, err := parseSessionsCSVForValidSessions(file)
	if err != nil {
		log.Error("Failed to parse the sessions csv file")
		return nil, err
	}

	return &SessionRepo{file, validSessionsMap}, nil
}

func parseSessionsCSVForValidSessions(file *os.File) (map[string]SessionInfo, error) {
	log.Info("Parsing sessions csv file...")
	if file == nil {
		return nil, fmt.Errorf("Can't parse nil file.")
	}

	var err error
	readerBufferSize, err := util.Max(sessionCSVRowSize, len(sessionCSVHeader))
	if err != nil {
		return nil, err
	}

	r := bufio.NewReaderSize(file, readerBufferSize)
	validSessionsMap := make(map[string]SessionInfo)
	for {
		rowBytes, err := r.ReadBytes(util.AsciiNewLine)
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}

		id, sessInfoPtr, err := parseRowBytes(rowBytes)
		if err != nil {
			return nil, err
		}
		validSessionsMap[id] = *sessInfoPtr
	}

	return validSessionsMap, nil
}

// Function parseSessionInfo parses a given csv line bytes and returns a session id value and
// session info.
// The returned error is not nil, if there's a parsing error.
func parseRowBytes(rowBytes []byte) (string, *SessionInfo, error) {
	rowBytesSize := len(rowBytes)
	if rowBytesSize < sessionCSVRowSize {
		return "", nil, fmt.Errorf(
			"Expected a sessions csv row to have length %d, but received %d bytes",
			sessionCSVRowSize,
			rowBytesSize,
		)
	}

	id := string(rowBytes[columnOffsetId:valueSizeId])
	userId := int64(binary.BigEndian.Uint64(rowBytes[columnOffsetUserId:(columnOffsetUserId + valueSizeUserId)]))

	createdAtBytesStr := string(rowBytes[columnOffsetCreatedAt:(columnOffsetCreatedAt + valueSizeCreatedAt)])
	createdAt, err := time.Parse(timeLayout, createdAtBytesStr)
	if err != nil {
		log.Error(
			"Failed to parse 'createdAt' bytes string '%s' as %s format",
			createdAtBytesStr,
			timeLayout,
		)
		return "", nil, err
	}

	expiresAtBytesStr := string(rowBytes[columnOffsetExpiresAt:(columnOffsetExpiresAt + valueSizeExpiresAt)])
	expiresAt, err := time.Parse(timeLayout, expiresAtBytesStr)
	if err != nil {
		log.Error(
			"Failed to parse 'expiresAt' bytes string '%s' as %s format",
			expiresAtBytesStr,
			timeLayout,
		)
		return "", nil, err
	}

	validBytes := rowBytes[columnOffsetValid:(columnOffsetValid + valueSizeValid)]
	valid, err := util.BytesSliceToBool(validBytes)
	if err != nil {
		log.Error("Failed to parse 'valid' bytes slice %#v to a boolean", validBytes)
		return "", nil, err
	}

	offset := int64(binary.BigEndian.Uint64(rowBytes[columnOffsetOffset:(columnOffsetOffset + valueSizeOffset)]))

	sessInfo := SessionInfo{
		UserId: userId,
		CreatedAt: createdAt,
		ExpiresAt: expiresAt,
		Valid: valid,
		Offset: offset,
	}

	return id, &sessInfo, nil
}
