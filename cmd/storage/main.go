package main

import (
	"fmt"
	"github.com/chigvero/modules/internal/storage"
)

func main() {
	st := storage.NewStorage()

	fmt.Println("it works", st)
}
