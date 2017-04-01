package main

import "fmt"

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human      // 匿名字段，那么默认Student就包含了Human的所有字段
	Skills     //不仅仅是struct字段哦，所有的内置类型和自定义类型都是可以作为匿名字段的
	int        // 内置类型作为匿名字段
	speciality string
}

func main() {
	// 我们初始化一个学生
	mark := Student{Human: Human{"Mark", 25, 120}, speciality: "Computer Science"}

	// 我们访问相应的字段
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)

	// 修改他的skill技能字段
	mark.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", mark.Skills)
	fmt.Println("She acquired two new ones ")
	mark.Skills = append(mark.Skills, "physics", "golang")
	fmt.Println("Her skills now are ", mark.Skills)
	// 修改匿名内置类型字段
	mark.int = 3
	fmt.Println("Her preferred number is", mark.int)
}
