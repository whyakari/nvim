from os import system

def baseds(**args):
    for _ in args:
          pass
    return args


data = ["Ubuntu", "Debian", "Xubuntu", "Kubuntu", "Lubuntu"]
data2 = ["Arch", "Manjaro"]
data3 = ["alpine"]

baseUbuntu = baseds(derived=data)
baseArch = baseds(derived=data2)
baseAlpine = baseds(derived=data3)

for distrosUbuntu in baseUbuntu["derived"]:
    try:

        name = open('/etc/issue').readline()

        if distrosUbuntu in name:
            print(name + "System based on Ubuntu Detected.")
            system(r'''sudo apt update; sudo apt upgrade -y && sudo apt install curl git python3 -y; curl -LO https://github.com/neovim/neovim/releases/latest/download/nvim.appimage && chmod u+x nvim.appimage && ./nvim.appimage --appimage-extract
./squashfs-root/AppRun --version &&
sudo mv squashfs-root /
sudo ln -s /squashfs-root/AppRun /usr/bin/nvim && curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.38.0/install.sh | bash; export NVM_DIR="$HOME/.nvm"; [ -s "$NVM_DIR/nvm.sh" ]; \. "$NVM_DIR/nvm.sh"; [ -s "$NVM_DIR/bash_completion" ]; \. "$NVM_DIR/bash_completion"; nvm install v16.2.0; pip install --upgrade pip; pip install wheel; pip install pynvim neovim; mkdir .config; cd .config; mkdir nvim; cd; sh -c 'curl -fLo "${XDG_DATA_HOME:-$HOME/.local/share}"/nvim/site/autoload/plug.vim --create-dirs; https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim'; git clone https://github.com/AkariOficial/nvim/; cd nvim; cp init.vim $HOME/.config/nvim/; cd; sudo rm -rf nvim; rm linux.py; clear; rm nvim.appimage && rm -rf squashfs-root;''')
            system("clear")
            system(r"""echo 'Received: ‘Distro not supported’.

If you believe what your is distro Supported, then contact-me in Telegram @AkariOficial or https://t.me/AkariOficial .'""")

        elif distrosUbuntu is None:
            pass

    except Exception as e:
        print("error, problaby is using termux.")
        quit()


for distrosArch in baseArch["derived"]:
    try:

        name = open('/etc/issue').readline()
        if distrosArch in name:
            print(name + "System based on Arch Detected.")
            system(r'''pacman -Syu; pacman -Syu npm; pacman -S python-pip; pacman -Syu nodejs-lts-gallium git python neovim; pip install --upgrade pip; pip install wheel; pip install pynvim neovim; mkdir .config; cd .config; mkdir nvim; cd; sh -c 'curl -fLo "${XDG_DATA_HOME:-$HOME/.local/share}"/nvim/site/autoload/plug.vim --create-dirs \
   https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim'; git clone https://github.com/AkariOficial/nvim/; cd nvim; cp init.vim $HOME/.config/nvim/; cd; rm -rf nvim; rm linux.py;''')
            system("clear")
            system(r"""echo 'Received: ‘Distro not supported’.

If you believe what your is distro Supported, then contact-me in Telegram @AkariOficial or https://t.me/AkariOficial .'""")
        elif distrosArch is None:
            pass

    except Exception as e:
        print("error.")
        quit()


for distroAlpine in baseAlpine["derived"]:
    try:
        name = open('/etc/issue').readline()
        if distroAlpine in name:
            print(name + "System on based Alpine Detected.")

            system(r'''apk update; apk upgrade; apk add nodejs-lts git python3 py3-pip neovim; pip install --upgrade pip; pip install wheel;
pip pynvim neovim; mkdir .config; cd .config; mkdir nvim; cd; sh -
c 'curl -fLo "${XDG_DATA_HOME:-$HOME/.local/share}"/nvim/site/auto
load/plug.vim --create-dirs \
>        https://raw.githubusercontent.com/junegunn/vim-plug/maste
r/plug.vim'; git clone https://github.com/AkariOficial/nvim/; cd n
vim; cp init.vim $HOME/.config/nvim/; cd; rm linux.py;''')

            system("clear")
            system(r"""echo 'Received: ‘Distro not supported'
If you believe what your is distro Supported, then contact-me
in Telegram @AkariOficial or https://t.me/AkariOficial .'""")
        elif distroAlpine is None:
            pass

    except Exception as e:
        print("erorr.")
        quit()

