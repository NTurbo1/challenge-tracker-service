package db

import (
	"github.com/nturbo1/challenge-tracker-service/db/session"
	"github.com/nturbo1/challenge-tracker-service/db/user"
	"github.com/nturbo1/challenge-tracker-service/log"
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
	log.Info("Initializing the database...")
	var err error
	UserRepository, err = user.CreateUserRepo(usersCSVFilePath)
	if err != nil {
		return err
	}

	SessionRepository, err = session.CreateSessionRepo(sessionsCSVFilePath)
	if err != nil {
		return err
	}

	log.Info("The database has been initialized!")
	return nil
}

func Flush() {
	log.Info("Flushing all data :0 ...")
	for _, repo := range repos {
		err := repo.FlushAllData()
		if err != nil {
			log.Error(
				"FixMe: FUCK!!! Failed to flush all data for a repository! You gotta " +
				"figure out something to handle this situation like a real man for real, man!" +
				"Now, you're gonna loose all of your fucking DATA!!!\n\nBTW, the error is: ", 
				err,
			)
		}
	}

	log.Info("Flushed all data successfully! :)")
}

func Close() {
	for _, repo := range repos {
		repo.Close()
	}
}
