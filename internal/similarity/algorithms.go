package similarity

// Algorithms 相似度算法集合
type Algorithms struct{}

// NewAlgorithms 创建新的算法实例
func NewAlgorithms() *Algorithms {
	return &Algorithms{}
}

// EditDistance 计算编辑距离（Levenshtein距离）
func (a *Algorithms) EditDistance(str1, str2 string) int {
	runes1 := []rune(str1)
	runes2 := []rune(str2)
	m, n := len(runes1), len(runes2)

	// 创建DP表
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 初始化
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	// 填充DP表
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if runes1[i-1] == runes2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, min(dp[i][j-1]+1, dp[i-1][j-1]+1))
			}
		}
	}

	return dp[m][n]
}

func (a *Algorithms) CalcCharSimilarity(str1, str2 string) float64 {
	editDit := a.EditDistance(str1, str2)
	maxlen := max(len([]rune(str1)), len([]rune(str2)))
	if maxlen == 0 {
		return 1.0
	}
	return max(0, 1.0-float64(editDit)/float64(maxlen))

}

// CalcSegSimilarity 分段计算长文本的相似度
func (a *Algorithms) CalcSegSimilarity(segs1, segs2 []string) float64 {
	if len(segs1) == 0 && len(segs2) == 0 {
		return 1.0
	}
	if len(segs1) == 0 || len(segs2) == 0 {
		return 0.0
	}

	// 计算每个段落的最佳匹配相似度
	totalSimilarity := 0.0
	totalWeight := 0.0

	//seg1为原文，segs2为比较版，使用seg2与seg1进行比较
	for _, seg2 := range segs2 {
		maxSimilarity := 0.0
		seglen := float64(len([]rune(seg2)))

		if seglen == 0 {
			continue
		}

		for _, seg1 := range segs1 {
			similarity := a.CalcCharSimilarity(seg1, seg2)
			if similarity > maxSimilarity {
				maxSimilarity = similarity
			}
		}

		totalSimilarity += maxSimilarity * seglen
		totalWeight += seglen
	}

	return totalSimilarity / totalWeight
}
