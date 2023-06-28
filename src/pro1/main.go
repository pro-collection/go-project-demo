package main

import (
	"flag"
	"log"
)

var name string

func main() {
	flag.Parse()

	goCmd := flag.NewFlagSet("go", flag.ExitOnError)

	goCmd.StringVar(&name, "name", "golang", "help - golang")

	phpCmd := flag.NewFlagSet("php", flag.ExitOnError)

	phpCmd.StringVar(&name, "n", "php", "help - php")

	args := flag.Args()

	if len(args) <= 0 {
		return
	}

	switch args[0] {
	case "go":
		{
			_ = goCmd.Parse(args[1:])
			break
		}
	case "php":
		{
			_ = phpCmd.Parse(args[1:])
			break
		}
	default:
		name = ""
		break
	}

	log.Printf("name: %s", name)
}
