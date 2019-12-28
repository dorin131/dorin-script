package main

import "os/user"

import "fmt"

import "github.com/dorin131/dorin-script/repl"

import "os"

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hi %s! Have a go at DorinScript!\n", user.Username)
	fmt.Printf("Type a command\n")
	repl.Start(os.Stdin, os.Stdout)
}
