package command

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type FlutterCommands struct {
	apk       string
	integrate []string
	linux     string
	test      string
	web       string
}

// Flutter Command Strings
var (
	flutterAPKBuild = FlutterCommands{
		apk: "flutter build apk lib/main.dart",
	}
	flutterIntegration = strings.Join(
		flutterIntegrate.integrate,
		" && ",
	)
	flutterLinuxBuild = FlutterCommands{
		linux: "flutter build linux lib/main.dart",
	}
	flutterRunTest = FlutterCommands{
		test: "flutter test",
	}
	flutterWebBuid = FlutterCommands{
		web: "flutter build web lib/main.dart",
	}
)

// Flutter Integration Commands
var flutterIntegrate = FlutterCommands{
	integrate: []string{
		"flutter analyze lib",
		"flutter doctor",
		"dart fix lib --dry-run",
		"dart fix lib --apply",
		"dart format lib",
		"dart info",
		"flutter pub deps",
		"flutter pub upgrade"},
}

// Build Package APK
func (apkBuild FlutterCommands) FlutterAPK() {
	var flutterAPK, flutterAPKError = exec.Command(
		apkBuild.apk,
		*flutterBuildFlag,
	).Output()
	if flutterAPKError != nil {
		log.Fatal(flutterAPKError)
	}
	fmt.Print(flutterAPK)
}

// Build Package Linux
func (linuxBuild FlutterCommands) FlutterLinux() {
	var flutterLinux, flutterLinuxError = exec.Command(
		linuxBuild.linux,
		*flutterBuildFlag,
	).Output()
	if flutterLinuxError != nil {
		log.Fatal(flutterLinuxError)
	}
	fmt.Print(flutterLinux)
}

// Build Package Web
func (webBuild FlutterCommands) FlutterWeb() {
	var flutterWeb, flutterWebError = exec.Command(
		webBuild.web,
		*flutterBuildFlag,
	).Output()
	if flutterWebError != nil {
		log.Fatal(flutterWebError)
	}
	fmt.Print(flutterWeb)
}

// Flutter Command Flag Execution
func FlutterFlagExecute() {
	if *flutterBuildFlag != "/bin/flutter_application" {
		flutterAPKBuild.FlutterAPK()
		flutterLinuxBuild.FlutterLinux()
		flutterWebBuid.FlutterWeb()
	}
	if *flutterIntegrationFlag == true {
		FlutterIntegration()
	}
	if *flutterDocFlag == true {
		dartDocument.DartDoc()
	}
	if *flutterTestFlag != "." {
		flutterRunTest.FlutterTest()
	}
}

func FlutterIntegration() {
	var flutterIntegrate, flutterIntegrateError = exec.Command(
		flutterIntegration,
	).Output()
	if flutterIntegrateError != nil {
		log.Fatal(flutterIntegrateError)
	}
	fmt.Print(flutterIntegrate)
}

// Run Test
func (runTest FlutterCommands) FlutterTest() {
	var flutterTest, flutterTestError = exec.Command(
		runTest.test,
		*flutterTestFlag,
	).Output()
	if flutterTestError != nil {
		log.Fatal(flutterTestError)
	}
	fmt.Print(flutterTest)
}
