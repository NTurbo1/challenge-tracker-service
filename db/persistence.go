package db

var db_files []string = []string{ "users.csv", "sessions.csv" }

var UsersRepository *UserRepo = nil
var SessionRepository *SessionRepo = nil

const usersCSVFilePath = "db/data/users.csv"
const sessionsCSVFilePath = "db/data/sessions.csv"

func InitDb() error {
	var err error
	UsersRepository, err = CreateUserRepo(usersCSVFilePath)
	if err != nil {
		return err
	}

	SessionRepository, err = CreateSessionRepo(sessionsCSVFilePath)
	if err != nil {
		return err
	}

	return nil
}
