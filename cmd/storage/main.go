package main

import (
	"fmt"
	"log"

	"github.com/chigvero/modules/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("hello"))
	if err != nil {
		log.Fatalf("%v", err)
	}
	restoredFile, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println("it is restored", *restoredFile)
}
