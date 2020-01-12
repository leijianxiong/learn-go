package main

import "fmt"

func main()  {
	s := []interface{}{"a", "b", "c"}

	s2 := arrayKeys(&s)
	fmt.Println(s2, (*s2)[0])
	s3 := arrayKeys(s2)
}

func f1(s []string) {
	s[0] = "a2"
}

func arrayKeys(s *[]interface{}) *[]interface{}  {
	var ks []interface{}
	for k, _ := range *s {
		ks = append(ks, k)
	}
	return &ks
}

func arrayValues()  {

}
