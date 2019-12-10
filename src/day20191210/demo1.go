package main

import "fmt"

/*
思考题
如果我们把一个值为nil的某个实现类型的变量赋给了接口变量，那么在这个接口变量上仍然可以调用该接口的方法吗？如果可以，有哪些注意事项？如果不可以，原因是什么？
https://time.geekbang.org/column/article/18037
 */

type Pet interface {
	Name() string
	setName(name string)
	m1()
	m2()
}

type Cat struct {
	name string
}

func (cat Cat) Name() string  {
	return cat.name
}

func (cat *Cat) setName(name string)  {
	cat.name = name
}

func (cat Cat) m0()  {
	fmt.Println("m0 method called")
}

func (cat *Cat) m1()  {
	fmt.Println("m1 method called")
}

func (cat *Cat) m2() {
	fmt.Println(1, cat.name)
	fmt.Println(2, cat.Name())
}

func NewCat(name string) *Cat {
	return &Cat{
		name: name,
	}
}

func main()  {

	cat := NewCat("cat1")
	cat = nil
	var pet = cat
	pet.m0() //1 => 空指针不能调用值方法 只能调用指针方法(参考pet.m1()), 方法里不能调用动态类型的属性(参考pet.m2()
	pet.m1() //2
	pet.m2() //3
}