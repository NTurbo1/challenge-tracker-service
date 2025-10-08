package user

import (
	"bufio"
	"os"
	"fmt"
	"strings"

	"github.com/nturbo1/challenge-tracker-service/util"
)

func CreateUserRepo(usersFilepath string) (*UserRepo, error) {
	fileExists := util.FileExists(usersFilepath)

	file, err := os.OpenFile(usersFilepath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Failed to open users storage file ", usersFilepath)
		return nil, err
	}

	if fileExists {
		users, err := parseUsersCSVFile(file)
		if err != nil {
			fmt.Println("Failed to parse the users csv file")
			return nil, err
		}

		return &UserRepo{file, users}, nil
	}

	err = util.WriteHeaderLineTo(file, usersCSVHeader)
	if err != nil {
		fmt.Println("Failed to write the header to the users csv file.")
		return nil, err
	}

	return &UserRepo{file, []*User{}}, nil
}

func parseUsersCSVFile(file *os.File) ([]*User, error) {
	fmt.Println("Parsing users csv file...")
	if file == nil {
		return nil, fmt.Errorf("Can't parse nil file.")
	}

	var users = []*User{}
	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++;

		if lineCount == 1 {
			fmt.Println("Line 1 (Header): ", line)
			if strings.Compare(usersCSVHeader, line) != 0 {
				return nil, fmt.Errorf(
					"Invalid users csv header: %s. Expected to be: %s", line, usersCSVHeader,
				)
			}
			continue
		}

		fmt.Printf("Line %d: %s\n", lineCount, line)
		userPtr, err := parseUserFromCSVLine(line)
		if err != nil {
			fmt.Println("Failed to parse line: ", line)
			return nil, err
		}
		fmt.Println("Parsed to: ", *userPtr)
		users = append(users, userPtr)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func parseUserFromCSVLine(CSVLine string) (*User, error) {
	if len(CSVLine) == 0 {
		return nil, fmt.Errorf("Empty csv line.")
	}
	cols := strings.Split(CSVLine, ",")
	numCols := len(cols)
	user := User{
		Firstname: cols[0],
	}

	if numCols > 1 {
		user.Lastname = cols[1]
		if numCols > 2 {
			user.Username = cols[2]
			if numCols > 3 {
				user.Password = cols[3]
			}
		}
	}

	return &user, nil
}
