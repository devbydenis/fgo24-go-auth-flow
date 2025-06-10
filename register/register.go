package register

import (
	"fmt"
	"os"
	"auth_flow/user"
	"crypto/rand"
)

func generateID() string {
	id := make([]byte, 8)
	_, err := rand.Read(id)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", id)
}

func RegisterUser() user.User {
	user := user.User  {
		ID:       generateID(),
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
	
	user.Password = user.EncryptPassword()
	
	return user
}
