package main

import (
	"local_workflow/lib/src/command"
	"log"
	"os"
)

func main() {
	go command.DartFlagValidation()
	go command.FlutterFlagValidation()
	go command.GoFlagValidation()
	if len(os.Args) < 2 {
		log.Fatal("Command Required, Help: --help")
	}
	switch os.Args[1] {
	case "dart":
		go command.DartFlagExecute()
	case "flutter":
		go command.FlutterFlagExecute()
	case "go":
		go command.GoFlagExecute()
	default:
		log.Fatal("Unknown Command, Help: --help")
	}
}
