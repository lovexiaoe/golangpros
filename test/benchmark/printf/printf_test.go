// 基准测试用于测试代码的性能，当你需要测试一个问题的不同解决方案时，十分有用。
//很多开发者用于测试不同的并发方案。benchmark测试有很多选项，借助go tools运行以应用这些选项。

//本例测试将一个数字转换成字符串的三种方法的效率。
package listing05_test

import (
	"fmt"
	"strconv"
	"testing"
)

// BenchmarkSprintf provides performance numbers for the
// fmt.Sprintf function.
//基准测试的名称以Benchmark开始，参数为testing.B的指针。并且使用循环来达到测试效果。
func BenchmarkSprintf(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", number)
	}
}

// BenchmarkFormat provides performance numbers for the
// strconv.FormatInt function.
func BenchmarkFormat(b *testing.B) {
	number := int64(10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(number, 10)
	}
}

// BenchmarkItoa provides performance numbers for the
// strconv.Itoa function.
func BenchmarkItoa(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}
}
