package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	fmt.Println(os.Args[1])
	raspivid := exec.Command("raspivid", "-o", os.Args[1]+".h264", "-t", "1000000000")

	starterr := raspivid.Start()
	if starterr != nil {
		fmt.Println("Can't start raspivid")
		fmt.Println(starterr.Error())
	}

	time.Sleep(10)

	killerr := raspivid.Process.Kill()
	if killerr != nil {
		fmt.Println("Can't kill raspivid")
		fmt.Println(killerr)
	}
}
