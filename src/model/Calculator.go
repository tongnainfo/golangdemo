package model

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"strings"
)

func Calculator(){
	inputReader := bufio.NewReader(os.Stdin)
	for i := 1; i < 100; i++ {
		fmt.Println("请输入有一个数字: ")
		var numone,numtwo  int
		var input, err = inputReader.ReadString('\n')
		if err == nil {
			i,error := strconv.Atoi(strings.Trim(input,"\n"))
			if error != nil{
				fmt.Println("字符串转换成整数失败")
				numone=0
			}else{
				numone=i
			}
		}
		fmt.Println("请输入有一个数字: ")
		var input1, err1 = inputReader.ReadString('\n')
		if err1 == nil {
			j,error := strconv.Atoi(strings.Trim(input1,"\n"))
			if error != nil{
				fmt.Println("字符串转换成整数失败")
				numtwo=0
			}
			numtwo=j
		}
		fmt.Println("我们将计算两个数字的和,他们的和是： ",numone+numtwo)
		fmt.Println("\n\n")
	}

}
