package similarity

import (
	"testing"
)

// TestNewChecker 测试Checker构造函数
func TestNewChecker(t *testing.T) {
	checker := NewChecker()

	// 首先确保构造函数返回了有效对象
	if checker == nil {
		t.Fatal("NewChecker should not return nil")
	}

	// 然后检查对象的字段是否正确初始化
	if checker.tokenizer == nil {
		t.Error("Checker tokenizer should not be nil")
	}
	if checker.algorithms == nil {
		t.Error("Checker algorithms should not be nil")
	}
}

// TestChecker_CalculateSimilarity 测试Checker的相似度计算
func TestChecker_CalculateSimilarity(t *testing.T) {
	checker := NewChecker()

	tests := map[string]struct {
		original      string
		plagiarized   string
		minSimilarity float64 // 最小预期相似度
		maxSimilarity float64 // 最大预期相似度
	}{
		"identical":  {"hello world", "hello world", 0.99, 1.01},
		"different":  {"hello world", "goodbye world", 0.0, 0.6},
		"empty_both": {"", "", 0.99, 1.01},
		"empty_one":  {"", "hello", 0.0, 0.01},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := checker.CalculateSimilarity(tt.original, tt.plagiarized)
			if result < tt.minSimilarity || result > tt.maxSimilarity {
				t.Errorf("%s: expected similarity between %f and %f, got %f",
					name, tt.minSimilarity, tt.maxSimilarity, result)
			}
		})
	}
}
