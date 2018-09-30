package methodtest

import (
	"testing"
)

func TestDelOneFromArray(t *testing.T) {
	//下面是一张不同参数和结果的表。
	var urls = []struct {
		array  []int
		n      int
		result []int
	}{
		{
			[]int{1, 2, 3, 5, 6},
			3,
			[]int{1, 2, 5, 6},
		},
		{
			[]int{9, 2, 3, 5, 6},
			6,
			[]int{9, 2, 5, 5},
		},
	}

	t.Log("\t测试方法delOneFromArray开始.")
	{
		for _, u := range urls {
			t.Log("\t从 ", u.array, " 中删除 ", u.n, " 应该得到: ", u.result)

			result := delOneFromArray(u.array, u.n)

			t.Log("\t实际得到: ", result)

			if compareSlice(result, u.result) {
				t.Log("\t删除成功。 ")
			} else {
				//t.Error说明测试出错，单并不会退出测试。
				t.Error("\t删除失败。 ")
			}

		}
	}
}

func compareSlice(slice1 []int, slice2 []int) bool {
	if slice1 == nil || slice2 == nil {
		return false
	}
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
