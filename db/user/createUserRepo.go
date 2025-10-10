package user

import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"strconv"

	"github.com/nturbo1/challenge-tracker-service/log"
)

func CreateUserRepo(usersFilepath string) (*UserRepo, error) {
	file, err := os.OpenFile(usersFilepath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Error("Failed to open users storage file ", usersFilepath)
		return nil, err
	}

	usersMap, err := parseUsersCSVFile(file)
	if err != nil {
		log.Error("Failed to parse the users csv file")
		return nil, err
	}

	return &UserRepo{file, usersMap}, nil
}

func parseUsersCSVFile(file *os.File) (map[string]*User, error) {
	log.Info("Parsing users csv file...")
	if file == nil {
		return nil, fmt.Errorf("Can't parse nil file.")
	}

	var usersMap = map[string]*User{}
	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			break;
		}

		lineCount++;

		if lineCount == 1 {
			log.Debug("Line 1 (Header): %s", line)
			if strings.Compare(usersCSVHeader, line) != 0 {
				return nil, fmt.Errorf(
					"Invalid users csv header: %s. Expected to be: %s", line, usersCSVHeader,
				)
			}
			continue
		}

		log.Debug("Line %d: %s\n", lineCount, line)
		userPtr, err := parseUserFromCSVLine(line)
		if err != nil {
			log.Error("Failed to parse line: %s", line)
			return nil, err
		}
		log.Debug("Parsed to: %s", userPtr)
		usersMap[userPtr.Username] = userPtr
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return usersMap, nil
}

func parseUserFromCSVLine(CSVLine string) (*User, error) {
	if len(CSVLine) == 0 {
		return nil, fmt.Errorf("Empty csv line.")
	}

	cols := strings.Split(CSVLine, ",")
	numCols := len(cols)
	if numCols != numUsersCSVHeaderCols {
		return nil, fmt.Errorf(
			"Line '%s' has %d columns but expected to have %d columns.", 
			CSVLine, numCols, numUsersCSVHeaderCols,
		)
	}

	userId, err := strconv.Atoi(cols[0])
	if err != nil {
		log.Error("Failed to convert user id %s to an integer.", cols[0])
		return nil, err
	}

	user := User{
		Id: userId,
		Username: cols[1],
		Firstname: cols[2],
		Lastname: cols[3],
		Password: cols[4],
	}

	return &user, nil
}
