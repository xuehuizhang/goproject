package model

type animal struct {
	Name string
	age int
}

type cat struct {
	animal
}

func NewCat(name string,age int) *cat  {
	return &cat{
		animal{Name:name,age:age,},
	}
}
