package db

import (
	"os"
	"net/http"
	"fmt"
	"github.com/nturbo1/challenge-tracker-service/util"
)

type CookieRepo struct {
	file *os.File
	Cookies []*http.Cookie
}

var cookiesCSVHeader = "Name, Value, MaxAge, Secure, HttpOnly, Domain, SameSite"

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
	return nil, fmt.Errorf("FixMe: Implement parseCookiesCSVFile!")
}

func findCookie(cookieName string) *http.Cookie {
	fmt.Println("FixMe: Implement findCookie function in persistence!")
	return nil
}
