package main

import (
	"fmt"
	"os/exec"
    "github.com/AkariOficial/nvim/arch"
    "github.com/AkariOficial/nvim/ubuntu"
)

func isArchLinux() bool {
	out, err := exec.Command("bash", "-c", "cat /etc/os-release | grep -oP '(?<=^ID=).+'").Output()
	if err != nil {
		return false
	}
	return string(out) == "archarm"
}

func isUbuntu() bool {
	out, err := exec.Command("bash", "-c", "cat /etc/os-release | grep -oP '(?<=^ID=).+'").Output()
	if err != nil {
		return false
	}
	return string(out) == "ubuntu"
}

func main() {
	if isArchLinux() {
		cmds := arch.Commands()

		for _, cmd := range cmds {
			err := exec.Command("bash", "-c", cmd).Run()
			if err != nil {
				fmt.Println("Erro ao executar o comando:", cmd)
				return
			}
		}
	} else if isUbuntu() {
		cmds := ubuntu.Commands()

		for _, cmd := range cmds {
			err := exec.Command("bash", "-c", cmd).Run()
			if err != nil {
				fmt.Println("Erro ao executar o comando:", cmd)
				return
			}
		}
	} else {
		fmt.Println("Sistema operacional não suportado.")
		return
	}

	fmt.Println("Configuração concluída com sucesso.")
}

