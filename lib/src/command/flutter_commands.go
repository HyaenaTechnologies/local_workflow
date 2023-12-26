package command

import (
	"flag"
	"os"
	"strings"
)

// Flutter Commands and Flags
var (
	flutterCommand      = flag.NewFlagSet("flutter", flag.ExitOnError)
	flutterCommandError = flutterCommand.Parse(os.Args[2:])
	flutterBuildFlag    = flutterCommand.String(
		"build",
		"/bin/flutter_application",
		"Build Package: --build Path: <String>")
	flutterDocFlag = flutterCommand.Bool(
		"doc",
		false,
		"Document Package: --doc <Bool>")
	flutterIntegrationFlag = flutterCommand.Bool(
		"integration",
		false,
		"Continuous Integration: --integration <Bool>")
	flutterTestFlag = flutterCommand.String(
		"test",
		".",
		"Run Test: --test Path: <String>")
)

// Flutter Command Strings
var (
	flutterAPKBuild    = "flutter build apk lib/main.dart"
	flutterIntegration = strings.Join(flutterIntegrateCommands, " && ")
	flutterLinuxBuild  = "flutter build linux lib/main.dart"
	flutterRunTest     = "flutter test"
	flutterWebBuid     = "flutter build web lib/main.dart"
)

// Flutter Integration Commands
var flutterIntegrateCommands = []string{
	"flutter analyze lib",
	"flutter doctor",
	"dart fix lib --dry-run",
	"dart fix lib --apply",
	"dart format lib",
	"dart info",
	"flutter pub deps",
	"flutter pub upgrade"}
