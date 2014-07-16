package uuid

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
)

func GenerateUuid() {
	uuid := uuid.NewRandom()
	fmt.Println("Generated uuid:", uuid)
}
