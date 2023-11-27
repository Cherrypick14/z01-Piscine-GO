package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(str string) {
	for _, r := range str {
		z01.PrintRune(r)
	}
}

func ReadFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "error"
	}
	return string(data)
}

func main() {
	args := os.Args[1:]
	done := false
	for _, filename := range args {
		if _, err := os.Stat(filename); err != nil {
			PrintStr("ERROR: open " + filename + ": no such file or directory\n")
			os.Exit(1)
		}
		PrintStr(ReadFile(filename))
		done = true
	}
	if !done {
		reader := io.TeeReader(os.Stdin, os.Stdout)
		ioutil.ReadAll(reader)
		os.Stdin.Close()
		os.Stdout.Close()
	}
}
