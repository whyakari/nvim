package ubuntu

import (
	"os/exec"
	"strings"
)

func UbuntuCommands() []string {
	return []string{
		"apt update",
		"apt upgrade -y",
		"apt install -y npm python3-pip nodejs git python3-neovim",
		"pip install --upgrade pip",
		"pip install wheel",
		"pip install pynvim neovim",
		"mkdir -p $HOME/.config/nvim",
		"curl -fLo \"$HOME/.local/share/nvim/site/autoload/plug.vim\" --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim",
		"rm -rf $HOME/nvim",
		"git clone -v https://github.com/AkariOficial/nvim $HOME/nvim",
		"cp $HOME/nvim/init.vim $HOME/.config/nvim/",
		"rm -rf $HOME/nvim",
	}
}

func IsUbuntu() bool {
	out, err := exec.Command("bash", "-c", "cat /etc/os-release | grep -oP '(?<=^ID=).+'").Output()
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(out)) == "ubuntu"
}

