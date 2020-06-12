package main

import (
	"./clint"
)

func main() {
	clint.StartInterface()
	for true {
		clint.Prompt()
	}
}
