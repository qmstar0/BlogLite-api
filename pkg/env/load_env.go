package env

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Load(path ...string) {
	if len(path) == 0 {
		path = append(path, ".env")
	}
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
