package main

import (
	"fmt"
	"os/exec"
)

func main() {
    // commands for executeon terminal.
	cmds := []string{
		"sudo pacman -Syu --noconfirm",
		"sudo pacman -Syu npm sudo python-pip nodejs-lts-gallium git python neovim --noconfirm",
		"pip install --upgrade pip",
		"pip install wheel",
		"pip install pynvim neovim",
		"mkdir -p $HOME/.config/nvim",
		"curl -fLo \"$HOME/.local/share/nvim/site/autoload/plug.vim\" --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim",
		"git clone https://github.com/AkariOficial/nvim $HOME/nvim",
		"cp $HOME/nvim/init.vim $HOME/.config/nvim/",
		"rm -rf $HOME/nvim",
	}

	// run each command in the terminal
	for _, cmd := range cmds {
		err := exec.Command("bash", "-c", cmd).Run()
		if err != nil {
			fmt.Println("error when executing command:", cmd)
			break
		}
	}
}

