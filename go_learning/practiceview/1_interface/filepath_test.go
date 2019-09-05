package __interface

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// 遍历文件树
func TestFilePath(t *testing.T) {
	filepath.Walk("D:\\翻墙软件", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			fmt.Println(path)
		}
		return err
	})
}