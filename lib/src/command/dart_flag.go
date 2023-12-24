package command

import (
	"flag"
	"fmt"
	"local_workflow/lib/src/integration"
	"log"
	"os"
	"os/exec"
)

// Dart Command Strings
var (
	dartCompileBuild = "dart compile exe lib/main.dart --output"
	dartRunTest      = "dart test"
)

// Dart Commands and Flags
var (
	dartCommand      = flag.NewFlagSet("dart", flag.ExitOnError)
	dartCommandError = dartCommand.Parse(os.Args[2:])
	dartCompileFlag  = dartCommand.String(
		"compile",
		"/bin/dart_application",
		"Compile Package: --compile Path: <String>")
	dartDocFlag = dartCommand.Bool(
		"doc",
		false,
		"Document Package: --doc <Bool>")
	dartIntegrationFlag = dartCommand.Bool(
		"integration",
		false,
		"Continuous Integration: --integration <Bool>")
	dartTestFlag = dartCommand.String(
		"test",
		"",
		"Run Test: --test Path: <String>")
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

// Dart Command Flag Execution
func DartFlagExecute() {
	if *dartCompileFlag != "" {
		DartCompile()
	}
	if *dartIntegrationFlag == true {
		integration.DartAnalyze()
		integration.DartFixDryRun()
		integration.DartFixApply()
		integration.DartFormat()
		integration.DartInfo()
		integration.DartPubDeps()
		integration.DartPubUpgrade()
	}
	if *dartDocFlag == true {
		integration.DartDoc()
	}
	if *dartTestFlag != "" {
		DartTest()
	}
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
