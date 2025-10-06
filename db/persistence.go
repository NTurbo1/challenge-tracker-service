package db

var db_files []string = []string{ "users.csv", "cookies.csv" }

var UsersRepository *UserRepo = nil
var CookieRepository *CookieRepo = nil

const usersCSVFilePath = "db/data/users.csv"

func InitDb() error {
	var err error
	UsersRepository, err = CreateUserRepo(usersCSVFilePath)

	if err != nil {
		return err
	}

	return nil
}
