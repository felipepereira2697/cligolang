package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//Creating a basic cli with Golang
	//To execute: go build .
	//./cligolang -channelName YourStringVariableHere

	//First parameter: variable name, second: the default value, third the channel name
	channel := flag.String("channelName", "Arya", "fp")
	flag.Parse()

	os.Mkdir(fmt.Sprintf("%s", *channel), 077)
	fmt.Println(*channel)

}
