package logger

import (
	"log"
	"encoding/json"
)

func LogError(err error) {
	output, _ :=json.Marshal(err)
	log.Println(string(output))
}