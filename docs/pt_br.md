Para instalar o neovim (baseado na minha configuração) siga esses passos

### (Aviso) Lembrando que, para instalar você baixa de acordo com sua GOARCH e GOOS:

#### Support Version
|GOARCH | Architecture | 64bits | 32bits
|-|-|-|-|
Android/Termux | `android` | **arm64** | no
Linux/Termux | `linux` | **arm64/x64** | no
Linux/PC | `linux` | **amd64/x64** | no
--------------------

1. Acesse os Releases: [https://github.com/whyakari/nvim/releases](https://github.com/whyakari/nvim/releases).

2. Escolha de acordo com sua arquitetura:
   ```shell
   wget example-v0.1.2-arch-arch.tar.gz
   ```
   
3. Extraia o arquivo baixado com `tar`:
     ```shell
    tar -xznf example-v0.1.2-arch-arch.tar.gz
    ```
> Certifique-se de usar o nome correto.

4. Por fim, execute o comando para iniciar a instalação:
    ```shell
    ./nvim
    ```
> Este comando irá instalar o nvim e a minha configuração.

### Certifique-se de seguir as instruções acima de acordo com a sua arquitetura para obter a instalação correta.

## AVISO
> Lembre-se de substituir o exemplo do URL e do arquivo `.tar` quando for extrair.
