from os import system

def baseds(**args):
    for _ in args:
          pass
    return args

data = ["Arch", "Manjaro", "EndeavourOS"]
baseArch = baseds(derived=data)

for distrosArch in baseArch["derived"]:
    try:
        name = open('/etc/issue').readline()
        if distrosArch in name:
            print(name + "System based on Arch Detected.")
            system(r'''sudo pacman -Syu --noconfirm; sudo pacman -Syu npm sudo python-pip nodejs-lts-gallium git python neovim --noconfirm; pip install --upgrade pip; pip install wheel; pip install pynvim neovim; mkdir .config; cd .config; mkdir nvim; cd; sh -c 'curl -fLo "${XDG_DATA_HOME:-$HOME/.local/share}"/nvim/site/autoload/plug.vim --create-dirs \
       https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim'; git clone https://github.com/AkariOficial/nvim/; cd nvim; cp init.vim $HOME/.config/nvim/; cd; rm -rf nvim; rm linux.py;''')
            system("clear")
            system(r"""echo 'Received: ‘Distro not supported’.

If you believe what your is distro Supported, then contact-me in Telegram @AkariOficial or https://t.me/AkariOficial .'""")
        elif distrosArch is None:
            pass

    except Exception as e:
        print("error.")
