/*
	map的读取和设置也类似slice一样，通过key来操作，只是map的key多了很多类型，可以是int，可以是string及所有完全定义了==与!=操作的类型。
	1，map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取
	2，map的长度是不固定的，也就是和slice一样，也是一种引用类型
	3，内置的len函数同样适用于map，返回map拥有的key的数量
	4，map的值可以很方便的修改，通过numbers["one"]=11可以很容易的把key为one的字典值改为11
	5，map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	var numbers map[string]int
	numbers = make(map[string]int)
	numbers["one"] = 1  //赋值
	numbers["ten"] = 10 //赋值
	numbers["three"] = 3

	fmt.Println("第三个数字是: ", numbers["three"]) // 读取数据
	// 打印出来如:第三个数字是: 3

	//删除map的元素
	// 初始化一个字典
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	delete(rating, "C") // 删除key为C的元素

	//map的迭代操作
	for _, v := range numbers {
		fmt.Println(v) //这里打印的map是无序的，每次打印的结果可能不一样，如果要得到有序的结果，则需要对map进行排序。
	}

	//map的排序，对map进行排序，需要创建一个新的slice，将map的可以放到slice中，然后对slice进行排序，再迭代取出map的值。
	maplen := len(numbers)
	s := make([]string, maplen)
	i := 0
	for k, _ := range numbers {
		s[i] = k
		i++
	}
	sort.Strings(s)
	for i = 0; i < maplen; i++ {
		fmt.Println(numbers[s[i]])
	}
}
