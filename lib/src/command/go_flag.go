package command

import (
	"flag"
	"fmt"
	"local_workflow/lib/src/integration"
	"log"
	"os"
	"os/exec"
)

// Go Command Strings
var (
	goDarwinAMD64Build  = "GOOS=darwin GOARCH=amd64 go build -o"
	goDarwinARM64Build  = "GOOS=darwin GOARCH=arm64 go build -o"
	goLinuxAMD64Buid    = "GOOS=linux GOARCH=amd64 go build -o"
	goLinuxARM64Buid    = "GOOS=linux GOARCH=arm64 go build -o"
	goRunTest           = "go test -v"
	goWindowsAMD64Build = "GOOS=windows GOARCH=amd64 go build -o"
	goWindowsARM64Build = "GOOS=windows GOARCH=arm64 go build -o"
)

// Go Commands and Flags
var (
	goCommand      = flag.NewFlagSet("go", flag.ExitOnError)
	goCommandError = goCommand.Parse(os.Args[2:])
	goBuildFlag    = goCommand.String(
		"build",
		"",
		"Build Package: --build Path: <String>")
	goDocFlag = goCommand.Bool(
		"doc",
		false,
		"Document Package: --doc <Bool>")
	goIntegrationFlag = goCommand.Bool(
		"integration",
		false,
		"Continuous Integration: --integration <Bool>")
	goTestFlag = goCommand.String(
		"test",
		"./test",
		"Run Test: --test Path: <String>")
)

// Build Package Darwin AMD64
func GoBuildDarwinAMD64() {
	var goDarwinAMD64, goDarwinAMD64Error = exec.Command(
		goDarwinAMD64Build,
		*goBuildFlag).Output()
	if goDarwinAMD64Error != nil {
		log.Fatal(goDarwinAMD64Error)
	}
	fmt.Print(goDarwinAMD64)
}

// Build Package Darwin ARM64
func GoBuildDarwinARM64() {
	var goDarwinARM64, goDarwinARM64Error = exec.Command(
		goDarwinARM64Build,
		*goBuildFlag).Output()
	if goDarwinARM64Error != nil {
		log.Fatal(goDarwinARM64Error)
	}
	fmt.Print(goDarwinARM64)
}

// Build Package Linux AMD64
func GoBuildLinuxAMD64() {
	var goLinuxAMD64, goLinuxAMD64Error = exec.Command(
		goLinuxAMD64Buid,
		*goBuildFlag).Output()
	if goLinuxAMD64Error != nil {
		log.Fatal(goLinuxAMD64Error)
	}
	fmt.Print(goLinuxAMD64)
}

// Build Package Linux ARM64
func GoBuildLinuxARM64() {
	var goLinuxARM64, goLinuxARM64Error = exec.Command(
		goLinuxARM64Buid,
		*goBuildFlag).Output()
	if goLinuxARM64Error != nil {
		log.Fatal(goLinuxARM64Error)
	}
	fmt.Print(goLinuxARM64)
}

// Build Package Windows AMD64
func GoBuildWindowsAMD64() {
	var goWindowsAMD64, goWindowsAMD64Error = exec.Command(
		goWindowsAMD64Build,
		*goBuildFlag).Output()
	if goWindowsAMD64Error != nil {
		log.Fatal(goWindowsAMD64Error)
	}
	fmt.Print(goWindowsAMD64)
}

// Build Package Windows ARM64
func GoBuildWindowsARM64() {
	var goWindowsARM64, goWindowsARM64Error = exec.Command(
		goWindowsARM64Build,
		*goBuildFlag).Output()
	if goWindowsARM64Error != nil {
		log.Fatal(goWindowsARM64Error)
	}
	fmt.Print(goWindowsARM64)
}

// Go Command Flag Execution
func GoFlagExecute() {
	if *goBuildFlag != "" {
		GoBuildDarwinAMD64()
		GoBuildDarwinARM64()
		GoBuildLinuxAMD64()
		GoBuildLinuxARM64()
		GoBuildWindowsAMD64()
		GoBuildWindowsARM64()
	}
	if *goIntegrationFlag == true {
		integration.GoEnv()
		integration.GoFix()
		integration.GoFmt()
		integration.GoList()
		integration.GoListModules()
		integration.GoModDownload()
		integration.GoModGraph()
		integration.GoModVerify()
		integration.GoVersion()
		integration.GoVet()
		integration.GoWorkSynchronize()
		integration.GoWorkUse()
	}
	if *goDocFlag == true {
		integration.GoDoc()
	}
	if *goTestFlag != "" {
		GoTest()
	}
}

// Run Test
func GoTest() {
	var goTest, goTestError = exec.Command(
		goRunTest,
		*goTestFlag).Output()
	if goTestError != nil {
		log.Fatal(goTestError)
	}
	fmt.Print(goTest)
}
