/*
	make用于内建类型（map、slice 和channel）的内存分配。new用于各种类型的内存分配。
	内建函数new本质上说跟其它语言中的同名函数功能一样：new(T)分配了零值填充的T类型的内存空间，并且返回其地址，
	即一个*T类型的值。用Go的术语说，它返回了一个指针，指向新分配的类型T的零值。有一点非常重要：

	new返回指针。

	内建函数make(T, args)与new(T)有着不同的功能，make只能创建slice、map和channel，并且返回一个有初始值(非零)的T类型，
	而不是*T。本质来讲，导致这三个类型有所不同的原因是指向数据结构的引用在使用前必须被初始化。
	例如，一个slice，是一个包含指向数据（内部array）的指针、长度和容量的三项描述符；
	在这些项目被初始化之前，slice为nil。对于slice、map和channel来说，make初始化了内部的数据结构，填充适当的值。

	make返回初始化后的（非零）值。

	总的来说，make需要传入参数才可以分配内存，返回的是对象，该对象一般包含指向数据的指针，及其他描述信息。如slice包含一个
	指向数组的指针，长度和容量等描述符。
	而new则不需要传入参数，系统默认分配内存并用零值填充，返回指向数据的指针。
*/

package main

import (
	"fmt"
)

func main() {

}
