### Minha config do neovim

<p float="left">
    <img src="https://user-images.githubusercontent.com/58480908/191308048-12ca1658-5185-4c47-8298-95accc888706.png" width=130/>
    <img src="https://user-images.githubusercontent.com/58480908/191308172-69a9e8b0-f435-45a9-8383-f266f7812226.png" width=130/>
</p> 

> **Only users of Termux <br> and Proot-Distro!** <br> do not try in Linux(__Ubuntu__, etc..) because process is quite different.


#### Requirement: install package ```wget```.

###### In Termux:
>   pkg in -y wget<br>
###### In Linux (Ubuntu):
>   apt install -y wget 
###### In Linux (**Arch**):
>   pacman -S wget

#### 🥥 Download (for termux):
    wget https://raw.githubusercontent.com/AkariOficial/nvim/main/download.sh && bash ./download.sh
#### 🦠 Download (for linux ```Ubuntu``` and derived's.</br>```Arch``` and derived's):
    wget https://raw.githubusercontent.com/AkariOficial/nvim/main/linux.py; python3 linux.py

##### write ```nvim``` in your terminal, and then ```:PlugInstall``` (I'm trusting you to know how to deal with terminal and neovim itself)
> by and thanks to modifying:<br>[outragedline 🇧🇷](https://github.com/outragedline/neovim-termux)
