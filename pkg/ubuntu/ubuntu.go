package ubuntu

import (
	"fmt"
	"os/exec"
	"strings"
)

func UbuntuCommands() []string {
	return []string{
		"apt update",
		"apt install -y curl npm python3-pip nodejs git python3-neovim",
		"pip install --upgrade pip --break-system-packages",
		"pip install wheel --break-system-packages",
		"pip install pynvim neovim --break-system-packages",
		"mkdir -p $HOME/.config/nvim",
		"curl -fLo \"$HOME/.local/share/nvim/site/autoload/plug.vim\" --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim",
		"rm -rf $HOME/nvim",
		"git clone -v https://github.com/whyakari/nvim $HOME/nvim",
		"cp $HOME/nvim/init.vim $HOME/.config/nvim/",
		"rm -rf $HOME/nvim",
	}
}

func IsUbuntu() bool {

    fmt.Println("Installing setup of the neovim, wait...")

	out, err := exec.Command("bash", "-c", "cat /etc/os-release | grep -oP '(?<=^ID=).+'").Output()
	if err != nil {
		return false
	}
    return strings.TrimSpace(string(out)) == "ubuntu"
}

