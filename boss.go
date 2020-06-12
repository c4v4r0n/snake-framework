package main

import (
	"os"

	"./clint"
)

func main() {
	os.Setenv("HTTP_PROXY", "http://localhost:8080")
	clint.StartInterface()
	for true {
		clint.Prompt()
	}
}
