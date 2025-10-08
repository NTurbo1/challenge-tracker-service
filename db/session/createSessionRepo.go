package session

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"bufio"
	"time"

	"github.com/nturbo1/challenge-tracker-service/util"
)

func CreateSessionRepo(filepath string) (*SessionRepo, error) {
	fileExists := util.FileExists(filepath)

	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Failed to open file: ", filepath)
		return nil, err
	}

	if fileExists {
		sessionsMap, err := parseSessionsMapFromCSVFile(file)
		if err != nil {
			return nil, err
		}

		return &SessionRepo{
			File: file,
			SessionsMap: sessionsMap,
		}, nil
	}

	return &SessionRepo{
		File: file,
		SessionsMap: map[string]SessionInfo{},
	}, nil
}

func parseSessionsMapFromCSVFile(file *os.File) (map[string]SessionInfo, error) {
	scanner := bufio.NewScanner(file)
	lineCount := 0
	sessionsMap := map[string]SessionInfo{}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			break;
		}

		lineCount++

		if lineCount == 1 {
			if strings.Compare(line, sessionCSVHeader) != 0 {
				return nil, fmt.Errorf(
					"Invalid session csv file header: %s. Expected: %s", line, sessionCSVHeader,
				)
			}
			continue
		}

		id, sessionInfoPtr, err := parseSessionInfo(line)
		if err != nil {
			return nil, err
		}
		sessionsMap[id] = *sessionInfoPtr
	}

	return sessionsMap, nil
}

// Function parseSessionInfo parses a given csv line and returns a session id value and
// session info.
// The return error is not nil, if there's a parsing error.
func parseSessionInfo(csvLine string) (string, *SessionInfo, error) {
	if len(csvLine) == 0 {
		return "", nil, fmt.Errorf("Empty csv line.")
	}
	
	cols := strings.Split(csvLine, ",")
	if len(cols) != numSessionCSVCols {
		return "", nil, fmt.Errorf("Invalid session csv line: %s", csvLine)
	}

	userId, err := strconv.Atoi(cols[1])
	if err != nil {
		fmt.Printf(
			"Failed to convert userId value '%s' from session csv file row '%s' to an integer.\n", 
			cols[1], csvLine,
		)
		return "", nil, err
	}

	createdAt, err := time.Parse(timeLayout, cols[2])
	if err != nil {
		fmt.Printf(
			"Failed to parse createdAt value '%s' from session csv file row '%s' to layout '%s'.\n", 
			cols[2], csvLine, timeLayout,
		)
		return "", nil, err
	}

	expiresAt, err := time.Parse(timeLayout, cols[3])
	if err != nil {
		fmt.Printf(
			"Failed to parse expiresAt value '%s' from session csv file row '%s' to layout '%s'.\n",
			cols[3], csvLine, timeLayout,
		)
		return "", nil, err
	}

	return cols[0], &SessionInfo{userId, createdAt, expiresAt}, nil
}
