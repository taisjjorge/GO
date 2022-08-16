//hello.go
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"

	"bufio"
	"net/http"
	"os"
	"reflect"
	"time"
)

func main() {

	showIntrodution()
	showNames()
	readingSites()

	for {
		showMenu()
		comand := commandRead()

		switch comand {
		case 1:
			initMonitoring()
		case 2:
			fmt.Println("Exibindo logs...\n")
			printLog()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Comando escolhido não é uma opção válida")
			os.Exit(-1)
		}
	}
}

func showIntrodution() {
	name := "Tais"
	age := 28

	fmt.Println("O tipo da variável name é", reflect.TypeOf(name))
	fmt.Println("O tipo da variável age é", reflect.TypeOf(age))
	fmt.Println("")
	fmt.Println("Bem vinds,", name)
	fmt.Println("")
}

func showMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func commandRead() int {
	var comand int
	fmt.Scanf("%d", &comand)
	fmt.Println("O comando escolhido foi:", comand)

	return comand
}

func initMonitoring() {
	fmt.Println("Monitorando...")
	sites := readingSites()

	fmt.Println(sites)
	fmt.Println("")

	//for i := 0; i < 5; i++ {
	for i, site := range sites {
		fmt.Println("Testando site", i, ":", site)
		testSites(site)
		println("")
		time.Sleep(5 * time.Second)
	}
	//}
}

func testSites(site string) {
	resp, err := http.Get(site)
	//fmt.Println(resp)
	if err != nil {
		fmt.Println("Vixe! Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site", site, "foi carregado com sucesso")
		registerLogs(site, true)
	} else {
		fmt.Println("Site", site, "indisponível. Status Code:", resp.StatusCode)
		registerLogs(site, false)
	}
}

func showNames() {
	nomes := []string{"tanana", "tananina", "tanananina"}
	fmt.Println("O slice tem", len(nomes))
	nomes = append(nomes, "Maria Eduarda")
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("Agora o slice tem", len(nomes))
	fmt.Println("Mas abriu espaço de", cap(nomes), "porque appendar dobra tamanho")
	fmt.Println("")
}

func readingSites() []string {
	var sites []string
	fileTxt, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Vixe! Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(fileTxt)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}
	fileTxt.Close()
	return sites
}

func registerLogs(site string, status bool) {

	fileLog, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}

	fileLog.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online:" + strconv.FormatBool(status) + "\n")
	fileLog.Close()

}

func printLog() {
	file, err :=  ioutil.ReadFile("logs.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}
