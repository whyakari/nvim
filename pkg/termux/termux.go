package termux

import (
	"os/exec"
	"strings"
)

func IsTermux() bool {
	out, err := exec.Command("bash", "-c", "uname -o").Output()
	if err != nil {
		return false
	}
    return strings.TrimSpace(string(out)) == "Android"
}

func TermuxCommands() []string {
	return []string{
		"apt update",
		"apt upgrade -y",
		"pkg install nodejs-lts git python neovim -y",
		"pip install --upgrade python-pip",
		"pip install wheel pynvim neovim",
		"mkdir -p $HOME/.config/nvim",
		"curl -fLo \"$HOME/.local/share/nvim/site/autoload/plug.vim\" --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim",
		"termux-setup-storage",
		"git clone -v https://github.com/AkariOficial/nvim $HOME/nvim",
		"cp $HOME/nvim/init.vim $HOME/.config/nvim/",
        "cd $HOME",
		"rm -rf $HOME/nvim",
		"rm $HOME/README.md",
        "rm $HOME/*.tar.gz",
	}
}

