package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	exibeIntroducao()

	for {
		exibeMenu()
		// _, idade := devolveNome() //com underline ignora um ou mais valores
		// fmt.Println(idade)

		comando := lerComando()

		switch comando {
		case 1:
			inciciarMonitoramento()
		case 2:
			fmt.Println("Logs Exibidos")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Comando Inválido!")
			os.Exit(-1)
		}
	}
}

// func devolveNome() (string, int) { //função de exemplo para exibir o returno multiplo de uma função
// 	nome := "Pedro"
// 	idade := 27
// 	return nome, idade
// }

func exibeIntroducao() {
	nome := "Pedro"
	versao := 1.2

	fmt.Println("Olá, Sr.", nome)
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
	sites := []string{
		"https://www.alura.com.br",
		"https://random-status-code.herokuapp.com/",
		"https://www.covidvisualizer.com/",
	}

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
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("O site", site, "carregado com sucesso!")
	} else {
		fmt.Println("O site", site, "não está carregando! Erro:", resp.StatusCode)
	}
}
