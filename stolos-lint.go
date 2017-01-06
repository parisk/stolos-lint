package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if _, err := os.Stat("docker-compose.yaml"); os.IsNotExist(err) {
		fmt.Fprintln(os.Stderr, "Could not find docker-compose.yml file in the current directory.")
		os.Exit(1)
	}
	dat, fsErr := ioutil.ReadFile("docker-compose.yaml")
	check(fsErr)
	fmt.Print("👍")
	fmt.Print(string(dat))
}
