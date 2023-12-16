package integration

import (
	"fmt"
	"log"
	"os/exec"
)

// Flutter Integration Strings
var (
	flutterAnalysis            = "flutter analyze lib"
	flutterDependencies        = "flutter pub deps"
	flutterPhysician           = "flutter doctor"
	flutterUpgradeDependencies = "flutter pub upgrade"
)

// Analyze Package
func FlutterAnalyze() {
	var flutterAnalyze, flutterAnalyzeError = exec.Command(
		flutterAnalysis).Output()
	if flutterAnalyzeError != nil {
		log.Fatal(flutterAnalyzeError)
	}
	fmt.Print(flutterAnalyze)
}

// Run Doctor
func FlutterDoctor() {
	var flutterDoctor, flutterDoctorError = exec.Command(
		flutterPhysician).Output()
	if flutterDoctorError != nil {
		log.Fatal(flutterDoctorError)
	}
	fmt.Print(flutterDoctor)
}

// Dependency Graph
func FlutterPubDeps() {
	var flutterPubDeps, flutterPubDepsError = exec.Command(
		flutterDependencies).Output()
	if flutterPubDepsError != nil {
		log.Fatal(flutterPubDepsError)
	}
	fmt.Print(flutterPubDeps)
}

// Upgrade Packages
func FlutterPubUpgrade() {
	var flutterPubUpgrade, flutterPubUpgradeError = exec.Command(
		flutterUpgradeDependencies).Output()
	if flutterPubUpgradeError != nil {
		log.Fatal(flutterPubUpgradeError)
	}
	fmt.Print(flutterPubUpgrade)
}
