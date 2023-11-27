package main

import (
	"fmt"
	"os"
)

func main() {
	myargs := os.Args[1:]
	checkTrue := false
	for i := 0; i < len(myargs); i++ {
		if myargs[i] == "01" || myargs[i] == "galaxy" || myargs[i] == "galaxy 01" {
			checkTrue = true
		}
	}
	if checkTrue {
		fmt.Println("Alert!!!")
	}
}
