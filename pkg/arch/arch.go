package arch

import (
	"os"
	"os/exec"
	"strings"
)

func IsRootUser() bool {
	return os.Geteuid() == 0
}

func IsArchLinux() bool {
	out, err := exec.Command("bash", "-c", "cat /etc/os-release | grep -oP '(?<=^ID=).+'").Output()
	if err != nil {
		return false
	}
    return strings.TrimSpace(string(out)) == "archarm"
}

func ArchCommands(useSudo bool) []string {
	commands := []string{
		"pacman -Syu --noconfirm",
		"pacman -Syu npm sudo python-pip nodejs-lts-gallium git python neovim --noconfirm",
		"pip install --upgrade pip --break-system-packages",
		"pip install wheel --break-system-packages",
		"pip install pynvim neovim --break-system-packages",
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

    if useSudo {
		// Adicione "sudo" antes de cada comando
		for i, cmd := range commands {
			commands[i] = "sudo " + cmd
		}
	}
	return commands
}
