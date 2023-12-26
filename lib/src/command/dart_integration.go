package command

import (
	"fmt"
	"log"
	"os/exec"
)

// Compile Package
func DartCompile() {
	var dartCompile, dartCompileError = exec.Command(
		dartCompileBuild,
		*dartCompileFlag).Output()
	if dartCompileError != nil {
		log.Fatal(dartCompileError)
	}
	fmt.Print(dartCompile)
}

// Document Package
func DartDoc() {
	var dartDoc, dartDocError = exec.Command(
		dartDocument).Output()
	if dartDocError != nil {
		log.Fatal(dartDocError)
	}
	fmt.Print(dartDoc)
}

// Dart Command Flag Execution
func DartFlagExecute() {
	if *dartCompileFlag != "/bin/dart_application" {
		DartCompile()
	}
	if *dartIntegrationFlag == true {
		DartIntegration()
	}
	if *dartDocFlag == true {
		DartDoc()
	}
	if *dartTestFlag != "." {
		DartTest()
	}
}

func DartIntegration() {
	var dartIntegrate, dartIntegrateError = exec.Command(
		dartIntegration).Output()
	if dartIntegrateError != nil {
		log.Fatal(dartIntegrateError)
	}
	fmt.Print(dartIntegrate)
}

// Run Test
func DartTest() {
	var dartTest, dartTestError = exec.Command(
		dartRunTest,
		*dartTestFlag).Output()
	if dartTestError != nil {
		log.Fatal(dartTestError)
	}
	fmt.Print(dartTest)
}
