package main

import (
	"fmt"
	"time"

	rpio "github.com/stianeikeland/go-rpio"
)

// func main() {
// err := rpio.Open()
// if err != nil {
// fmt.Println("failed to rpio.Open()")
// fmt.Println(err.Error())
// }
// defer rpio.Close()
// inputpin := rpio.Pin(27)
// outputpin := rpio.Pin(17)
//
// inputpin.Input()
// outputpin.Output()
//
// for {
// time.Sleep(1 * time.Second)
// if inputpin.Read() == rpio.High {
// fmt.Println("High")
// outputpin.High()
// } else if inputpin.Read() == rpio.Low {
// fmt.Println("Low")
// outputpin.Low()
// }
// }
// }

func main() {
	err := rpio.Open()
	if err != nil {
		fmt.Println("failed to rpio.Open()")
		fmt.Println(err.Error())
	}
	defer rpio.Close()
	inputpin := rpio.Pin(27)
	outputpin := rpio.Pin(17)

	inputpin.Input()
	inputpin.PullUp()
	inputpin.Detect(rpio.FallEdge) // enable falling edge event detection
	outputpin.Output()

	for i := 0; i < 2; {
		if inputpin.EdgeDetected() { // check if event occured
			fmt.Println("Motion detected")
			outputpin.High()
			i++
		} else {
			outputpin.Low()
		}
		time.Sleep(time.Second / 2)
	}
	inputpin.Detect(rpio.NoEdge) // disable edge event detection
}
