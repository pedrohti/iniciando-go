package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	exibeIntroducao()

	for {
		exibeMenu()
		comando := lerComando()

		switch comando {
		case 1:
			inciciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLog()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Comando Inválido!")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Seu Nome"
	versao := 1.3

	fmt.Println("Olá, Sr(a).", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")

	return comandoLido
}

func inciciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := leArquivo()

	for i := 0; i < monitoramentos; i++ {
		for indice, item := range sites {
			fmt.Println("Testando site", indice, ":", item)
			testaSite(item)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro ao testar os sites:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site", site, "carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("O site", site, "não está carregando! Erro:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de log:", err)
	}

	arquivo.WriteString("Online: " + strconv.FormatBool(status) + " - " + time.Now().Format("02/01/2006 15:04:05") + " - " + site + "\n")
	arquivo.Close()
}

func imprimeLog() {
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo de log:", err)
	}

	fmt.Println(string(arquivo))
}
