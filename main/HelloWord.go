package main

import (
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
	fmt.Println("s="+s)
	//fmt.Println("hello,world")
}
