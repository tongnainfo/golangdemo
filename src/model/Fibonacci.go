package model

import "fmt"

func fibonacci() func() int {
	last1 := 0
	last2 := 0
	return func() int {
		sum:= last1 + last2
		last1 = last2;
		last2 = sum
		if last2 == 0 {
			last2 = 1
		}
		return last1
	}
}
func Fibonacci() {
	f,x:= fibonacci(),fibonacci()
	for i := 0; i < 50; i++ {
		fmt.Println(f())
	}
	for i := 0; i < 10; i++ {
		fmt.Println(x())
	}

}
