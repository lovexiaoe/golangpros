// 数组的组合算法，也是对二维数组的应用。
package main

import (
	"errors"
	"fmt"
)

// pernutate1 和permutate2方法思想类似。permutate2简化一些，但是最后需要转换成实际的组合。

func main() {
	s := []int{1, 0, -1, 2, -2}
	if solution1, err := permutate1(s, 4); err == nil {
		fmt.Println(solution1)
		fmt.Println(len(solution1))
	}
	if solution2, err := permutate2(s, 4); err == nil {
		fmt.Println(solution2)
		fmt.Println(len(solution2))
	}
}

// 递归完成一个数组中n个元素的组合。根据序列，从小到大生产组合。
// 如C(5,3)。可以用5进制的排序过的3位数表示，各位数字不重复。最小的数字为012，最大的数字为234。
// 从012开始每次加1，当每位达到最大时进位。如124进位后为134，而不是132。
func permutate1(nums []int, n int) ([][]int, error) {
	lennum := len(nums)
	if n < 1 || lennum < 1 || lennum < n {
		return nil, errors.New("传入参数错误")
	}
	var solution = [][]int{}
	topsolution := make([]int, n)
	sequeSlice1(topsolution, 0)
	solution = append(solution, topsolution)
	for k := 0; ; {
		secondsolution := make([]int, n)
		if solution[k][0] >= lennum-n {
			break
		}
		k++
		copy(secondsolution, solution[k-1])
		for j := 0; j < n; j++ {
			if secondsolution[n-j-1] < lennum-j-1 {
				secondsolution[n-j-1]++
				break
			} else if secondsolution[n-j-1] == lennum-j-1 {
				//这里是进位的逻辑，当某位达到最大时，如果前一位没达到最大，则给前一位加1，后面的位顺序设置为前一位加1.
				if n-j-1 > 0 && secondsolution[n-j-2] < lennum-j-2 {
					secondsolution[n-j-2]++
					sequeSlice1(secondsolution, n-j-2)
					break
				}
			}
		}
		//fmt.Println(solution)
		solution = append(solution, secondsolution)
	}
	return solution, nil
}

// 当i>n时，将array[i+1]的元素设置为array[i]+1, 如传入([1,4,3],1),返回[1,4,5]
func sequeSlice1(slice []int, n int) []int {
	if n < 0 {
		n = 0
	}
	if n >= len(slice) {
		return slice
	}
	for i := n + 1; i < len(slice); i++ {
		slice[i] = slice[i-1] + 1
	}
	return slice
}

// 递归完成一个数组中n个元素的组合。
// 一个集合中的元素可以用1表示选取，0表示未选。
// 如C(5,3),包括组合[11100],[00111]
// 1，从右向左找第一个10。进行交换。
// 2，将上次找到10位置右边的1全部移动到左边。
// 3，重复1,2。
func permutate2(nums []int, n int) ([][]int, error) {
	lennum := len(nums)
	if n < 1 || lennum < 1 || lennum < n {
		return nil, errors.New("传入参数错误")
	}
	var solution = [][]int{}
	topsolution := make([]int, lennum)
	for i := 0; i < lennum; i++ {
		if i < n {
			topsolution[i] = 1
		} else {
			topsolution[i] = 0
		}
	}
	solution = append(solution, topsolution)

	for k := 0; ; {
		//fmt.Println(solution)
		secondsolution := make([]int, lennum)
		copy(secondsolution, solution[k])
		var j int
		for j = lennum - 1; j > 0; j-- {
			if secondsolution[j] == 0 && secondsolution[j-1] == 1 {
				swap(secondsolution, j, j-1)
				secondsolution = sequeSlice2(secondsolution, j)
				solution = append(solution, secondsolution)
				k++
				break
			}
		}
		if j <= 0 {
			break
		}

	}
	return solution, nil
}

// 将i>n的元素排序。
func sequeSlice2(slice []int, n int) []int {
	slen := len(slice)
	if n < 0 {
		n = 0
	}
	if n >= slen {
		return slice
	}
	for i := n + 1; i < slen; i++ {
		for j := i + 1; j < slen; j++ {
			if slice[i] < slice[j] {
				swap(slice, i, j)
			}
		}
	}
	return slice
}

// slice 元素交换
func swap(slice []int, m int, n int) {
	swap := slice[m]
	slice[m] = slice[n]
	slice[n] = swap
}
