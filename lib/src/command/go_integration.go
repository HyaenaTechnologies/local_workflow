package command

import (
	"fmt"
	"log"
	"os/exec"
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

// Build Package Darwin AMD64
func (darwinAMD64Build GoCommands) GoBuildDarwinAMD64() {
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
func (darwinARM64Build GoCommands) GoBuildDarwinARM64() {
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
func (linuxAMD64Build GoCommands) GoBuildLinuxAMD64() {
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
func (linuxARM64Build GoCommands) GoBuildLinuxARM64() {
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
func (windowsAMD64Build GoCommands) GoBuildWindowsAMD64() {
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
func (windowsARM64Build GoCommands) GoBuildWindowsARM64() {
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
		goDarwinAMD64Build.GoBuildDarwinAMD64()
		goDarwinARM64Build.GoBuildDarwinARM64()
		goLinuxAMD64Buid.GoBuildLinuxAMD64()
		goLinuxARM64Buid.GoBuildLinuxARM64()
		goWindowsAMD64Build.GoBuildWindowsAMD64()
		goWindowsARM64Build.GoBuildWindowsARM64()
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
