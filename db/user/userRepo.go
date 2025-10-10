package user

import (
	"fmt"
	"os"
	"strconv"

	"github.com/nturbo1/challenge-tracker-service/log"
)

type User struct {
	Id int
	Firstname string
	Lastname string
	Username string
	Password string
}
func (u *User) String() string {
	return "{id: " + strconv.FormatInt(int64(u.Id), 10) + ", firstname: " + u.Firstname +
	", lastname: " + u.Lastname + ", username: " + u.Username + ", password: " + u.Password +
	"}" // NOTE: Careful with printing out the password!!!
}

type UserRepo struct {
	File *os.File
	UsersMap map[string]*User
}

// Function FindByUsername searches for a user with a given username and returns
// nil if not found.
func (ur *UserRepo) FindByUsername(username string) *User {
	if userPtr, exists := ur.UsersMap[username]; exists {
		return userPtr
	}

	return nil
}

func (ur *UserRepo) FlushAllData() error {
	return fmt.Errorf("FixMe: Implement (*UserRepo).FlushAllData()!!!")
}

func (ur *UserRepo) Close() error {
	log.Info("Closing the user repo...")
	err := ur.File.Close()
	if err != nil {
		log.Error("While closing the user repo file: ", err)
	}

	return err
}
