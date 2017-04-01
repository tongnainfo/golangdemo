package model

import "fmt"

func Demo()  {
	x:=make([]int,0)
	x=append(x,1)
	//x=x[0:6]
	fmt.Println(x,len(x),cap(x))
	for i := 1; i < 100; i++ {
		if i%2 == 0 {
			x=append(x,i)
		}
	}
	fmt.Println(x,len(x),cap(x))

}
