package main

import (
	"flag"
	"fmt"
)

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
}

var token string

func main() {

	if token == "" {
		fmt.Println("No token provided. Please provide the token: -t <bot token>")
		return
	}
	setup(token)
}
