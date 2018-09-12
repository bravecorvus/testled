package main

import (
	"fmt"
	"time"

	rpio "github.com/stianeikeland/go-rpio"
)

func main() {
	err := rpio.Open()
	if err != nil {
		fmt.Println("failed to rpio.Open()")
		fmt.Println(err.Error())
	}
	pin := rpio.Pin(17)
	pin.Output()
	for {
		time.Sleep(1 * time.Second)
		pin.Toggle()
	}
}
