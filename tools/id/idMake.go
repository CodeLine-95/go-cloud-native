package id

import (
	"fmt"
	"math/rand"
	"time"
)

var Make Maker

func init() {
	rand.Seed(time.Now().UnixNano())
	Make = NewMaker(rand.Uint32())
}

func GetStringID() string {
	return fmt.Sprintf("%v", Make.Make())
}
