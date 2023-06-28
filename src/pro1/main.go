package main

import (
	"flag"
	"log"
)


func main() {
	var name string

	flag.StringVar(&name, "name", "", "help")

	flag.Parse()

	log.Printf("Hello %s", name)
}
