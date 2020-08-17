package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

// The process could be executed asynchronously too.
// This is done by calling the start methods of the Cmd structure. In this case, the
// process is executed, but the main gorutine does not wait until it ends.
// The Wait method could be used to wait until the process ends. After the Wait method finishes,
// the resources of the process are released.
func main() {
	prc := exec.Command("ls", "-a")
	out := bytes.NewBuffer([]byte{})
	prc.Stdout = out
	err := prc.Start()
	if err != nil {
		fmt.Println(err)
	}

	prc.Wait()

	if prc.ProcessState.Success() {
		fmt.Println("Process run successfully with output:")
		fmt.Println(out.String())
	}
}
