package uuid

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
)

func GenerateUuid() uuid.UUID {
	uuid := uuid.NewRandom()
	fmt.Println("Generated uuid:", uuid)
	return uuid
}
