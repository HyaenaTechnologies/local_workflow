package command

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type GoCommands struct {
	darwinAMD64  string
	darwinARM64  string
	document     string
	integrate    []string
	linuxAMD64   string
	linuxARM64   string
	test         string
	windowsAMD64 string
	windowsARM64 string
}

// Go Command Strings
var (
	goDarwinAMD64Build = GoCommands{
		darwinAMD64: "GOOS=darwin GOARCH=amd64 go build -o",
	}
	goDarwinARM64Build = GoCommands{
		darwinARM64: "GOOS=darwin GOARCH=arm64 go build -o",
	}
	goDocument = GoCommands{
		document: "go doc",
	}
	goIntegration = strings.Join(
		goIntegrate.integrate,
		" && ",
	)
	goLinuxAMD64Buid = GoCommands{
		linuxAMD64: "GOOS=linux GOARCH=amd64 go build -o",
	}
	goLinuxARM64Buid = GoCommands{
		linuxARM64: "GOOS=linux GOARCH=arm64 go build -o",
	}
	goRunTest = GoCommands{
		test: "go test -v",
	}
	goWindowsAMD64Build = GoCommands{
		windowsAMD64: "GOOS=windows GOARCH=amd64 go build -o",
	}
	goWindowsARM64Build = GoCommands{
		windowsARM64: "GOOS=windows GOARCH=arm64 go build -o",
	}
)

// Go Integration Commands
var goIntegrate = GoCommands{
	integrate: []string{
		"go env",
		"go fix ./lib",
		"go fmt ./lib",
		"go list",
		"go list -m",
		"go mod download",
		"go mod graph",
		"go mod verify",
		"go version",
		"go vet ./lib",
		"go work sync",
		"go work use ./"},
}

// Build Package Darwin AMD64
func (darwinAMD64Build GoCommands) GoDarwinAMD64() {
	var goDarwinAMD64, goDarwinAMD64Error = exec.Command(
		darwinAMD64Build.darwinAMD64,
		*goBuildFlag,
	).Output()
	if goDarwinAMD64Error != nil {
		log.Fatal(goDarwinAMD64Error)
	}
	fmt.Print(goDarwinAMD64)
}

// Build Package Darwin ARM64
func (darwinARM64Build GoCommands) GoDarwinARM64() {
	var goDarwinARM64, goDarwinARM64Error = exec.Command(
		darwinARM64Build.darwinARM64,
		*goBuildFlag,
	).Output()
	if goDarwinARM64Error != nil {
		log.Fatal(goDarwinARM64Error)
	}
	fmt.Print(goDarwinARM64)
}

// Build Package Linux AMD64
func (linuxAMD64Build GoCommands) GoLinuxAMD64() {
	var goLinuxAMD64, goLinuxAMD64Error = exec.Command(
		linuxAMD64Build.linuxAMD64,
		*goBuildFlag,
	).Output()
	if goLinuxAMD64Error != nil {
		log.Fatal(goLinuxAMD64Error)
	}
	fmt.Print(goLinuxAMD64)
}

// Build Package Linux ARM64
func (linuxARM64Build GoCommands) GoLinuxARM64() {
	var goLinuxARM64, goLinuxARM64Error = exec.Command(
		linuxARM64Build.linuxARM64,
		*goBuildFlag,
	).Output()
	if goLinuxARM64Error != nil {
		log.Fatal(goLinuxARM64Error)
	}
	fmt.Print(goLinuxARM64)
}

// Build Package Windows AMD64
func (windowsAMD64Build GoCommands) GoWindowsAMD64() {
	var goWindowsAMD64, goWindowsAMD64Error = exec.Command(
		windowsAMD64Build.windowsAMD64,
		*goBuildFlag,
	).Output()
	if goWindowsAMD64Error != nil {
		log.Fatal(goWindowsAMD64Error)
	}
	fmt.Print(goWindowsAMD64)
}

// Build Package Windows ARM64
func (windowsARM64Build GoCommands) GoWindowsARM64() {
	var goWindowsARM64, goWindowsARM64Error = exec.Command(
		windowsARM64Build.windowsARM64,
		*goBuildFlag,
	).Output()
	if goWindowsARM64Error != nil {
		log.Fatal(goWindowsARM64Error)
	}
	fmt.Print(goWindowsARM64)
}

// Document Package
func (documentation GoCommands) GoDoc() {
	var goDoc, goDocError = exec.Command(
		documentation.document,
	).Output()
	if goDocError != nil {
		log.Fatal(goDocError)
	}
	fmt.Print(goDoc)
}

// Go Command Flag Execution
func GoFlagExecute() {
	if *goBuildFlag != "" {
		goDarwinAMD64Build.GoDarwinAMD64()
		goDarwinARM64Build.GoDarwinARM64()
		goLinuxAMD64Buid.GoLinuxAMD64()
		goLinuxARM64Buid.GoLinuxARM64()
		goWindowsAMD64Build.GoWindowsAMD64()
		goWindowsARM64Build.GoWindowsARM64()
	}
	if *goIntegrationFlag == true {
		GoIntegration()
	}
	if *goDocFlag == true {
		goDocument.GoDoc()
	}
	if *goTestFlag != "./test" {
		goRunTest.GoTest()
	}
}

func GoIntegration() {
	var goIntegrate, goIntegrateError = exec.Command(
		goIntegration,
	).Output()
	if goIntegrateError != nil {
		log.Fatal(goIntegrateError)
	}
	fmt.Print(goIntegrate)
}

// Run Test
func (runTest GoCommands) GoTest() {
	var goTest, goTestError = exec.Command(
		runTest.test,
		*goTestFlag,
	).Output()
	if goTestError != nil {
		log.Fatal(goTestError)
	}
	fmt.Print(goTest)
}
