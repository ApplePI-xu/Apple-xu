package main

import (
	"psgCheck/internal/similarity"
	"testing"
)

func TestIntegration_TableDriven(t *testing.T) {
	type testCase struct {
		name     string
		orig     string
		plag     string
		expectLo float64
		expectHi float64
	}

	tests := []testCase{
		{
			name:     "完全相同",
			orig:     "hello world\n第二段。",
			plag:     "hello world\n第二段。",
			expectLo: 1.0,
			expectHi: 1.0,
		},
		{
			name:     "部分相同",
			orig:     "hello world\n第二段。",
			plag:     "hello norld\n第二段!",
			expectLo: 0.8,
			expectHi: 1.0,
		},
		{
			name:     "完全不同",
			orig:     "abc def",
			plag:     "xyz 123",
			expectLo: 0.0,
			expectHi: 0.2,
		},
		{
			name:     "全为空",
			orig:     "",
			plag:     "",
			expectLo: 1.0,
			expectHi: 1.0,
		},
		{
			name:     "部分为空-原文空",
			orig:     "",
			plag:     "HelloWorld",
			expectLo: 0.0,
			expectHi: 0.0,
		},
		{
			name:     "部分为空-抄袭空",
			orig:     "HelloWorld",
			plag:     "",
			expectLo: 0.0,
			expectHi: 0.0,
		},
		{
			name:     "英文大小写不同",
			orig:     "Hello World",
			plag:     "hello world",
			expectLo: 0.95,
			expectHi: 1.0,
		},
		{
			name:     "中文符号不同",
			orig:     "你好，世界后续内容！",
			plag:     "你好  世界后续内容？",
			expectLo: 0.8,
			expectHi: 1.0,
		},
		{
			name:     "多段落部分相同",
			orig:     "第一段内容。\n\n第二段内容。\n\n第三段内容。",
			plag:     "第一段内容。\n\n不同内容。\n\n第三段内容。",
			expectLo: 0.6,
			expectHi: 0.9,
		},
		{
			name:     "多段落完全不同",
			orig:     "A。\n\nB。\n\nCDEFG。",
			plag:     "X。\n\nY。\n\nZHIJK。",
			expectLo: 0.0,
			expectHi: 0.3,
		},
		{
			name:     "长文本高相似",
			orig:     "Go is an open source programming language.\nIt is easy to learn.",
			plag:     "Go is an open source programming language.\nIt is easy to use.",
			expectLo: 0.7,
			expectHi: 1.0,
		},
		{
			name:     "长文本低相似",
			orig:     "Go is an open source programming language.\nIt is easy to learn.",
			plag:     "Python is a popular scripting language.\nIt is widely used.",
			expectLo: 0.0,
			expectHi: 0.5,
		},
	}

	checker := similarity.NewChecker()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sim := checker.CalculateSimilarity(tc.orig, tc.plag)
			if sim < tc.expectLo || sim > tc.expectHi {
				t.Errorf("%s: 相似度 %.4f 不在期望范围 [%.2f, %.2f]", tc.name, sim, tc.expectLo, tc.expectHi)
			}
		})
	}
}
