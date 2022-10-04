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

const monitorings = 3
const monitoriongDelay = 5

func main() {
	showIntroduction()
	
	for {
		showMenu()
		
		command := readCommand()
		
		switch command {
		case 1:
			initMonitoring()
			
		case 2:
			printLogs()

		case 0:
			fmt.Println("Saindo do programa ...")
			os.Exit(0)
			
		default: 
			fmt.Println("Nao conheco este comando")
			os.Exit(-1)
		}
	}
}

func showMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa")
}

func showIntroduction() {
	name := "Gabriel Borel"
	version := 1.1
	
	fmt.Println("Olá, sr. ", name)
	fmt.Println("Este programa está na versão", version)
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	
	fmt.Println("O comando escolhido foi", command)
	fmt.Println("")
	
	return command
}

func testWebsite(website string) {
	response, err := http.Get(website)
	
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Website:", website, "OK")
		registerLogs(website, true)
	} else {
		fmt.Println("Website:", website, "está com problemas", response.StatusCode)
		registerLogs(website, false)
	}
}
	
func readWebsitesFromFile() []string {
	var websites []string

	file, err := os.Open("websites.txt")
	
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		websites = append(websites, line)
		
		if err == io.EOF {
			break
		}
	}

	file.Close()

	return websites
}

func registerLogs(website string, status bool) {
	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	today := time.Now().Format("02/01/2006 15:04:05")
	file.WriteString(today + " - " + website + " - online: " + strconv.FormatBool(status) + "\n")

	fmt.Println(file)

	file.Close()
}

func printLogs() {
	fmt.Println("Exibindo logs ...")

	file, err := ioutil.ReadFile("logs.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(file))
}

func initMonitoring() {
	fmt.Println("Monitorando ...")

	websites := readWebsitesFromFile()

	for i := 0; i < monitorings; i++ {
		for _, website := range websites {
			testWebsite(website)
		}

		time.Sleep(monitoriongDelay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}