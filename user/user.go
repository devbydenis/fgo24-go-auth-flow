package user

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

type User struct {
	ID string
	Email string
	Password string
}

type UserLogin struct {
	Email string
	Password string
}

type UserList struct {
	UserList []User
}

func (ul *UserList) HandleLoginUser(userLogin UserLogin) {
	for _, user := range ul.UserList {
		if user.Email == userLogin.Email && user.Password == userLogin.Password {
			fmt.Println("Login Success", userLogin)
		} else {
			fmt.Println("Login Failed. Please Check your email and password")
		}
	}	
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

func (ul UserLogin) EncryptPassword() string {
	hasher := md5.New()
	 _, err := io.WriteString(hasher, ul.Password)
	 if err != nil {
	  panic(err)
	 }
	return hex.EncodeToString(hasher.Sum(nil))
}
