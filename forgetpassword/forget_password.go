package forgetpassword

import (
	"auth_flow/user"
	"fmt"
	"os"
)

// var UserList user.UserList

func ForgetPassword(users *user.UserList) user.UserForgetPassword {
	fmt.Println("current users:", users)
	userForgetPassword := user.UserForgetPassword {
		Email: "",
		Password: "",
	} 
	
	// Input email forget password
	fmt.Println("Masukan email untuk ganti password baru")
	_, err := fmt.Scanln(&userForgetPassword.Email)	
	if err != nil {
		fmt.Println("Email tidak boleh kosong")
		os.Exit(0)
	}
	emailFound := false
	for _, user := range users.UserList {
		if userForgetPassword.Email == user.Email {
			emailFound = true
			break
		}
	}
	if !emailFound {
		fmt.Println("Email tidak ditemukan")
		os.Exit(0)
	}
	
	// Input new password forget password
	fmt.Println("Masukan password baru!")
	_, err = fmt.Scanln(&userForgetPassword.Password)
	if err != nil {
		fmt.Println("Password tidak boleh kosong")
		os.Exit(0)
	}
	
	fmt.Println("Masukan konfirmasi password baru!")
	var confirm string
	_, err = fmt.Scanln(&confirm)
	if err != nil {
		fmt.Println("Konfirmasi password tidak boleh kosong")
		os.Exit(0)
	}
	if userForgetPassword.Password != confirm {
		fmt.Println("Password tidak cocok")
		os.Exit(0)
	}
	
	userForgetPassword.Password = userForgetPassword.EncryptPassword()
	
	return userForgetPassword
}