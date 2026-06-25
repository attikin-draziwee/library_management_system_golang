package main

import (
	"fmt"

	"github.com/attikin-draziwee/library_management_system_golang/storage"
	"github.com/attikin-draziwee/library_management_system_golang/utils"
)

func main() {
	fmt.Println("Hello, World!")
	if err := storage.InitStorage(); err != nil {
		utils.Display(err)
	}
}
