package user

import (
	"fmt"
	"os"
	"strings"
)

type User struct {
	Firstname string
	Lastname string
	Username string
	Password string
}

type UserRepo struct {
	File *os.File
	Users []*User
}

func (ur *UserRepo) FindUserByUsername(username string) (*User, error) {
	users := ur.Users
	if len(users) == 0 {
		return nil, fmt.Errorf("User not found by username %s", username)
	}

	for i := 0; i < len(users); i++ {
		user := users[i]
		if strings.Compare(user.Username, username) == 0 {
			return user, nil
		}
	}

	return nil, fmt.Errorf("User not found by username %s", username)
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
