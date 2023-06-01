package main

import (
	"fmt"
	"os/exec"
	"github.com/AkariOficial/nvim/pkg/arch"
    "github.com/AkariOficial/nvim/pkg/ubuntu"
    "github.com/AkariOficial/nvim/pkg/termux"
)

func main() {

	if arch.IsArchLinux() {
		cmds := arch.ArchCommands()

		for _, cmd := range cmds {
			err := exec.Command("bash", "-c", cmd).Run()
			if err != nil {
				fmt.Println("Erro ao executar o comando:", cmd)
				return
			}
		}
	} else if ubuntu.IsUbuntu() {
		cmds := ubuntu.UbuntuCommands()

		for _, cmd := range cmds {
			err := exec.Command("bash", "-c", cmd).Run()
			if err != nil {
				fmt.Println("Erro ao executar o comando:", cmd)
				return
			}
        }
    } else if termux.IsTermux() {
        cmds := termux.TermuxCommands()

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
