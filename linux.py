from os import system

def baseds(**args):
    for _ in args:
          pass
    return args


data = ["Ubuntu", "Debian", "Xubuntu", "Kubuntu", "Lubuntu"]
data2 = ["Arch", "Manjaro"]


baseUbuntu = baseds(derived=data)
baseArch = baseds(derived=data2)


for distrosUbuntu in baseUbuntu["derived"]:
    try:

        name = open('/etc/issue').readline()

        if distrosUbuntu in name:
            print(name + "System based on Ubuntu Detected.")
            system(r'''apt update; apt upgrade -y; apt curl install git python3 neovim -y; curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.38.0/install.sh | bash; nvm install v16.2.0; export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"; pip install --upgrade pip; pip install wheel; pip install pynvim neovim; mkdir .config; cd .config; mkdir nvim; cd; sh -c 'curl -fLo "${XDG_DATA_HOME:-$HOME/.local/share}"/nvim/site/autoload/plug.vim --create-dirs \
   https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim'; git clone https://github.com/AkariOficial/nvim/; cd nvim; cp init.vim /root/.config/nvim/; cd; rm -rf nvim''')

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
            system('''pacman -Syu; pacman -Syu nodejs-lts-gallium git python neovim && pip install --upgrade pip; pip install wheel; pip install pynvim neovim && mkdir .config; cd .config; mkdir nvim; cd && sh -c 'curl -fLo "${XDG_DATA_HOME:-$HOME/.local/share}"/nvim/site/autoload/plug.vim --create-dirs \
   https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim' && git clone https://github.com/AkariOficial/nvim/; cd nvim; cp init.vim /root/.config/nvim/; cd; rm -rf nvim''')
        elif distrosArch is None:
            pass

    except Exception as e:
        quit()


#print(name + """Distro not supported. If you believe what your is distro Supported, then contact-me in Telegram @AkariOficial.""")
