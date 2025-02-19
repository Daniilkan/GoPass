package main

import (
	"GoPass/internal/DataBase"
	"GoPass/internal/Generating"
	"GoPass/internal/Vault"
	"fmt"
)

func main() {
	_ = database.InitDataBase()
	fmt.Println("Hello! It's Password Vault, you can choose write or read or clear password using 0/1/2 or type exit")
	for {
		choice := ""
		fmt.Scanln(&choice)
		if choice == "exit" {
			break
		} else {
			switch choice {
			case "0":
				fmt.Println("Please enter service and password in one line you can do -g in password to generate it")
				service := ""
				password := ""
				fmt.Scanln(&service, &password)
				if password == "-g" {
					password, _ = generating.GeneratePassword(16)
				}
				if len(password) == 0 || len(service) == 0 {
					fmt.Println("Enter valid service or password")
					continue
				}
				fmt.Println("Your password is added:", password)
				password, _ = vault.Encrypt(password, []byte("absolute cinemas"))
				_ = database.AddPassWord(service, password)
			case "1":
				fmt.Println("Please enter service in one line")
				service := ""
				fmt.Scanln(&service)
				pass, _ := database.GetPassWord(service)
				pass, _ = vault.Decrypt(pass, []byte("absolute cinemas"))
				if len(pass) == 0 {
					fmt.Println("Your password is empty")
				} else {
					fmt.Println("Password for", service, "is:", string(pass))
				}
			case "2":
				fmt.Println("Please enter service in one line")
				service := ""
				fmt.Scanln(&service)
				_ = database.ClearPassWord(service)
				fmt.Println("Your password is cleared")
			}
		}
	}
}
