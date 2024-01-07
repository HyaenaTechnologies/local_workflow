package command

import (
	"flag"
	"os"
	"strings"
)

// Go Commands and Flags
var (
	goCommand      = flag.NewFlagSet("go", flag.ExitOnError)
	goCommandError = goCommand.Parse(os.Args[2:])
	goBuildFlag    = goCommand.String(
		"build",
		"",
		"Build Package: --build Path: <String>",
	)
	goDocFlag = goCommand.Bool(
		"doc",
		false,
		"Document Package: --doc <Bool>",
	)
	goIntegrationFlag = goCommand.Bool(
		"integration",
		false,
		"Continuous Integration: --integration <Bool>",
	)
	goTestFlag = goCommand.String(
		"test",
		"./test",
		"Run Test: --test Path: <String>",
	)
)

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
		goIntegrateCommands.integrate,
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
var goIntegrateCommands = GoCommands{
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
