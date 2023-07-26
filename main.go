package main

import (
	"fmt"

	"github.com/google/uuid"
)

//uuid

func main() {
	uuidObj, _ := uuid.NewUUID()
	fmt.Println(" ", uuidObj.String())

	uuidObj2, _ := uuid.NewUUID()
	fmt.Println(" ", uuidObj2.String())
}

// $ go run main.go
//   9cb85d7e-2ba2-11ee-b1a6-16d4241b0fa5
//   9cb9b277-2ba2-11ee-b1a6-16d4241b0fa5

