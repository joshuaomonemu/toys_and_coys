package helper

import (
	"encoding/json"
	"log"
)

func Unmarshal(r []byte, p *map[string]interface{}) {
	err := json.Unmarshal(r, p)
	if err != nil {
		log.Fatalln(err)
	}
}
