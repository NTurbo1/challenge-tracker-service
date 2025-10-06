package db

import (
	"os"
	"strings"
	"net/http"
	"fmt"
	"bufio"
	"strconv"
	"github.com/nturbo1/challenge-tracker-service/util"
)

type CookieRepo struct {
	file *os.File
	Cookies []*http.Cookie
}

const cookiesCSVHeader = "Name, Value, MaxAge, Secure, HttpOnly, Domain"
const numCookiesCSVCols = 6

func CreateCookieRepo(cookiesFilePath string) (*CookieRepo, error) {
	fileExists := util.FileExists(cookiesFilePath)

	file, err := os.OpenFile(cookiesFilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Failed to open cookies storage file.")
		return nil, err
	}

	if fileExists {
		cookies, err := parseCookiesCSVFile(file)
		if err != nil {
			fmt.Println("Failed to parse the cookies csv file.")
			return nil, err
		}

		return &CookieRepo{file, cookies}, nil
	}

	err = util.WriteHeaderLineTo(file, cookiesCSVHeader)
	if err != nil {
		fmt.Println("Failed to write the header to the cookies csv file.")
		return nil, err
	}

	return &CookieRepo{file, []*http.Cookie{}}, nil
}

func parseCookiesCSVFile(file *os.File) ([]*http.Cookie, error) {
	if file == nil {
		return nil, fmt.Errorf("Can't parse nil file.")
	}

	cookies := []*http.Cookie{}
	lineCount := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineCount++

		if lineCount == 1 {
			if strings.Compare(cookiesCSVHeader, line) != 0 {
				return nil, fmt.Errorf(
					"Invalid cookies csv header: %s. Expected to be: %s", line, cookiesCSVHeader,
				)
			}
			fmt.Printf("Line %d (Header): %s\n", lineCount, line)
			continue
		}

		fmt.Printf("Line %d: %s\n", lineCount, line)
		cookiePtr, err := parseCookieCSVLine(line)
		if err != nil {
			return nil, err
		}
		fmt.Println("Parsed to: ", *cookiePtr)
		cookies = append(cookies, cookiePtr)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return cookies, nil
}

func parseCookieCSVLine(csvLine string) (*http.Cookie, error) {
	if len(csvLine) == 0 {
		return nil, fmt.Errorf("Empty line in cookies csv file.")
	}
	cols := strings.Split(csvLine, ",")
	if len(cols) != numCookiesCSVCols {
		return nil, fmt.Errorf("Invalid row in the cookies csv file: %s", csvLine)
	}

	maxAge, err := strconv.Atoi(cols[2])
	if err != nil {
		fmt.Printf(
			"Failed to convert cookie maxAge value %s to integer from line %s\n", cols[2], csvLine,
		)
		return nil, err
	}
	secure, err := strconv.ParseBool(cols[3])
	if err != nil {
		fmt.Printf(
			"Failed to convert cookie secure value %s to boolean from line %s\n", cols[3], csvLine,
		)
		return nil, err
	}
	httpOnly, err := strconv.ParseBool(cols[4])
	if err != nil {
		fmt.Printf(
			"Failed to convert cookie httpOnly value %s to boolean from line %s\n", cols[4], csvLine,
		)
		return nil, err
	}

	cookiePtr := &http.Cookie{
		Name: cols[0],
		Value: cols[1],
		MaxAge: maxAge,
		Secure: secure,
		HttpOnly: httpOnly,
		Domain: cols[5],
	}

	return cookiePtr, nil
}

func findCookie(cookieName string) *http.Cookie {
	fmt.Println("FixMe: Implement findCookie function in persistence!")
	return nil
}
