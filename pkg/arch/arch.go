package arch

import (
	"os/exec"
)

func UpdateSystem() error {
	cmd := exec.Command("sudo", "pacman", "-Syu", "--noconfirm")
	return cmd.Run()
}

func isArchLinux() bool {
	out, err := exec.Command("bash", "-c", "cat /etc/os-release | grep -oP '(?<=^ID=).+'").Output()
	if err != nil {
		return false
	}
	return string(out) == "arch"
}

func ArchCommands() []string {
	return []string{
		"sudo pacman -Syu --noconfirm",
		"sudo pacman -Syu npm sudo python-pip nodejs-lts-gallium git python neovim --noconfirm",
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

