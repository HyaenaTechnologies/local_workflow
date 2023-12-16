package command

import (
	"log"
)

// Dart Command Flag Validation
func DartFlagValidation() {
	if dartCommandError != nil {
		log.Fatal(dartCommandError)
	}
	if *dartCompileFlag == "" &&
		*dartIntegrationFlag == false {
		dartCommand.PrintDefaults()
		log.Fatalf("Help: --help")
	}
}

// Flutter Command Flag Validation
func FlutterFlagValidation() {
	if *flutterBuildFlag == "" &&
		*flutterIntegrationFlag == false {
		flutterCommand.PrintDefaults()
		log.Fatalf("Help: --help")
	}
	if flutterCommandError != nil {
		log.Fatal(flutterCommandError)
	}
}

// Go Command Flag validation
func GoFlagValidation() {
	if *goBuildFlag == "" &&
		*goIntegrationFlag == false {
		goCommand.PrintDefaults()
		log.Fatalf("Help: --help")
	}
	if goCommandError != nil {
		log.Fatal(goCommandError)
	}
}
