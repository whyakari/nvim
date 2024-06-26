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
		"cd $HOME",
        "pacman -Syuu sudo --noconfirm",
		"sudo pacman -Syu --noconfirm",
		"pacman -Syu lua npm python-pip nodejs git python neovim --noconfirm",
		"pacman -S python-wheel --noconfirm",
		"pacman -S python-pynvim python-neovim --noconfirm",
		"mkdir -p $HOME/.config/nvim",
		"curl -fLo \"$HOME/.local/share/nvim/site/autoload/plug.vim\" --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim",
		"git clone -v https://github.com/whyakari/nvim $HOME/nvim_tmp",
		"cp $HOME/nvim_tmp/init.lua $HOME/.config/nvim/",
        "cp -r $HOME/nvim_tmp/lua $HOME/.config/nvim/",
        "cp $HOME/nvim_tmp/lazy-lock.json $HOME/.config/nvim/",
        "cd $HOME",
		"rm -rf nvim",
		"rm -rf $HOME/nvim_tmp",
		"rm $HOME/README.md",
        "rm $HOME/*.tar.gz",
	}
}
