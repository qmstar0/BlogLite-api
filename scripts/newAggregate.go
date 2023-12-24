package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 检查是否提供了命令行参数
	if len(os.Args) < 2 {
		fmt.Println("请提供文件夹名作为命令行参数。")
		return
	}

	// 获取命令行参数作为文件夹名
	folderName := os.Args[1]

	// 创建主文件夹
	err := os.Mkdir(folderName, os.ModePerm)
	if err != nil {
		fmt.Printf("创建文件夹失败：%v\n", err)
		return
	}

	// 切换到主文件夹
	err = os.Chdir(folderName)
	if err != nil {
		fmt.Printf("切换到文件夹失败：%v\n", err)
		return
	}

	// 创建子文件夹
	subFolders := []string{
		"events",
		"repository",
		"entity",
		"service",
		"valueobject",
		"commands",
	}
	for _, subFolder := range subFolders {
		err := os.Mkdir(subFolder, os.ModePerm)
		if err != nil {
			fmt.Printf("创建文件夹失败：%v\n", err)
			return
		}
	}

	// 在 entity 文件夹下创建 (folderName).go 文件
	entityFile, err := os.Create(filepath.Join("entity", folderName+".go"))
	if err != nil {
		fmt.Printf("创建 entity 文件失败：%v\n", err)
		return
	}
	defer entityFile.Close()
	_, err = entityFile.WriteString("package entity\n")
	if err != nil {
		fmt.Printf("写入entity/%s.go失败", folderName)
	}

	// 在 repository 文件夹下创建 (folderName).go 文件
	repositoryFile, err := os.Create(filepath.Join("repository", folderName+".go"))
	if err != nil {
		fmt.Printf("创建 repository 文件失败：%v\n", err)
		return
	}
	defer repositoryFile.Close()
	_, err = repositoryFile.WriteString("package repository\n")
	if err != nil {
		fmt.Printf("写入repository/%s.go失败", folderName)
	}

	fmt.Println("文件夹和文件创建完成。")
}
