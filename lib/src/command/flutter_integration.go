package command

import (
	"fmt"
	"log"
	"os/exec"
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
	if *flutterBuildFlag != "/bin/flutter_application" {
		FlutterBuildAPK()
		FlutterBuildLinux()
		FlutterBuildWeb()
	}
	if *flutterIntegrationFlag == true {
		FlutterIntegration()
	}
	if *flutterDocFlag == true {
		DartDoc()
	}
	if *flutterTestFlag != "." {
		FlutterTest()
	}
}

func FlutterIntegration() {
	var flutterIntegrate, flutterIntegrateError = exec.Command(
		flutterIntegration).Output()
	if flutterIntegrateError != nil {
		log.Fatal(flutterIntegrateError)
	}
	fmt.Print(flutterIntegrate)
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
