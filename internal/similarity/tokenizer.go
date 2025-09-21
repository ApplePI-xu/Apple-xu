package similarity

import (
	"psgCheck/internal/utils"
	"regexp"
	"strings"
)

// Tokenizer 文本分词器
type Tokenizer struct{}

// NewTokenizer 创建新的分词器
func NewTokenizer() *Tokenizer {
	return &Tokenizer{}
}

func (t *Tokenizer) SplitTextIntoSegments(text string) []string {
	// 预处理文本
	fh := utils.NewFileHandler()
	text = fh.PreprocessText(text)

	// 统一换行符
	text = strings.ReplaceAll(text, "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")

	// 按段落分割（双换行符）
	paragraphs := strings.Split(text, "\n\n")

	var segments []string
	for _, paragraph := range paragraphs {
		if len(paragraph) == 0 {
			continue
		}
		// 如果段落太长，按句子进一步分割
		if len(paragraph) > 200 {
			sentences := t.SplitSegIntoSentences(paragraph)
			segments = append(segments, sentences...)
		} else {
			segments = append(segments, paragraph)
		}
	}

	return segments
}

func (t *Tokenizer) SplitSegIntoSentences(text string) []string {
	regexp := regexp.MustCompile(`[。！？!."“”]+`)
	sentences := regexp.Split(text, -1)

	var result []string
	for _, sentence := range sentences {
		if len(sentence) == 0 {
			continue
		}
		result = append(result, strings.TrimSpace(sentence))
	}
	return result
}
