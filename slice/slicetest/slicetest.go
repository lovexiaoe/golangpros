/*
	在go中一般很少直接使用array。slice并不是真正意义上的动态数组，slice只是对数组的一部分的封装和展示，是一种引用类型，其总是指向一个底层array。
*/

package main

import (
	"fmt"
)

func main() {
	//声明一个slice和array类似，只是少了长度。
	//var fslice []int
	//初始化一个slice和array也类似。
	//slice := []byte{'a', 'b', 'c', 'd'}

	//slice可以从一个数组或一个已经存在的slice中再次声明。slice通过array[i:j]来获取，其中i是数组的开始位置，j是结束位置，但不包含array[j]，它的长度是j-i。
	// 声明一个含有10个元素元素类型为byte的数组
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// 声明两个含有byte的slice
	var a, b, c, d []byte
	// a指向数组的第3个元素开始，并到第五个元素结束，
	a = ar[2:5]
	//现在a含有的元素: ar[2]、ar[3]和ar[4]
	// b是数组ar的另一个slice
	b = ar[3:5]
	// b的元素是：ar[3]和ar[4]
	c = ar[:3] //等价于 ar[0:3] c包含元素: a,b,c
	d = ar[5:] // 等价于ar[5:10] ar包含元素: f,g,h,i,j
	//e = ar[:]  // 等价于ar[0:10] 这样ar包含了全部的元素

	//slice使用make初始化
	s1 := make([]int, 3, 10) //数组长度为3，cap容量为10，当slice分配的元素超过cap时，程序会重新分配内存（消耗一定的资源）。
	fmt.Println(s1)

	//slice有几个有用的内置函数。
	fmt.Println(len(a)) //获取slice的长度
	fmt.Println(cap(a)) //获取slice的最大容量。参考内建函数中的cap，数组切片的初始cap为数组的cap减去起始下标。
	fmt.Println(cap(b))
	fmt.Println(cap(c))
	//向slice追加一个或者多个元素，会改变slice所引用的数组。但是当slice中没有剩余空间（即（cap-len）==0）时，
	//此时将动态分配新的数组空间（大小由系统自动决定）。返回的slice数组指针将指向这个空间，而原数组的内容将保持不变；
	fmt.Printf("%p\n", c)
	c = append(c, 'm')
	//追加后，c为a,b,c,m。追加没有超过c的cap，所以不会重新分配空间。
	//ar为a,b,c,m,e,f,g,h,i,j。ar变化了，但是其cap不会变化。
	fmt.Println(cap(c))
	fmt.Printf("%v,%p\n", string(c), c) //c的地址没有改变

	//原始数组ar保持不变，cap也不会变化。
	fmt.Printf("%v,%p\n", string(d), d)
	d = append(d, 'l')
	fmt.Println(cap(d))
	//追加后，超过的ar的cap，系统分配新的空间，将追加后的d复制到新的空间中，d的指针指向新的空间。
	fmt.Printf("%v,%p\n", string(d), d) //d添加l后，超出了数组的cap，系统复制并重新分配内存，地址发生了改变，

	//数组copy
	s2 := []int{1, 2, 3, 4, 5, 6}
	s3 := []int{8, 9, 10}
	copy(s2[1:3], s3[1:2]) //第一个参数为目标数组，第二个为原数组。也可以不加截取，整个slice拷贝，
	fmt.Println(s2)        //[1,9,3,4,5,6]

	//slice的迭代
	for i, v := range s3 {
		fmt.Println(i, "=", v)
	}
}
