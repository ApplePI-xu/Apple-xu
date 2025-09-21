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
