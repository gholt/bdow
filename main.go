package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	ps "github.com/mitchellh/go-ps"
)

func main() {
	log.Print("starting")
	processes, err := ps.Processes()
	if err != nil {
		log.Fatalf("error obtaining process list %v", err)
	}
	pid := 0
	for _, process := range processes {
		if process.Executable() == "BlackDesert64.exe" {
			pid = process.Pid()
		}
	}
	if pid == 0 {
		log.Fatal("pid for BlackDesert64.exe not found")
	}
	log.Printf("pid %d", pid)
	for {
		netstat := exec.Command("cmd", "/C netstat -ano")
		output, err := netstat.Output()
		if err != nil {
			log.Fatalf("error running netstat %v", err)
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
				log.Fatalf("error finding process for pid %d %v", pid, err)
			}
			err = process.Kill()
			if err != nil {
				log.Fatalf("error killing process %v", err)
			}
			log.Fatal("killed")
		}
		time.Sleep(time.Second * 300)
	}
}
