package main

import (
	"github.com/nickthedev/cutedoc"
	"log"
	"os"
)

// main loads the documentation config in either the current directory or a specified one and generates documentation using it.
func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	if doc, err := cutedoc.New(); err == nil {
		if err = cutedoc.Run(doc); err != nil {
			log.Fatalf("An error occured generating documentation: \n\t%v.", err)
		}
	} else {
		log.Fatalf("An error occurred loading the documentation config: \n\t%v.", err)
	}
}
