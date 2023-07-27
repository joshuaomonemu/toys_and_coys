package helper

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func Unmarshal(r []byte, p *map[string]interface{}) {
	err := json.Unmarshal(r, p)
	if err != nil {
		log.Fatalln(err)
	}
}

func Randoms() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	s := make([]int32, 5)
	err := binary.Read(r, binary.BigEndian, &s)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))
}
