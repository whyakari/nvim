#!/bin/bash

apt update; apt upgrade -y; pkg in nodejs-lts git python neovim -y && pip install --upgrade pip; pip install wheel pynvim neovim && mkdir .config; cd .config; mkdir nvim; cd && termux-setup-storage && sh -c 'curl -fLo "${XDG_DATA_HOME:-$HOME/.local/share}"/nvim/site/autoload/plug.vim --create-dirs \
   https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim' && git clone https://github.com/AkariOficial/nvim/; cd nvim; cp init.vim /data/data/com.termux/files/home/.config/nvim/; cd; rm -rf nvim \
   bash chmod +x download.sh; rm download.sh
