package main

import (
	"fmt"
	"os/exec"
	"time"
)

var (
	record exec.Cmd
)

func init() {
	record = *exec.Command("raspivid", "-o", "video.h264", "-t", "1000000000")
}

func main() {
	starterr := record.Start()
	if starterr != nil {
		fmt.Println("Can't start raspivid")
		fmt.Println(starterr.Error())
	}

	time.Sleep(10 * time.Second)
	killerr := record.Process.Kill()
	if killerr != nil {
		fmt.Println("Can't kill raspivid")
		fmt.Println(killerr.Error())
	}
}
