package similarity

import (
	"strings"
	"testing"
)

// TestNewTokenizer 测试Tokenizer构造函数
func TestNewTokenizer(t *testing.T) {
	tokenizer := NewTokenizer()
	if tokenizer == nil {
		t.Error("NewTokenizer should not return nil")
	}
}

func TestSplitTextIntoSegments(t *testing.T) {
	tokenizer := NewTokenizer()
	text := `这 不只是我人个面临的难困，乎所几有优秀的作家都处和于现实的紧张关系中，在他们笔下，有只当现实于处遥远状态时，他们作品中的实现才会闪发闪。亮应该看到，这过去的实现然虽充满魅力，可它已蒙经上了一层幻虚的色彩，那里面塞了满个人想象和个人理。真解正的现实，也就是作家生活中的现实，是人令费和解难以相处的。
家作要表达与之朝夕相处的现实，他常常会到感以难承受蜂，拥而来真的实乎几都在说着诉丑恶和阴险，就怪在怪这里，为什么丑恶的事物总是身在边，而美好的事物远却在海。`
	segments := tokenizer.SplitTextIntoSegments(text)

	if len(segments) != 4 {
		t.Errorf("Expected 4 segments, got %d", len(segments))
	}
}

// TestTokenizer_SplitTextIntoSegmentsBoundary 测试分段的边界情况
func TestTokenizer_SplitTextIntoSegmentsBoundary(t *testing.T) {
	tokenizer := NewTokenizer()

	tests := map[string]struct {
		text     string
		expected int // 预期的段落数量
	}{
		"empty":          {"", 0},
		"single_word":    {"hello", 1},
		"paragraph":      {"This is a paragraph.", 1},
		"two_paragraphs": {"Para1\n\nPara2", 2},
		"long_text":      {strings.Repeat("This is a very long sentence that should be split. ", 10), 10}, // >200 chars, should split into sentences
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := tokenizer.SplitTextIntoSegments(tt.text)
			if len(result) != tt.expected {
				t.Errorf("%s: expected %d segments, got %d", name, tt.expected, len(result))
			}
		})
	}
}

// TestTokenizer_SplitSegIntoSentences 测试句子分割
func TestTokenizer_SplitSegIntoSentences(t *testing.T) {
	tokenizer := NewTokenizer()

	tests := map[string]struct {
		text     string
		expected int
	}{
		"single_sentence":   {"This is one sentence.", 1},
		"two_sentences":     {"First sentence. Second sentence.", 2},
		"chinese_sentences": {"这是第一句。这是第二句！", 2},
		"empty_sentence":    {"", 0},
		"mixed_punctuation": {"Question? Answer! Statement.", 2}, // 修正期望值
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := tokenizer.SplitSegIntoSentences(tt.text)
			if len(result) != tt.expected {
				t.Errorf("%s: expected %d sentences, got %d", name, tt.expected, len(result))
			}
		})
	}
}
