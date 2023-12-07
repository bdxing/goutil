package goutil

import "math/rand"

// StrGenerateRandom 从种子字符串里随机生成 n 个字符组成字符串。
// 注意：math/rand 伪随机特性
// 提前调用 rand.Seed(time.Now().UTC().UnixNano()) 初始化随机种子
func StrGenerateRandom(n int, allowedChar []rune) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = allowedChar[rand.Intn(len(allowedChar))]
	}
	return string(b)
}
