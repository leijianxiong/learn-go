package main


func main()  {
	f1 := func(int) {

	}

	f2 := func() {}

	i := 1
	//i.(F1)
	F1(i)
	_ = F1(f1)
	_ = F1(f2)
}






type F1 func(i int)