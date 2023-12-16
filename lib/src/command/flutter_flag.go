package command

import (
	"flag"
	"fmt"
	"local_workflow/lib/src/integration"
	"log"
	"os"
	"os/exec"
)

// Flutter Command Strings
var (
	flutterAPKBuild   = "flutter build apk lib/main.dart"
	flutterLinuxBuild = "flutter build linux lib/main.dart"
	flutterRunTest    = "flutter test"
	flutterWebBuid    = "flutter build web lib/main.dart"
)

// Flutter Commands and Flags
var (
	flutterCommand      = flag.NewFlagSet("flutter", flag.ExitOnError)
	flutterCommandError = flutterCommand.Parse(os.Args[2:])
	flutterBuildFlag    = flutterCommand.String(
		"build",
		"",
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
		"",
		"Run Test: --test Path: <String>")
)

// Build Package APK
func FlutterBuildAPK() {
	var flutterAPK, flutterAPKError = exec.Command(
		flutterAPKBuild,
		*flutterBuildFlag).Output()
	if flutterAPKError != nil {
		log.Fatal(flutterAPKError)
	}
	fmt.Print(flutterAPK)
}

// Build Package Linux
func FlutterBuildLinux() {
	var flutterLinux, flutterLinuxError = exec.Command(
		flutterLinuxBuild,
		*flutterBuildFlag).Output()
	if flutterLinuxError != nil {
		log.Fatal(flutterLinuxError)
	}
	fmt.Print(flutterLinux)
}

// Build Package Web
func FlutterBuildWeb() {
	var flutterWeb, flutterWebError = exec.Command(
		flutterWebBuid,
		*flutterBuildFlag).Output()
	if flutterWebError != nil {
		log.Fatal(flutterWebError)
	}
	fmt.Print(flutterWeb)
}

// Flutter Command Flag Execution
func FlutterFlagExecute() {
	if *flutterBuildFlag != "" {
		FlutterBuildAPK()
		FlutterBuildLinux()
		FlutterBuildWeb()
	}
	if *flutterIntegrationFlag == true {
		integration.FlutterAnalyze()
		integration.FlutterDoctor()
		integration.DartFixDryRun()
		integration.DartFixApply()
		integration.DartFormat()
		integration.DartInfo()
		integration.FlutterPubDeps()
		integration.FlutterPubUpgrade()
	}
	if *flutterDocFlag == true {
		integration.DartDoc()
	}
	if *flutterTestFlag != "" {
		FlutterTest()
	}
}

// Run Test
func FlutterTest() {
	var flutterTest, flutterTestError = exec.Command(
		flutterRunTest,
		*flutterTestFlag).Output()
	if flutterTestError != nil {
		log.Fatal(flutterTestError)
	}
	fmt.Print(flutterTest)
}
