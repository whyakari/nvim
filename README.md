# minhas configurações do neovim

1 -
   - apt update; apt upgrade -y
   - pkg in nodejs-lts git python neovim -y

##

2 -
   - pip install --upgrade pip
   - pip install wheel pynvim  neovim

##

3 -
    # criar configuração manualmente(home do termux)
        -  mkdir .config; cd .config; mkdir nvim

    # navegar pra onde esta o init.vim
       - termux-setup-storage  # digite sim
       - cd /sdcard/Download

    # copiar config para:
       - cp init.vim /data/data/com.termux/files/home/.config/nvim/

##

4 - instalar os plugins
  ## digite: `nvim` e depois `:PlugInstall`
 
  ## sair do nvim
    - digite: ctrl x ou :q

  ## instalar extensões
    - digite: `:CocInstall coc-pyright`
