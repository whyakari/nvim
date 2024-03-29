package main

import (
	"fmt"
	"os/exec"
	"github.com/whyakari/nvim/pkg/arch"
    "github.com/whyakari/nvim/pkg/termux"
)

func main() {

	if arch.IsArchLinux() {
		cmds := arch.ArchCommands()

		for _, cmd := range cmds {
			err := exec.Command("bash", "-c", cmd).Run()
			if err != nil {
				fmt.Println("Error when executing the command:", cmd)
				return
			}
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
