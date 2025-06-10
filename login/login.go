package login

import (
	"auth_flow/user"
	"fmt"
	"os"
)

func LoginUser() user.UserLogin {
	userLogin := user.UserLogin  {
		Email:    "",
		Password: "",
	}

	// defer fmt.Println("INFO: Login Berhasil")
	
	fmt.Println("Masukan email untuk login")
	_, err := fmt.Scanln(&userLogin.Email)
	if err != nil {
		fmt.Println("Email tidak boleh kosong")
		os.Exit(0)
	}
	
	fmt.Println("Masukan password untuk login")
	_, err = fmt.Scanln(&userLogin.Password)
	if err != nil {
		fmt.Println("Password tidak boleh kosong")
		os.Exit(0)
	}
	
	userLogin.Password = userLogin.EncryptPassword()
	
	return userLogin
}
