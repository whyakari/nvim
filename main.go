package main

import (
	"fmt"
	"os/exec"
	"github.com/whyakari/nvim/pkg/arch"
    "github.com/whyakari/nvim/pkg/ubuntu"
    "github.com/whyakari/nvim/pkg/termux"
)

func main() {
	if arch.IsArchLinux() {
		useSudo := !arch.IsRootUser() // Verifica se o usuário não é root
		cmds := arch.ArchCommands(useSudo)

		for _, cmd := range cmds {
			err := exec.Command("bash", "-c", cmd).Run()
			if err != nil {
				fmt.Println("Error when executing the command:", cmd)
				return
			}
		}
	} else if ubuntu.IsUbuntu() {
		// ...
	} else if termux.IsTermux() {
		// ...
	} else {
		// ...
	}

	fmt.Println("Setup completed successfully (:")
}
