package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan error, 1)
	cmds := []*exec.Cmd{}

	signal.Notify(sigs)

	for _, child := range os.Args[1:] {
		cmdSplit := strings.Fields(child)
		cmd := exec.Command(cmdSplit[0], cmdSplit[1])

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		cmds = append(cmds, cmd)

		go func() {
			done <- cmd.Run()
		}()
	}

	go func() {
		for {
			sig := <-sigs
			for _, c := range cmds {
				c.Process.Signal(sig)
			}
		}
	}()
	err := <-done
	if err != nil {
		fmt.Println("error:", err)
	}
}
