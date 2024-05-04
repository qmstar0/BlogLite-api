package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	openapiFileName     = "openapi.json"
	rootDir             = "api/openapi"
	openapiGenerateTool = "oapi-codegen.exe"
)

func main() {
	runCmd(rootDir)
}

func runCmd(dir string) {
	openapiFile := filepath.Join(dir, openapiFileName)

	file := filepath.Join(dir, "config.yaml")
	cmd := exec.Command(
		openapiGenerateTool,
		"-config",
		file,
		openapiFile,
	)
	fmt.Println("执行命令:", cmd.String())
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
