package helper

import (
	"encoding/json"
	"log"
)
// decoding the string
func Unmarshal(r []byte, p *map[string]interface{}) {
	err := json.Unmarshal(r, p)
	if err != nil {
		log.Fatalln(err)
	}
}
