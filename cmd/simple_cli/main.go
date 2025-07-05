package main

import (
	"cli-tmpl/internal/app/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("execute failed err: %s", err.Error())
	}
}
