package utils

import (
	"fmt"

	"github.com/google/uuid"
)

func generateNewUUID() (uuid.UUID, error) {
	uuid, err := uuid.NewRandom()
	return uuid, err
}

//GetNewUUID generates new random UUID string
func GetNewUUID() string {
	uuid, err := generateNewUUID()

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return uuid.String()
}
