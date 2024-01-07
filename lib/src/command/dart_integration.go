package command

import (
	"fmt"
	"log"
	"os/exec"
)

type DartCommands struct {
	build     string
	document  string
	integrate []string
	test      string
}

// Compile Package
func (compileBuild DartCommands) DartCompile() {
	var dartCompile, dartCompileError = exec.Command(
		compileBuild.build,
		*dartCompileFlag,
	).Output()
	if dartCompileError != nil {
		log.Fatal(dartCompileError)
	}
	fmt.Print(dartCompile)
}

// Document Package
func (documentation DartCommands) DartDoc() {
	var dartDoc, dartDocError = exec.Command(
		documentation.document,
	).Output()
	if dartDocError != nil {
		log.Fatal(dartDocError)
	}
	fmt.Print(dartDoc)
}

// Dart Command Flag Execution
func DartFlagExecute() {
	if *dartCompileFlag != "/bin/dart_application" {
		dartCompileBuild.DartCompile()
	}
	if *dartIntegrationFlag == true {
		DartIntegration()
	}
	if *dartDocFlag == true {
		dartDocument.DartDoc()
	}
	if *dartTestFlag != "." {
		dartRunTest.DartTest()
	}
}

func DartIntegration() {
	var dartIntegrate, dartIntegrateError = exec.Command(
		dartIntegration,
	).Output()
	if dartIntegrateError != nil {
		log.Fatal(dartIntegrateError)
	}
	fmt.Print(dartIntegrate)
}

// Run Test
func (runTest DartCommands) DartTest() {
	var dartTest, dartTestError = exec.Command(
		runTest.test,
		*dartTestFlag,
	).Output()
	if dartTestError != nil {
		log.Fatal(dartTestError)
	}
	fmt.Print(dartTest)
}
