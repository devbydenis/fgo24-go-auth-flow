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

type UserList struct {
	UserList []User
}

type UserLogin struct {
	Email string
	Password string
}

type UserForgetPassword struct {
	Email string
	Password string
}

func changePassword(users *User, newPassword string) {
		users.Password = newPassword
	}
// func (ul *UserList) HandleForgetPassword(userForget UserForgetPassword) {
	// 	for _, user := range ul.UserList {
	// 		if user.Email == userForget.Email {
	// 			fmt.Println("Password Match", user.Email == userForget.Email)
	// 			fmt.Println(user)
	// 			user.Password = userForget.Password
	// 			fmt.Println("Password Changed")
	// 			break
	// 		} else {
	// 			fmt.Println("Email Not Found, Please register first!")
	// 		}
	// 	}
	// }

	
func (ul *UserList) ChangePassword(forgetData UserForgetPassword) {
	for _, user := range ul.UserList {
		if user.Email == forgetData.Email {
			fmt.Println("Password Match", user.Email == forgetData.Email)
			fmt.Println(user)
			fmt.Println(forgetData)
			user.Password = forgetData.Password
			fmt.Println("after change password", user)
			break
		} else {
			fmt.Println("Email Not Found, Please register first!")
		}
	}
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

func (ul UserForgetPassword) EncryptPassword() string {
	hasher := md5.New()
	 _, err := io.WriteString(hasher, ul.Password)
	 if err != nil {
	  panic(err)
	 }
	return hex.EncodeToString(hasher.Sum(nil))
}

