Para instalar o neovim (baseado na minha configuração) siga esses passos

### (Aviso) Lembrando que, para instalar você baixa de acordo com sua GOARCH e GOOS:

#### Support Version
|GOARCH | Architecture | 64bits | 32bits
|-|-|-|-|
Android/Termux | `android` | **arm64** | no
Linux/Termux | `linux` | **arm64** | no
Linux/PC | `linux` | **amd64** | no
--------------------

1. Acesse os Releases: [https://github.com/whyakari/nvim/releases](https://github.com/whyakari/nvim/releases).

2. Pressione o link e copie o URL do zip de acordo com a sua arquitetura. Por exemplo, para a arquitetura **Linux ARM64**, utilize o comando para download:
   ```shell
   wget https://github.com/whyakari/nvim/releases/download/v0.6.1/nvim-v0.6.1-linux-arm64.tar.gz
   ```
   
3. Extraia o arquivo baixado usando o `tar`:
     ```shell
    tar -xznf nvim-v0.6.1-linux-arm64.tar.gz
    ```
> Certifique-se de substituir o nome do arquivo pela correspondente à sua arquitetura.

4. Por fim, execute o comando para iniciar a instalação e configuração automatizada do Neovim automaticamente:
    ```shell
    ./nvim
    ```
> Este comando irá instalar o NVim e realizar a configuração inicial.

### Certifique-se de seguir as instruções acima de acordo com a sua arquitetura para obter a instalação correta do Neovim.
> Lembre-se de substituir o exemplo do URL e do arquivo `.tar`
