package main

import (
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"io"
)

type User struct {
	email string
	password  string
}

func (u User) encryptPassword() string {
	hasher := md5.New()
	 _, err := io.WriteString(hasher, u.password)
	 if err != nil {
	  panic(err)
	 }
	return hex.EncodeToString(hasher.Sum(nil))
}

func registerUser() {
	user := User{
		email:    "",
		password: "",
	}
	defer fmt.Println("INFO: Berhasil register")
	fmt.Println("Masukan email untuk register")
	_, err := fmt.Scanln(&user.email)
	if err != nil {
		fmt.Println("Email tidak boleh kosong")
		return
	}
	
	fmt.Println("Masukan password untuk register")
	_, err = fmt.Scanln(&user.password)
	if err != nil {
		fmt.Println("Password tidak boleh kosong")
		return
	}
	
	fmt.Println("Email:", user.email)
	fmt.Println("Password:", user.encryptPassword())
}

func main() {
	registerUser()
}
