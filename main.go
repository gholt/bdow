package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	ps "github.com/mitchellh/go-ps"
)

func main() {
	processes, err := ps.Processes()
	if err != nil {
		panic(err)
	}
	pid := 0
	for _, process := range processes {
		if process.Executable() == "BlackDesert64.exe" {
			pid = process.Pid()
		}
	}
	if pid == 0 {
		panic("PID for BlackDesert64.exe not found")
	}
	for {
		netstat := exec.Command("cmd", "/C netstat -ano")
		output, err := netstat.Output()
		if err != nil {
			panic(err)
		}
		lineEndSearch := " " + strconv.Itoa(pid)
		connected := false
		for _, line := range strings.Split(string(output), "\r\n") {
			if strings.HasSuffix(line, lineEndSearch) {
				connected = true
				break
			}
		}
		if !connected {
			process, err := os.FindProcess(pid)
			if err != nil {
				panic(err)
			}
			err = process.Kill()
			if err != nil {
				panic(err)
			}
			fmt.Println("Killed", time.Now())
			return
		}
		time.Sleep(time.Second * 300)
	}
}
