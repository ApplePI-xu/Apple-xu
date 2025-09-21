package utils

import "fmt"

type Formatter struct{}

func NewFormatter() *Formatter {
	return &Formatter{}
}

func (f *Formatter) FmtSimilarityAsPct(similarity float64) string {
	percentage := similarity * 100
	return fmt.Sprintf("%.2f%%", percentage)
}
