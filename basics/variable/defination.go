package variable

func main() {
	//定义三个类型都是“type”的变量
	var vname1, vname2, vname3 int
	/*
		定义三个类型都是"type"的变量,并且分别初始化为相应的值
		vname1为1，vname2为2，vname3为3
	*/
	var vname1, vname2, vname3 int = 1, 2, 3
	//type可以省略,go会自动推倒出类型。
	var vname1, vname2, vname3 = 1, 2, 3

	/*
		定义三个变量，它们分别初始化为相应的值
		vname1为1，vname2为2，vname3为3
		编译器会根据初始化的值自动推导出相应的类型，不过这种声明方法只能在函数内部使用。
		函数外部使用则会无法编译通过，所以一般用var方式来定义全局变量。
	*/
	vname1, vname2, vname3 := 1, 2, 3
}
