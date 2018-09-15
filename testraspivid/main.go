package main

import (
	"fmt"
	"os/exec"
	"time"
)

func record(output string, endsignal chan int) {
	raspivid := exec.Command("raspivid", "-o", output+".h264", "-t", "1000000000")
	starterr := raspivid.Start()
	if starterr != nil {
		fmt.Println("Can't start raspivid")
		fmt.Println(starterr.Error())
	}
	_ = endsignal
	killerr := raspivid.Process.Kill()
	if killerr != nil {
		fmt.Println("Can't kill raspivid")
		fmt.Println(killerr.Error())
	}

	MP4Box := exec.Command("MP4Box", "-add", output+".h264", output+".mp4")
	_, killerr = MP4Box.Output()
	if killerr != nil {
		fmt.Println("Can't kill MP$Box")
		fmt.Println(killerr.Error())
	}

}

func main() {
	x := make(chan int)
	go record("video2.h264", x)

	time.Sleep(10)
	x <- 1

}
