package arch

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"golang.org/x/crypto/ssh/terminal"
)

func isRoot() bool {
	return os.Geteuid() == 0
}

func addSudoToCommands(commands []string, password string) []string {
	sudoCommands := make([]string, len(commands))
	for i, command := range commands {
		sudoCommands[i] = "echo " + password + " | sudo -S " + command
	}
	return sudoCommands
}

func removeRootCommands(commands []string) []string {
	rootCommands := []string{
		"pacman -Syu --noconfirm",
		"pacman -Syu npm sudo python-pip nodejs-lts-gallium git python neovim --noconfirm",
	}

	filteredCommands := make([]string, 0)
	for _, command := range commands {
		if !contains(rootCommands, command) {
			filteredCommands = append(filteredCommands, command)
		}
	}

	return filteredCommands
}

func contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func ArchCommands() []string {
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

	if !isRoot() {
		fmt.Print("Senha do sudo: ")
		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Println("\nErro ao ler a senha do sudo:", err)
			os.Exit(1)
		}
		password := string(bytePassword)

		commands = addSudoToCommands(commands, password)
	} else {
		commands = removeRootCommands(commands)
	}

	return commands
}

func IsArchLinux() bool {
	out, err := exec.Command("bash", "-c", "cat /etc/os-release | grep -oP '(?<=^ID=).+'").Output()
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(out)) == "archarm"
}

