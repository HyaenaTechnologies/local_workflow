package integration

import (
	"fmt"
	"log"
	"os/exec"
)

// Go Integration Strings
var (
	goDependencyGraph = "go mod graph"
	goDocument        = "go doc"
	goEnvironment     = "go env"
	goFixCode         = "go fix ./lib"
	goFormatCode      = "go fmt ./lib"
	goGetDependencies = "go mod download"
	goModules         = "go list -m"
	goPackages        = "go list"
	goVerifyModules   = "go mod verify"
	goVersionNumber   = "go version"
	goVetCode         = "go vet ./lib"
	goWorkSync        = "go work sync"
	goWorkspace       = "go work use ./"
)

// Document Package
func GoDoc() {
	var goDoc, goDocError = exec.Command(
		goDocument).Output()
	if goDocError != nil {
		log.Fatal(goDocError)
	}
	fmt.Print(goDoc)
}

// Environment Information
func GoEnv() {
	var goEnv, goEnvError = exec.Command(
		goEnvironment).Output()
	if goEnvError != nil {
		log.Fatal(goEnvError)
	}
	fmt.Print(goEnv)
}

// Fix Package
func GoFix() {
	var goFix, goFixError = exec.Command(
		goFixCode).Output()
	if goFixError != nil {
		log.Fatal(goFixError)
	}
	fmt.Print(goFix)
}

// Format Package
func GoFmt() {
	var goFmt, goFmtError = exec.Command(
		goFormatCode).Output()
	if goFmtError != nil {
		log.Fatal(goFmtError)
	}
	fmt.Print(goFmt)
}

// List Packages
func GoList() {
	var goList, goListError = exec.Command(
		goPackages).Output()
	if goListError != nil {
		log.Fatal(goListError)
	}
	fmt.Print(goList)
}

// List Modules
func GoListModules() {
	var goListModules, goListModulesError = exec.Command(
		goModules).Output()
	if goListModulesError != nil {
		log.Fatal(goListModulesError)
	}
	fmt.Print(goListModules)
}

// Download Packages
func GoModDownload() {
	var goModDownload, goModDownloadError = exec.Command(
		goGetDependencies).Output()
	if goModDownloadError != nil {
		log.Fatal(goModDownloadError)
	}
	fmt.Print(goModDownload)
}

// Dependency Graph
func GoModGraph() {
	var goModGraph, goModGraphError = exec.Command(
		goDependencyGraph).Output()
	if goModGraphError != nil {
		log.Fatal(goModGraphError)
	}
	fmt.Print(goModGraph)
}

// Verify Module
func GoModVerify() {
	var goModVerify, goModVerifyError = exec.Command(
		goVerifyModules).Output()
	if goModVerifyError != nil {
		log.Fatal(goModVerifyError)
	}
	fmt.Print(goModVerify)
}

// Go Version
func GoVersion() {
	var goVersion, goVersionError = exec.Command(
		goVersionNumber).Output()
	if goVersionError != nil {
		log.Fatal(goVersionError)
	}
	fmt.Print(goVersion)
}

// Vet Package
func GoVet() {
	var goVet, goVetError = exec.Command(
		goVetCode).Output()
	if goVetError != nil {
		log.Fatal(goVetError)
	}
	fmt.Print(goVet)
}

// Synchronize Workspace
func GoWorkSynchronize() {
	var goSynchronize, goSynchronizeError = exec.Command(
		goWorkSync).Output()
	if goSynchronizeError != nil {
		log.Fatal(goSynchronizeError)
	}
	fmt.Print(goSynchronize)
}

// Use Workspace
func GoWorkUse() {
	var goWorkUse, goWorkUseError = exec.Command(
		goWorkspace).Output()
	if goWorkUseError != nil {
		log.Fatal(goWorkUseError)
	}
	fmt.Print(goWorkUse)
}
