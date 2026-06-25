package main

import (
	"fmt"

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
	books, _ := storage.SearchBooksByQuery("")
	for _, book := range books {
		utils.DisplayLine(book)
	}
	// storage.AppendBook(utils.FormatBookLine(16, "Шата", "Ри Гува", 2024, true))
}
