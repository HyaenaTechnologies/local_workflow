package command

import (
	"fmt"
	"log"
	"os/exec"
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

// Document Package
func GoDoc() {
	var goDoc, goDocError = exec.Command(
		goDocument).Output()
	if goDocError != nil {
		log.Fatal(goDocError)
	}
	fmt.Print(goDoc)
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
		GoIntegration()
	}
	if *goDocFlag == true {
		GoDoc()
	}
	if *goTestFlag != "./test" {
		GoTest()
	}
}

func GoIntegration() {
	var goIntegrate, goIntegrateError = exec.Command(
		goIntegration).Output()
	if goIntegrateError != nil {
		log.Fatal(goIntegrateError)
	}
	fmt.Print(goIntegrate)
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
