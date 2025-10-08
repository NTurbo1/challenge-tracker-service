package session

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"bufio"
	"time"
)

func CreateSessionRepo(filepath string) (*SessionRepo, error) {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Failed to open session storage file: ", filepath)
		return nil, err
	}

	sessionsMap, err := parseSessionsMapFromCSVFile(file)
	if err != nil {
		fmt.Println("Failed to parse the sessions csv file")
		return nil, err
	}

	return &SessionRepo{file, sessionsMap}, nil
}

func parseSessionsMapFromCSVFile(file *os.File) (map[string]SessionInfo, error) {
	fmt.Println("Parsing sessions csv file...")
	if file == nil {
		return nil, fmt.Errorf("Can't parse nil file.")
	}

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
			fmt.Println("Line 1 (Header): ", line)
			if strings.Compare(line, sessionCSVHeader) != 0 {
				return nil, fmt.Errorf(
					"Invalid session csv file header: %s. Expected: %s", line, sessionCSVHeader,
				)
			}
			continue
		}

		fmt.Printf("Line %d: %s\n", lineCount, line)
		id, sessionInfoPtr, err := parseSessionInfo(line)
		if err != nil {
			fmt.Println("Failed to parse line: ", line)
			return nil, err
		}
		fmt.Println("Parsed to: " + id + "," + sessionInfoPtr.String())
		sessionsMap[id] = *sessionInfoPtr
	}

	if err := scanner.Err(); err != nil {
		return nil, err
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
