//数组的排列算法。 也是对二维数组的应用。
package main

import (
	"errors"
	"fmt"
)

func main() {
	s := []int{1, 0, -1, 2, -2}
	if solution, err := permutate(s); err == nil {
		fmt.Println(solution)
		fmt.Println(len(solution))
	}
}

func fourSum(nums []int, target int) ([][]int, error) {
	return nil, nil
}

// 递归完成一个数组的排列。
// 先从数组中选出一个，再递归地将剩下的数组排列，再合并。
func permutate(nums []int) ([][]int, error) {
	n := len(nums)
	lennum := len(nums)
	if n < 1 || lennum < 1 || lennum < n {
		return nil, errors.New("传入参数错误")
	}
	var solution = [][]int{}
	var leavesolution [][]int
	var err error
	if n == 1 {
		solution = append(solution, nums)
	} else {
		for i := 0; i < lennum; i++ {
			leavenums := delOneFromArray(nums, nums[i])
			if leavesolution, err = permutate(leavenums); err != nil {
				return nil, err
			}
			for _, v := range leavesolution {
				solution = append(solution, append(v, nums[i]))
			}
		}
	}
	return solution, nil
}

// 从slice中删除第n个元素，返回。
func delOneFromArray(slice []int, n int) []int {
	length := len(slice)
	result := []int{}
	for i := 0; i < length; i++ {
		if n != slice[i] {
			result = append(result, slice[i])
		}
	}
	return result
}
