package similarity

type Checker struct {
	tokenizer  *Tokenizer
	algorithms *Algorithms
}

func NewChecker() *Checker {
	return &Checker{
		tokenizer:  NewTokenizer(),
		algorithms: NewAlgorithms(),
	}
}

func (c *Checker) CalculateSimilarity(original, plagiarized string) float64 {
	oriSegs := c.tokenizer.SplitTextIntoSegments(original)
	plaSegs := c.tokenizer.SplitTextIntoSegments(plagiarized)

	return c.algorithms.CalcSegSimilarity(oriSegs, plaSegs)
}
