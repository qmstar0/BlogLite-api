package env

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func init() {
	load(".env")
}

func Load(path ...string) {
	for i := range path {
		load(path[i])
	}
}

func load(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		_ = file.Close()
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key, value := parts[0], parts[1]
			_ = os.Setenv(key, value)
		}
	}
}
