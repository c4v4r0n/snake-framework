package clint

import (
	"fmt"

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
	fmt.Println("\n[*]Digite help para obter ajuda[*]")
}

//Prompt ...
func Prompt() {
	var cmd string

	fmt.Printf("\n[+]cmd>> ")
	fmt.Scanf("%s", &cmd)

	switch cmd {
	case "help":
		fmt.Println("\nLista de Comandos:")
		fmt.Println("\nBuster : Fuzzer de url multithread")
		fmt.Println("CVE-2018-15877 : Execução de comandos remoto")

	case "Buster":
		modules.Buster()

	case "CVE-2018-15877":
		modules.CVE201815877()
	}
}
