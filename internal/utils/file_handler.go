package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// FileHandler 文件处理器
type FileHandler struct{}

// NewFileHandler 创建新的文件处理器
func NewFileHandler() *FileHandler {
	return &FileHandler{}
}

func (fh *FileHandler) ReadFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("无法打开文件 %s: %v", filePath, err)
	}
	defer file.Close()

	var content strings.Builder
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		content.WriteString(strings.TrimSpace(line)) // 去除行首尾空格
		content.WriteString(" ")
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("读取文件 %s 时出错: %v", filePath, err)
	}
	// 清理多余的空格
	return content.String(), nil
}

func (fh *FileHandler) WriteFile(filePath string, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("无法创建输出文件 %s: %v", filePath, err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("写入文件 %s 时出错: %v", filePath, err)
	}

	return nil
}

func (fh *FileHandler) PreprocessText(text string) string {
	text = strings.TrimSpace(text) // 去除首尾空格
	// 替换多个空格为单个空格
	return strings.Join(strings.Fields(text), " ")
}
