package main

import (
	"auth_flow/login"
	"auth_flow/register"
	"auth_flow/user"
	"fmt"
)
	var UserList user.UserList

func main() {
	
	// fmt.Println(UserList)
	for {
		fmt.Println("Halo! silakan masukan perintah dengan angka!")
		fmt.Println("Silakan pilih menu berikut:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Forget Password")
		fmt.Println("4. Exit")
		
		var input int
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Input tidak valid. Masukin input yang valid!")
			return
		}
		
		switch choice := input; choice {
		case 1:
			UserList.HandleLoginUser(login.LoginUser())
		case 2:
			UserList.AddUser(register.RegisterUser())
			fmt.Println("Userlist saat ini: ", UserList)
		case 3:
			fmt.Println("Forget Password")
		case 4:
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Pilihan hanya ada 1 - 4. Pilihan lu orang ga valid!")
			return
		}
	
	}
}
