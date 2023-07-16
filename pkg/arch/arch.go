package arch

import (
	"fmt"
	"os/exec"
	"strings"
)

func IsArchLinux() bool {

    fmt.Println("Installing setup of the neovim, wait...")

	out, err := exec.Command("bash", "-c", "cat /etc/os-release | grep -oP '(?<=^ID=).+'").Output()
	if err != nil {
		return false
	}
    return strings.TrimSpace(string(out)) == "archarm" || strings.TrimSpace(string(out)) == "arch"
}

func ArchCommands() []string {
	return []string{
        "pacman -Syuu sudo --noconfirm",
		"sudo pacman -Syu --noconfirm",
		"sudo pacman -Syu npm python-pip nodejs-lts-gallium git python neovim --noconfirm",
		"pip install --upgrade pip --break-system-packages",
		"pip install wheel --break-system-packages",
		"pip install pynvim neovim --break-system-packages",
		"sudo mkdir -p $HOME/.config/nvim",
		"sudo curl -fLo \"$HOME/.local/share/nvim/site/autoload/plug.vim\" --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim",
		"sudo rm -rf $HOME/nvim",
		"sudo git clone -v https://github.com/whyakari/nvim $HOME/nvim",
		"sudo cp $HOME/nvim/init.vim $HOME/.config/nvim/",
        "cd $HOME",
		"sudo rm -rf $HOME/nvim",
        "sudo rm $HOME/README.md",
        "sudo rm $HOME/*.tar.gz",
	}
}
