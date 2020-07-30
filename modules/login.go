package modules

import (
	"fmt"
	"net/http"
	"strings"
)

//Login ...
func Login(url string) string {
	var username, password, stringForm string
	stringForm = "&wp-submit=Acessar&redirect_to=" + url
	stringForm += "%2Fwp-admin%2F&testcookie=1"
	fmt.Printf("Usuario: ")
	fmt.Scanf("%s", &username)
	fmt.Printf("Senha: ")
	fmt.Scanf("%s", &password)

	client := new(http.Client)
	formData := strings.NewReader("log=" + username + "&pwd=" + password + stringForm)
	req, _ := http.NewRequest("POST", url+"/wp-login.php", formData)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(req)

	cookie := strings.Join(resp.Header["Set-Cookie"], " ")
	indC := strings.Index(cookie, "path=/wordpress/ ")
	cookie = cookie[indC+17:]

	return cookie
}
