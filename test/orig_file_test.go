package main

import (
	"os"
	"psgCheck/internal/similarity"
	"testing"
)

func TestOrigAndVariantsFromFile(t *testing.T) {
	basePath := "d:/goproject/src/softwareEng/3123004185/test/test_data/"
	origPath := basePath + "orig.txt"
	cases := []struct {
		file     string
		expectLo float64
		expectHi float64
	}{
		{"orig_0.8_add.txt", 0.6, 0.9},
		{"orig_0.8_del.txt", 0.6, 0.9},
		{"orig_0.8_dis_1.txt", 0.6, 0.9},
		{"orig_0.8_dis_10.txt", 0.6, 0.85},
		{"orig_0.8_dis_15.txt", 0.5, 0.8},
	}

	origData, err := os.ReadFile(origPath)
	if err != nil {
		t.Fatalf("无法读取原文文件: %v", err)
	}

	checker := similarity.NewChecker()

	for _, c := range cases {
		c := c // 避免闭包变量问题
		t.Run(c.file, func(t *testing.T) {
			plagData, err := os.ReadFile(basePath + c.file)
			if err != nil {
				t.Errorf("无法读取抄袭文件 %s: %v", c.file, err)
				return
			}
			sim := checker.CalculateSimilarity(string(origData), string(plagData))
			if sim < c.expectLo || sim > c.expectHi {
				t.Errorf("%s: 相似度 %.4f 不在期望范围 [%.2f, %.2f]", c.file, sim, c.expectLo, c.expectHi)
			}
		})
	}
}
