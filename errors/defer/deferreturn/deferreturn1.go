// return先于defer执行，defer对有名返回值或者指针返回值会产生影响。对无名返回值不会产生影响。
package main

func main() {
	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))
	println(*DeferFunc4(1))
}

//return 时 t=1； defer对t操作 t+=3; t为有名返回值，所以返回4
func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

//return 时 t=1 ； defer对t操作 t+=3; 无名返回值，所以虽然t=4,但是返回1
func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

//return 时 t=2 ； defer对t操作 t+=1; t为有名返回值，所以t=3，返回3
func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

//return 时 t=1 ； 函数为指针返回值，defer对t操作 t+=3，对返回结果产生影响;所以t=4，返回4
func DeferFunc4(i int) *int {
	t := i
	defer func() {
		t += 3
	}()
	return &t
}
