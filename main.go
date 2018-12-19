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
	fmt.Println(time.Now())
	fmt.Println(time.Now(), "\n   ", bdow())
}

func bdow() error {
	processes, err := ps.Processes()
	if err != nil {
		return fmt.Errorf("error obtaining process list %v", err)
	}
	pid := 0
	for _, process := range processes {
		if process.Executable() == "BlackDesert64.exe" {
			pid = process.Pid()
		}
	}
	if pid == 0 {
		return fmt.Errorf("PID for BlackDesert64.exe not found")
	}
	for {
		netstat := exec.Command("cmd", "/C netstat -ano")
		output, err := netstat.Output()
		if err != nil {
			return fmt.Errorf("error running netstat %v", err)
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
				return fmt.Errorf("error finding process for pid %d %v", pid, err)
			}
			err = process.Kill()
			if err != nil {
				return fmt.Errorf("error killing process %v", err)
			}
			return fmt.Errorf("killed")
		}
		time.Sleep(time.Second * 300)
	}
	return nil
}
