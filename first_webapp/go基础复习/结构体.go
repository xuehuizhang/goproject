package main

import (
	"first_webapp/model"
	"fmt"
)
//结构体是值类型
type Cat struct{
	Name string
	Age int
	Color string
	Hobby string
}

type Person struct {
	Name string
	Age int
	Scores [5]float64
	ptr *int
	slice []int
	map1 map[int]string
}

//结构体是自定义数据类型，代表一类事物
//结构体变量是具体的，实际的，代表一个具体变量
func main()  {
	var c Cat=Cat{Name:"小花",Age:19,Color:"白色",Hobby:"吃鱼"}
	fmt.Println(c)

	var p Person
	p.map1=make(map[int]string)
	p.map1[1]="李四"
	fmt.Println(p)

	/*ai:=NewAnimal("猫",19)
	fmt.Println(ai.Name)*/

	a1:=model.NewCat("里斯",18)
	fmt.Println(a1.Name)
}
