## minhas configurações do neovim
#####
  ##### atualizar pacotes e instalar dependências necessárias
    apt update; apt upgrade -y; pkg in nodejs-lts git python neovim -y
#####
  ##### atualizar o pip e instalar dependências necessárias
    pip install --upgrade pip; pip install wheel pynvim neovim
#####
  ##### crie a configuração manualmente necessária para funcionar na ```home``` do termux
    mkdir .config; cd .config; mkdir nvim; cd
####
  ##### dê acesso de armazenamento no termux usando o comando:
    termux-setup-storage 
###
 confirme usando ```sim```
####
  #### download do ```plug.vim```
    sh -c 'curl -fLo "${XDG_DATA_HOME:-$HOME/.local/share}"/nvim/site/autoload/plug.vim --create-dirs \
       https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim'
 
##### download do meu ```init.vim```
    git clone https://github.com/Akaritsu/nvim/; cd nvim; cp init.vim /data/data/com.termux/files/home/.config/nvim/; cd; rm -rf nvim
#####
  ##### instalar os plugins digite nvim e depois ```PlugInstall``` e vá dando enter até terminar
  ##### sair do nvim, digite: 
    ctrl x ou :q
  ##### instalar extensões, digite: 
    CocInstall coc-pyright
####
  #### github a qual usei como base [outragedline](https://github.com/outragedline/neovim-termux)
