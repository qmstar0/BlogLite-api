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
	dirs, err := os.ReadDir(rootDir)
	if err != nil {
		panic(err)
	}
	for i := range dirs {
		dir := filepath.Join(rootDir, dirs[i].Name())
		fmt.Println("handler dir:", dir)
		runCmd(dir)
	}
}

func runCmd(dir string) {
	openapiFile := filepath.Join(dir, openapiFileName)

	fs, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for i := range fs {
		name := fs[i].Name()
		if name == openapiFileName {
			continue
		} else {
			file := filepath.Join(dir, name)
			cmd := exec.Command(
				openapiGenerateTool,
				"-config",
				file,
				openapiFile,
			)
			fmt.Println("执行命令:", cmd.String())
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			err = cmd.Run()
			if err != nil {
				panic(err)
			}
		}
	}
}
