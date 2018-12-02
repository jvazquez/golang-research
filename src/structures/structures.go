package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello world")
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning! %s", t)
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening. ", t.Hour())
	}
}
