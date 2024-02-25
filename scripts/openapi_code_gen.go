package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("你需要输入以下参数:\n" +
			"\t1. openapi规范的yaml文件\n" +
			"\t2. package name, 该选项决定程序生成的文件的包名\n" +
			"\t3. output dir, 该选项决定输出目录\n")
		return
	}

	tool := "oapi-codegen"
	yaml := os.Args[1]
	pkgName := os.Args[2]
	outputDir := os.Args[3]

	for strings.HasSuffix(outputDir, "/") {
		outputDir = outputDir[:len(outputDir)-1]
	}

	// 定义 oapi-codegen 命令及其参数
	//cmd1 := exec.Command(tool,
	//	"-package", pkgName,
	//	"-generate", "types",
	//	"-o", fmt.Sprintf("%s/openapi_types.gen.go", outputDir),
	//	fmt.Sprintf("api/openapi/%s.yaml", server))

	cmd2 := exec.Command(tool,
		"-package", pkgName,
		"-generate", "chi-server",
		"-o", fmt.Sprintf("%s/openapi_api.gen.go", outputDir),
		yaml)

	err := cmd2.Run()
	if err != nil {
		fmt.Println("执行发生错误", err, cmd2.String())
		return
	}
}
