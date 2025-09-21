package main

import (
	"fmt"
	"os"
	"psgCheck/internal/similarity"
	"psgCheck/internal/utils"
	"time"
)

func main() {
	start := time.Now()

	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "参数错误: %d\n", len(os.Args)-1)
		fmt.Fprintf(os.Stderr, "使用方法: %s <原文文件路径> <抄袭版文件路径> <输出文件路径>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "示例: %s /path/to/orig.txt /path/to/plagiarized.txt /path/to/result.txt\n", os.Args[0])
		os.Exit(1)
	}

	originalFilePath := os.Args[1]
	plagiarizedFilePath := os.Args[2]
	outputFilePath := os.Args[3]

	fileHandler := utils.NewFileHandler()
	formatter := utils.NewFormatter()
	checker := similarity.NewChecker()

	// 读取文件
	oriText, err := fileHandler.ReadFile(originalFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "读取原文文件失败: %v\n", err)
		os.Exit(1)
	}
	plagText, err := fileHandler.ReadFile(plagiarizedFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "读取抄袭版文件失败: %v\n", err)
		os.Exit(1)
	}

	// 计算相似度
	similarity := checker.CalculateSimilarity(oriText, plagText)

	// 格式化结果
	result := formatter.FmtSimilarityAsPct(similarity)

	// 写入结果文件
	msg := originalFilePath + "   " + plagiarizedFilePath + "   " + result
	err = fileHandler.WriteFile(outputFilePath, msg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "写入结果文件失败: %v\n", err)
		os.Exit(1)
	}

	// 输出结果
	fmt.Printf("查重完成！相似度: %s\n", result)
	fmt.Printf("结果已保存到: %s\n", outputFilePath)

	elapsed := time.Since(start)
	fmt.Println("耗时:", elapsed)
}
