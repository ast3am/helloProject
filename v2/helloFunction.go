package v2

import (
	"fmt"
	"time"
)

func SayVersion() {
	fmt.Println("hello from version 2.0.1")
}

func SayTime() {
	fmt.Println(time.Now())
}
