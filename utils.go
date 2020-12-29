package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type example struct {
	Description string
	Command     string
}

type command struct {
	ID          string
	Keywords    []string
	Description string
	Examples    []example
}

func exitGracefully(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

type arguments struct {
	listAll bool
	query   string
}

func fetchInformation(commands []command, args arguments) (string, error) {
	var sb strings.Builder
	var err error
	if args.listAll {
		sb.WriteString("Supported commands:\n")
		for _, command := range commands {
			id := command.ID
			desc := command.Description
			sb.WriteString("-" + id + "\n")
			sb.WriteString("\t" + desc + "\n")
		}
	} else if idx := sort.Search(len(commands), func(i int) bool { return commands[i].ID == args.query }); idx >= 0 && idx < len(commands) {
		command := commands[idx]
		sb.WriteString(command.ID + "\n")
		sb.WriteString(command.Description + "\n")
		for _, example := range command.Examples {
			sb.WriteString("- " + example.Description + "\n")
			sb.WriteString("\t" + example.Command + "\n")
		}
	} else {
		err = errors.New("Not supported command " + args.query)
	}
	return sb.String(), err
}

func getArgs() (arguments, error) {
	// Define option flags
	listAll := flag.Bool("list", false, "List all supported commands.")

	// Parse all command line args
	flag.Parse()

	// The query command
	query := flag.Arg(0)

	args := arguments{listAll: *listAll, query: query}

	// Check if listAll flag is true, then there should be no query
	if *listAll && query != "" {
		return arguments{}, errors.New("Cannot enter query when --list is set")
	}

	return args, nil
}

func readJSONData(path string) ([]command, error) {
	bs, err := ioutil.ReadFile(path)
	var commands []command

	if err != nil {
		return commands, err
	}
	if !json.Valid(bs) {
		return commands, errors.New("Invalid JSON representation")
	}
	err = json.Unmarshal(bs, &commands)
	if err != nil {
		return commands, err
	}
	return commands, nil
}
