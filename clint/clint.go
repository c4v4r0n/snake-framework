package clint

import (
	"fmt"
	"os"

	"../modules"
)

// StartInterface ...
func StartInterface() {
	fmt.Println(" ____  _   _    _    _  _______ ")
	fmt.Println("/ ___|| \\ | |  / \\  | |/ / ____|")
	fmt.Println("\\___ \\|  \\| | / _ \\ | ' /|  _|  ")
	fmt.Println(" ___) | |\\  |/ ___ \\| . \\| |___ ")
	fmt.Println("|____/|_| \\_/_/   \\_\\_|\\_\\_____|")
	fmt.Println("                                ")
	fmt.Println(" _____ ____      _    __  __ _______        _____  ____  _  __")
	fmt.Println("|  ___|  _ \\    / \\  |  \\/  | ____\\ \\      / / _ \\|  _ \\| |/ / ")
	fmt.Println("| |_  | |_) |  / _ \\ | |\\/| |  _|  \\ \\ /\\ / / | | | |_) | ' /  ")
	fmt.Println("|  _| |  _ <  / ___ \\| |  | | |___  \\ V  V /| |_| |  _ <| . \\  ")
	fmt.Println("|_|   |_| \\_\\/_/   \\_\\_|  |_|_____|  \\_/\\_/  \\___/|_| \\_\\_|\\_\\ ")
	fmt.Println("\n[*]Digite Help para obter ajuda[*]")
}

//Prompt ...
func Prompt() {
	var cmd string

	fmt.Printf("\n[+]cmd>> ")
	fmt.Scanf("%s", &cmd)

	switch cmd {
	case "Help":
		fmt.Println("\nLista de Comandos:")
		fmt.Println("	Exit		: Encerra a execução")
		fmt.Println("	Buster		: Fuzzer de url multithread")
		fmt.Println("	XSSKeylog 	: Keylogger para XSS")
		fmt.Println("\nLista de modulos de exploits")
		fmt.Println("	CVE-2018-15877	: Execução de comandos remoto")

	case "Buster":
		modules.Buster()
	case "XSSKeylog":
		var port string
		fmt.Printf("Porta: ")
		fmt.Scanf("%s", &port)
		modules.XSSServer(port)

	case "CVE-2018-15877":
		fmt.Printf("Referencias: https://www.exploit-db.com/exploits/45274\n\n")
		modules.CVE201815877()

	case "Exit":
		os.Exit(0)
	}
}
