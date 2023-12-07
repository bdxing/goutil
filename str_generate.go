package goutil

import "math/rand"

// StrGenerateRandom 从种子 allowedChar 字符串里随机生成 n 个字符组成字符串。
// 注意：math/rand 伪随机特性
// 提前调用 rand.Seed(time.Now().UTC().UnixNano()) 初始化随机种子，可避免预判生成特征。
func StrGenerateRandom(allowedChar []rune, n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = allowedChar[rand.Intn(len(allowedChar))]
	}
	return string(b)
}
