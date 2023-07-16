package termux

import (
	"fmt"
	"os/exec"
	"strings"
)

func IsTermux() bool {

    fmt.Println("Installing setup of the neovim, wait...")

	out, err := exec.Command("bash", "-c", "uname -o").Output()
	if err != nil {
		return false
	}
    return strings.TrimSpace(string(out)) == "Android"
}

func TermuxCommands() []string {
	return []string{
		"apt update",
		"pkg install nodejs-lts git python neovim -y",
		"pip install --upgrade python-pip --break-system-packages",
		"pip install wheel pynvim neovim --break-system-packages",
		"mkdir -p $HOME/.config/nvim",
		"curl -fLo \"$HOME/.local/share/nvim/site/autoload/plug.vim\" --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim",
        "rm -rf $HOME/nvim",
		"git clone -v https://github.com/whyakari/nvim $HOME/nvim",
		"cp $HOME/nvim/init.vim $HOME/.config/nvim/",
        "cd $HOME",
		"rm -rf $HOME/nvim",
		"rm $HOME/README.md",
        "rm $HOME/*.tar.gz",
	}
}

