package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("现在的时间是：", time.Now())
	fmt.Println("现在的时间是：", time.Now().Format("2006-01-02 15:04:05"))
	time1 := time.Now()
	fmt.Println("现在的时间是：", time1.Add(time.Second*3600).Format("2006-01-02 15:04:05"))

	//获得一个unixDate
	t, err := time.Parse(time.UnixDate, "Sat Mar 7 11:06:39 PST 2015")
	if err != nil {
		panic(err)
	}

	//不进行任何格式化打印
	fmt.Println("default format:", t)

	//unixDate格式
	fmt.Println("Unix format", t.Format(time.UnixDate))

	// 给时间附加时区.
	fmt.Println("Same, in UTC:", t.UTC().Format(time.UnixDate))

	//定义一个时间显示的函数
	do := func(name, layout, want string) {
		got := t.Format(layout)
		if want != got {
			fmt.Printf("error:for %q got %q; expected %q\n", layout, got, want)
			return
		}
		fmt.Printf("%-15s %q gives %q\n", name, layout, got)
	}

	fmt.Printf("\n时间格式化：\n\n")

	//格式化字符串在形式上表现为时间截，是固定的
	//	Jan 2 15:04:05 2006 MST
	//上面的时间截按顺序表现的值为，在格式化字符串中可以直接引用
	//	1   2 3   4  5    6 -7

	//一个简单的例子
	do("Basic", "Mon Jan 2 15:04:05 MST 2006", "Sat Mar 7 11:06:39 PST 2015")

	//对于固定宽度值的输出，如时间（7或者07）只有一个或者两个字符，在输出格式字符串中使用_代替空格
	//这里我们只打印天，天在我们的格式字符串中是2，在值中是7
	do("No Pad", "<2>", "<7>")

	// 在需要时可以加上下划线，表示0
	do("Spaces", "<_2>", "< 7>")

	// 相似的0表示0
	do("Zeros", "<02>", "<07>")

	//如果一个值已经是正确的宽度了，就不需要加0或者_，如取秒时，为39秒，则不需要加，
	//但是分和秒合起来取时需要加。
	do("Suppressed pad", "04:05", "06:39")

	// Unix格式使用了一个下划线去补齐天数.
	// 和上面的简单例子比较一下吧.
	do("Unix", time.UnixDate, "Sat Mar  7 11:06:39 PST 2015")

	//小时可以通过3PM,3pm和15h取。
	do("AM/PM", "3PM==3pm==15h", "11AM==11am==11h")

	//在时间后面添加分秒
	t, err = time.Parse(time.UnixDate, "Sat Mar  7 11:06:39.1234 PST 2015")
	if err != nil {
		panic(err)
	}
	//要打印出分秒，需要在分后面加一个'0的字符串'或者'9的字符串'，如果是0的字符串，
	//那么分秒是指定的字符串，并后缀0

	do("0s for fraction", "15:04:05.000000", "11:06:39.123400")

	//如果是9的字符串，那么就没有后缀的0。
	do("9s for fraction", "15:04:05.99999999", "11:06:39.1234")
}
