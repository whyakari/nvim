package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"github.com/whyakari/nvim/pkg/arch"
	"github.com/whyakari/nvim/pkg/termux"
	"github.com/whyakari/nvim/pkg/ubuntu"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	if arch.IsArchLinux() {
		cmds := arch.ArchCommands()

		for _, cmd := range cmds {
			err := runCommandWithSudo(cmd)
			if err != nil {
				fmt.Println("Error when executing the command:", cmd)
				return
			}
		}
	} else if ubuntu.IsUbuntu() {
		cmds := ubuntu.UbuntuCommands()

		for _, cmd := range cmds {
			err := runCommandWithSudo(cmd)
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

func runCommandWithSudo(cmd string) error {
	fmt.Print("Password of the sudo: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("\nErro ao ler a senha do sudo:", err)
		return err
	}
	password := string(bytePassword)

	sudoCmd := exec.Command("sudo", "-S", "bash", "-c", cmd)
	sudoCmd.Stdin = strings.NewReader(password + "\n")
	sudoCmd.Stdout = os.Stdout
	sudoCmd.Stderr = os.Stderr

	err = sudoCmd.Run()
	if err != nil {
		return err
	}

	return nil
}

