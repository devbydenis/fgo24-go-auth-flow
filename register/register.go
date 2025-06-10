package register

import (
	"fmt"
	"os"
	"auth_flow/user"
)

func RegisterUser() user.User {
	user := user.User  {
		ID:       0,
		Email:    "",
		Password: "",
	}
	defer fmt.Println("INFO: Berhasil register")
	fmt.Println("Masukan email untuk register")
	_, err := fmt.Scanln(&user.Email)
	if err != nil {
		fmt.Println("Email tidak boleh kosong")
		os.Exit(0)
	}
	
	fmt.Println("Masukan password untuk register")
	_, err = fmt.Scanln(&user.Password)
	if err != nil {
		fmt.Println("Password tidak boleh kosong")
		os.Exit(0)
	}
	
	fmt.Println("Masukan konfirmasi password")
	var confirm string
	_, err = fmt.Scanln(&confirm)
	if err != nil {
		fmt.Println("Konfirmasi password tidak boleh kosong")
		os.Exit(0)
	}
	
	if user.Password != confirm {
		fmt.Println("Password tidak cocok")
		os.Exit(0)
	}
	
	fmt.Println("INFO: Berhasil register")
	user.Password = user.EncryptPassword()
	// fmt.Println(user)
	return user
}
