package main

import (
	"fmt"

	"github.com/attikin-draziwee/library_management_system_golang/service"
	"github.com/attikin-draziwee/library_management_system_golang/storage"
	"github.com/attikin-draziwee/library_management_system_golang/utils"
)

func main() {
	fmt.Println("Hello, World!")
	if err := storage.InitStorage(); err != nil {
		utils.DisplayLine(err)
	}

	// err := storage.RemoveBook(9)
	// if err != nil {
	// 	utils.DisplayLine(err)
	// }
	// err := storage.UpdateBookStatusInFile(16, "Нет")
	// utils.DisplayLine(err)

	// books, err := storage.ReadAllBooks()
	// if err != nil {
	// 	utils.DisplayLine(err)
	// }
	// for _, book := range books {
	// 	utils.DisplayLine(book)
	// }

	// utils.Display(storage.FindBookLineByID(5))
	// utils.DisplayLine(storage.GetBookInfo(7))
	// books, _ := storage.SearchBooksByQuery("")
	// for _, book := range books {
	// 	utils.DisplayLine(book)
	// }
	// storage.AppendBook(utils.FormatBookLine(16, "Шата", "Ри Гува", 2024, true))
	// users, err := storage.SearchUsersByQuery("Никита")
	// if err != nil {
	// 	utils.DisplayLine(err)
	// }
	// for _, user := range users {
	// 	utils.DisplayLine(user)
	// }
	// utils.Display(service.GenerateNextUserID())
	// utils.Display(service.GenerateNextUserID())
	// utils.Display(service.GenerateNextUserID())

	// err := service.AddUser("Джон", "john228eax.you")
	// fmt.Println(err)
	// service.ListUsers()
	// service.SearchUsers("ник")
	// exists, name, err := service.UserExists(19)
	// utils.DisplayLine(exists, name, err)
	// service.DeleteUser(3)
	// fmt.Println(service.BorrowBook(10, 11))
	// service.ListBook()
	service.ListUserBooks(10)
}
