package data

import (
	"github.com/maprost/pqx"
	"github.com/maprost/pqx/pqarg"
)

type UserAccount struct {
	Email        string
	PasswordHash string
}

func CreateUser() error {
	return pqx.Create(UserAccount{})
}

func (u *UserAccount) Insert() error {
	return pqx.Insert(u)
}

func DeleteUser(email string) error {
	args := pqarg.New()
	result, err := pqx.Query(
		"DELETE FROM UserAccount "+
			" WHERE email = "+args.Next(email), args)
	if err != nil {
		err
	}
	result.Close()
	return nil
}

func GetUsers() ([]UserAccount, error) {
	users := []UserAccount{}
	prototype := UserAccount{}

	err := pqx.SelectList(&prototype, func() {
		users = append(users, prototype)
	})

	return users, err
}
