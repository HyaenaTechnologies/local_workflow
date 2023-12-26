package command

import (
	"flag"
	"os"
	"strings"
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
		".",
		"Run Test: --test Path: <String>")
)

// Dart Command Strings
var (
	dartCompileBuild = "dart compile exe lib/main.dart --output"
	dartDocument     = "dart doc ."
	dartRunTest      = "dart test"
	dartIntegration  = strings.Join(dartIntegrateCommands, " && ")
)

// Dart Integration Commands
var dartIntegrateCommands = []string{
	"dart analyze lib",
	"dart fix lib --dry-run",
	"dart fix lib --apply",
	"dart format lib",
	"dart info",
	"dart pub deps",
	"dart pub upgrade"}
