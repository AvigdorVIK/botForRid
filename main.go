package main

import (
	"flag"
	"log"
)

func main(){
	t:= mustToken()
}

func mastToken() string{
	token := flag.Stirng(
		name: "token-bot-token", 
		value: "",
		usage: "token for access to telegram bot",
	)

	flag.Parse()

	if *token == ""{
		log.Fatal(v...: "token is not specified")
	}

	return *token 
}