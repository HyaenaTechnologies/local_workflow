package integration

import (
	"fmt"
	"log"
	"os/exec"
)

// Dart Integration Strings
var (
	dartAnalysis            = "dart analyze lib"
	dartApplyFixes          = "dart fix lib --apply"
	dartDependencies        = "dart pub deps"
	dartDocument            = "dart doc ."
	dartDryRunFixes         = "dart fix lib --dry-run"
	dartEnvironment         = "dart info"
	dartFormatCode          = "dart format lib"
	dartUpgradeDependencies = "dart pub upgrade"
)

// Analyze Package
func DartAnalyze() {
	var dartAnalyze, dartAnalyzeError = exec.Command(
		dartAnalysis).Output()
	if dartAnalyzeError != nil {
		log.Fatal(dartAnalyzeError)
	}
	fmt.Print(dartAnalyze)
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

// Environment Information
func DartInfo() {
	var dartInfo, dartInfoError = exec.Command(
		dartEnvironment).Output()
	if dartInfoError != nil {
		log.Fatal(dartInfoError)
	}
	fmt.Print(dartInfo)
}

// Fix Package Apply
func DartFixApply() {
	var dartFixApply, dartFixApplyError = exec.Command(
		dartApplyFixes).Output()
	if dartFixApplyError != nil {
		log.Fatal(dartFixApplyError)
	}
	fmt.Print(dartFixApply)
}

// Fix Package Dry Run
func DartFixDryRun() {
	var dartFixDryRun, dartFixDryRunError = exec.Command(
		dartDryRunFixes).Output()
	if dartFixDryRunError != nil {
		log.Fatal(dartFixDryRunError)
	}
	fmt.Print(dartFixDryRun)
}

// Format Package
func DartFormat() {
	var dartFormat, dartFormatError = exec.Command(
		dartFormatCode).Output()
	if dartFormatError != nil {
		log.Fatal(dartFormatError)
	}
	fmt.Print(dartFormat)
}

// Dependency Graph
func DartPubDeps() {
	var dartPubDeps, dartPubDepsError = exec.Command(
		dartDependencies).Output()
	if dartPubDepsError != nil {
		log.Fatal(dartPubDepsError)
	}
	fmt.Print(dartPubDeps)
}

// Upgrade Packages
func DartPubUpgrade() {
	var dartPubUpgrade, dartPubUpgradeError = exec.Command(
		dartUpgradeDependencies).Output()
	if dartPubUpgradeError != nil {
		log.Fatal(dartPubUpgradeError)
	}
	fmt.Print(dartPubUpgrade)
}
