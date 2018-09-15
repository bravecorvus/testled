package main

import (
	"fmt"
	"os/exec"
	"time"
)

func raspivid(output string, endsignal chan int) {
	command := exec.Command("raspivid", "-o", "video.h264", "-t", "1000000000")
	starterr := command.Start()
	if starterr != nil {
		fmt.Println("Can't start raspivid")
		fmt.Println(starterr.Error())
	}
	_ = endsignal
	killerr := command.Process.Kill()
	if killerr != nil {
		fmt.Println("Can't kill raspivid")
		fmt.Println(killerr.Error())
	}
}

func main() {
	x := make(chan int)
	raspivid("video2.h264", x)

	time.Sleep(10)
	x <- 1

}
