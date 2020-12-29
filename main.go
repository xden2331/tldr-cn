package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	jsonpath = "./test/testJSONFile.json"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] <command>\nOptions:", os.Args[0])
		flag.PrintDefaults()
	}

	commands, err := readJSONData(jsonpath)
	if err != nil {
		exitGracefully(err)
	}
	args, err := getArgs()
	if err != nil {
		exitGracefully(err)
	}
	info, err := fetchInformation(commands, args)
	if err != nil {
		exitGracefully(err)
	}
	fmt.Println(info)
}
