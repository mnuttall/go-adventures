package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("hello, time is " + time.Now().String())
		time.Sleep(time.Millisecond * 5000)
	}
}
