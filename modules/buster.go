package modules

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

var ther = make([]int, 0, 100)

//Buster ...
func Buster() {
	var wlPath, host string

	fmt.Printf("Wordlist: ")
	fmt.Scanf("%s", &wlPath)
	fmt.Printf("Url: ")
	fmt.Scanf("%s", &host)

	var file, _ = os.Open(wlPath)
	var scanner = bufio.NewScanner(file)

	dirname := make(chan string)
	quit := make(chan int)
	defer file.Close()
	url := host

	go func() {
		for scanner.Scan() {
			dirname <- scanner.Text()
		}
		defer func() { quit <- 1 }()
	}()

	func() {
		for true {
			if len(ther) < cap(ther) {
				select {
				case v := <-dirname:
					ther = append(ther, 1)
					go worker(url, v)
				case <-quit:
					fmt.Println("\n============\n\nFinish")
					return
				default:
					continue
				}
			} else {
				time.Sleep(1000 * time.Millisecond)
			}
		}
	}()
}

func worker(url, dirName string) {
	resp, err := http.Get(url + "/" + dirName)
	if err == nil {
		if resp.StatusCode != 404 {
			fmt.Println("/"+dirName+"  Status ", resp.StatusCode)
		}
	}
	ther = ther[:len(ther)-1]
}
