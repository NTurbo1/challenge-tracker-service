package db

import (
	"fmt"
	
	"github.com/nturbo1/challenge-tracker-service/db/session"
	"github.com/nturbo1/challenge-tracker-service/db/user"
)

var db_files []string = []string{ "users.csv", "sessions.csv" }

var UserRepository *user.UserRepo = nil
var SessionRepository *session.SessionRepo = nil

const usersCSVFilePath = "db/data/users.csv"
const sessionsCSVFilePath = "db/data/sessions.csv"

type Repo interface {
	FlushAllData() error
	Close() error
}

var repos []Repo = []Repo{UserRepository, SessionRepository}

func InitDb() error {
	fmt.Println("Initializing the database...")
	var err error
	UserRepository, err = user.CreateUserRepo(usersCSVFilePath)
	if err != nil {
		return err
	}

	SessionRepository, err = session.CreateSessionRepo(sessionsCSVFilePath)
	if err != nil {
		return err
	}

	fmt.Println("The database has been initialized!")
	return nil
}

func Flush() {
	fmt.Println("Flushing all data :0 ...")
	for _, repo := range repos {
		err := repo.FlushAllData()
		if err != nil {
			fmt.Println(
				"FixMe: FUCK!!! Failed to flush all data for a repository! You gotta " +
				"figure out something to handle this situation like a real man for real, man!" +
				"Now, you're gonna loose all of your fucking DATA!!!\n\nBTW, the error is: ", 
				err,
			)
		}
	}

	fmt.Println("Flushed all data successfully! :)")
}

func Close() {
	for _, repo := range repos {
		repo.Close()
	}
}
