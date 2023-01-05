package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	for {
		leSitesDoArquivo()
		exibeIntroducao()
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs....")
		case 3:
			fmt.Println("Fechando Aplicação....")
			os.Exit(0)
		default:
			fmt.Println("comando inválido")
			os.Exit(-1)
		}
	}

	/* nome, _ := exibeNomeIdade()
	fmt.Println(nome) */

	/* if comando == 1 {
		fmt.Println("Monitorando.....")
	} else if comando == 2 {
		fmt.Println("Exibindo logs....")
	} else if comando == 0 {
		fmt.Println("Fechando Aplicação....")
	} else {
		fmt.Println("comando inválido")
	} */

}

func exibeNomeIdade() (string, int) {
	nome := "Rafael"
	idade := 21
	return nome, idade
}

func exibeIntroducao() {
	var versao float32 = 1.1
	nome := "Rafael"
	fmt.Println("este software está na versão: ", versao)
	fmt.Println("olá senhor", nome, "Selecione uma opção abaixo")
}

func exibeMenu() {
	fmt.Println("1- iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("3- Sair do Programa")
}

func leComando() int {

	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("o comando escolhido foi o ", comandoLido)
	fmt.Println("o endereço da variavel comando é", &comandoLido)
	fmt.Println("")

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando.....")
	sites := []string{
		"https://www.google.com", "http://www.alura.com.br",
		"http://www.caelum.com.br", "https://random-status-code.herokuapp.com/"}

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println(reflect.TypeOf(site))
			testaSite(site, i)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testaSite(site string, i int) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("o Site;", site, " na posição ", i, " foi carregado com sucesso!")
	} else {
		fmt.Println("o site: ", site, " na posição ", i, " está com problemas... status: ", resp.StatusCode)
	}

}

func leSitesDoArquivo() []string {
	var sites []string

	// arquivo, err := os.Open("sites.txt")
	arquivo, err := ioutil.ReadFile("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	fmt.Println(string(arquivo))

	return sites
}
