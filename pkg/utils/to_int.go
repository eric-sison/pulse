package utils

import (
	"log"
	"strconv"
)

func ToInt(strVal string) int {
	intVal, err := strconv.Atoi(strVal)

	if err != nil {
		log.Fatal(err)
	}

	return intVal
}
