package similarity

import (
	"testing"
)

// abs 计算浮点数绝对值
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// TestNewAlgorithms 测试构造函数
func TestNewAlgorithms(t *testing.T) {
	algorithms := NewAlgorithms()
	if algorithms == nil {
		t.Error("NewAlgorithms should not return nil")
	}
}

// TestEditDistance 测试编辑距离计算
func TestEditDistance(t *testing.T) {
	algorithms := NewAlgorithms()

	tests := map[string]struct {
		input1 string
		input2 string
		output int
	}{
		// 经典测试用例
		"kitten_sitting":      {"kitten", "sitting", 3},
		"flaw_lawn":           {"flaw", "lawn", 2},
		"intention_execution": {"intention", "execution", 5},

		// 边界情况
		"both_empty":   {"", "", 0},
		"one_empty":    {"", "hello", 5},
		"other_empty":  {"hello", "", 5},
		"same_strings": {"hello", "hello", 0},
		"single_char":  {"a", "b", 1},
		"unicode":      {"你好", "你们", 1},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := algorithms.EditDistance(tt.input1, tt.input2)
			if result != tt.output {
				t.Errorf("expected %d, got %d", tt.output, result)
			}
		})
	}
}

// TestCalcCharSimilarity 测试字符相似度计算
func TestCalcCharSimilarity(t *testing.T) {
	algorithms := NewAlgorithms()

	tests := map[string]struct {
		input1 string
		input2 string
		output float64
	}{
		// 基于编辑距离的相似度测试
		"kitten_sitting":      {"kitten", "sitting", 0.571429},
		"flaw_lawn":           {"flaw", "lawn", 0.500000},
		"intention_execution": {"intention", "execution", 0.444444},

		// 边界情况
		"both_empty":       {"", "", 1.0},
		"one_empty":        {"", "hello", 0.0},
		"other_empty":      {"hello", "", 0.0},
		"same_strings":     {"hello", "hello", 1.0},
		"single_char_same": {"a", "a", 1.0},
		"single_char_diff": {"a", "b", 0.0},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := algorithms.CalcCharSimilarity(tt.input1, tt.input2)
			if abs(result-tt.output) > 0.000001 {
				t.Errorf("expected %f, got %f", tt.output, result)
			}
		})
	}
}

// TestCalcSegSimilarity 测试分段相似度计算
func TestCalcSegSimilarity(t *testing.T) {
	algorithms := NewAlgorithms()

	tests := map[string]struct {
		segs1    []string
		segs2    []string
		expected float64
	}{
		// 功能测试
		"chinese_text":   {[]string{"你好", "今天是晴天"}, []string{"你啊哈好", "今天天气晴朗"}, 0.5},
		"identical_segs": {[]string{"flaw", "lawn"}, []string{"flaw", "lawn"}, 1.0},
		"exact_match":    {[]string{"intention", "execution"}, []string{"intention", "execution"}, 1.0},

		// 边界情况
		"both_empty":        {[]string{}, []string{}, 1.0},
		"first_empty":       {[]string{}, []string{"hello"}, 0.0},
		"second_empty":      {[]string{"hello"}, []string{}, 0.0},
		"empty_segments":    {[]string{"", ""}, []string{"", ""}, 1.0},
		"one_empty_segment": {[]string{"hello", ""}, []string{"hello"}, 1.0},
		"multiple_segments": {[]string{"hello", "world"}, []string{"hello", "world"}, 1.0},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := algorithms.CalcSegSimilarity(tt.segs1, tt.segs2)
			if abs(result-tt.expected) > 0.000001 {
				t.Errorf("%s: expected %f, got %f", name, tt.expected, result)
			}
		})
	}
}
