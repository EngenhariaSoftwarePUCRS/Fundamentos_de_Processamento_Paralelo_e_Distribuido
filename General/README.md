## Pasta para instalação e testes com a linguagem Go


**Sobre Instalação**

Instale Go, instruções em: [https://golang.org/doc/tutorial/getting-started](https://golang.org/doc/tutorial/getting-started)

Nas versões mais recentes de go, existe a necessidade de criar um módulo.
1. Crie um diretório para seus fontes, por exemplo goprogs.
2. Dentro desta pasta voce ira criar seus programas em Go.
3. Entre no diretório goprogs (cd gopros)
4. Entre: "go mod init goprogs".    Este comando criará um arquivo go.mod em goprogs
5. Edite em todos os arquivos e pastas para que os "import ./...." tornem-se "import goprogs/...."
6. Na pasta onde esta o arquivo nomeArquivo.go que contém main, comande "go run nomeArquivo.go*

### Links Úteis

- [Getting Started](https://go.dev/doc/tutorial/getting-started)
- [Download](https://go.dev/doc/install)
- [Guided Learning](https://go.dev/learn/#guided-learning-journeys)
- [Playground](https://go.dev/play/)
- [ElementosBásicosDeGo - Fernando Dotti](https://www.youtube.com/watch?v=3nod1UIcgWY)
