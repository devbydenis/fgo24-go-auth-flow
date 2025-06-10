package user

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

type User struct {
	ID int
	Email string
	Password string
}

type UserList struct {
	UserList []User
}

func (ul *UserList) AddUser(user User) {
	ul.UserList = append(ul.UserList, user)
}

func (u User) EncryptPassword() string {
	hasher := md5.New()
	 _, err := io.WriteString(hasher, u.Password)
	 if err != nil {
	  panic(err)
	 }
	return hex.EncodeToString(hasher.Sum(nil))
}

