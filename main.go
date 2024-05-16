package main

import (
	"fmt"
	"github.com/whyakari/nvim/pkg/arch"
	"github.com/whyakari/nvim/pkg/termux"
	"os/exec"
)

func main() {

	if arch.IsArchLinux() {
		cmds := arch.ArchCommands()

		for _, cmd := range cmds {
			fmt.Printf("Executing: %s\n", cmd)
			out, err := exec.Command("bash", "-c", cmd).CombinedOutput()
			if err != nil {
				fmt.Printf("Error executing command '%s': %s\nOutput: %s\n", cmd, err, string(out))
				return
			}
			fmt.Printf("Output: %s\n", string(out))
		}

	} else if termux.IsTermux() {
		cmds := termux.TermuxCommands()

		for _, cmd := range cmds {
			err := exec.Command("bash", "-c", cmd).Run()
			if err != nil {
				fmt.Println("Error when executing the command:", cmd)
				return
			}
		}
	} else {
		fmt.Println("Operating system not supported.")
		return
	}
	fmt.Println("Setup completed successfully (:")
}
