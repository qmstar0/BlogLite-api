package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("你需要输入以下参数:\n" +
			"\t1. server name, 程序会寻找/api/protobuf/$server.proto\n")
		return
	}

	tool := "protoc"
	server := os.Args[1]

	// 定义 oapi-codegen 命令及其参数
	cmd1 := exec.Command(tool,
		"-I", "api/protobuf", fmt.Sprintf("%s.proto", server),
		"--go_out=internal/common/proto",
	)
	err := cmd1.Run()
	if err != nil {
		fmt.Println("执行发生错误", err, cmd1.String())
		return
	}
}
