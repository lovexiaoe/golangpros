package main

type student struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	// 错误写法，新手容易犯的错误。
	for _, stu := range stus {
		m[stu.Name] = &stu //这里引用stu的地址，然而程序运行时只会引用stus最后一个元素的值。
	}

	for k, v := range m {
		println(k, "=>", v.Name)
	}

	// 正确
	//	for i := 0; i < len(stus); i++ {
	//		m[stus[i].Name] = &stus[i]
	//	}
	for _, stu := range stus {
		stu1 := stu //将stu保存为局部变量stu1。避免上面的错误。
		m[stu.Name] = &stu1
	}
	for k, v := range m {
		println(k, "=>", v.Name)
	}
}
