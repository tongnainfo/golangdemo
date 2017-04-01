package main

import (
	"fmt"
	"os"
	"math/rand"
	"bytes"
	"strconv"
	_"model"
	"bufio"
	"time"
)
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(20 * time.Millisecond)
		fmt.Println(s)
	}
}
func showMessage(i int)  {
	fmt.Println("num is:",i)
}
type Man struct {
	Age int
}
func main() {
	defer showMessage(100)
	for i:=1;i<10;i++{
		defer showMessage(i)
	}
}
func xx(p * int)  {
	*p=5;
}
func makeBig(p  int)  {
	p=20;
}
func x() {
	var x []int
	var array [10]int
	for i := 0; i < 10; i++ {
		array[i]=i
	}
	//[0,4)
	x = array[0:4]
	fmt.Print(x)
}
func demo1() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	var input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
}
func Demo() {
	fmt.Println("Hello, 世界")
	fmt.Println(os.Args)
	sayHi("i am ada")
	fmt.Println("rand", rand.Intn(10))
	a, b := max(1, 56)
	fmt.Println(a, b)
}



func sayHi(msg string) {
	fmt.Println(msg)
}
func max(one, two int) (int, int) {
	if one > two {
		return one, 1
	} else {
		return two, 2
	}
}
func shengfa() {
	var buffer = bytes.Buffer{}
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			buffer.WriteString(strconv.Itoa(i));
			buffer.WriteString("*")
			buffer.WriteString(strconv.Itoa(j))
			buffer.WriteString("=")
			buffer.WriteString(strconv.Itoa(i * j))
			buffer.WriteString("\t")
		}
		buffer.WriteString("\n")
	}
	fmt.Println(buffer.String())
}
