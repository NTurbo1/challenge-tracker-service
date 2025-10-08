package user

import (
	"fmt"
	"os"
)

type User struct {
	Firstname string
	Lastname string
	Username string
	Password string
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
	fmt.Println("Closing the user repo...")
	err := ur.File.Close()
	if err != nil {
		fmt.Println("While closing the user repo file: ", err)
	}

	return err
}
