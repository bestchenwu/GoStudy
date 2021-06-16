package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//var s string = ""
	//for i :=1 ;i<len(os.Args);i++{
	//	s+=" "+os.Args[i]
	//}
	//fmt.Println("s="+s)
	s := ""
	//t := ""
	for _,arg := range os.Args[1:]{
		s+=""+arg
	}
	//fmt.Println("s="+s)
	//fmt.Println(strings.Join(os.Args[1:], " "))
	//练习1.1
	//for i:=1;i<len(os.Args);i++{
	//	fmt.Println(i)
	//	fmt.Println("value="+os.Args[i])
	//}
	//1.3 查找重复的行
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		counts[input.Text()]++
	}
	for line,n := range counts{
		if n>1{
			fmt.Println("%d\t%s\n",n,line)
		}
	}
	//fmt.Println("hello,world")
}
