package atest

import "fmt"

type X int

var XX []X

func init() {
	fmt.Println(">>>>>>>", &XX)
	XX = append(XX, 1000)
}

func Regist(x X) {
	fmt.Printf("<<<<<<<<<<<", &XX)
	XX = append(XX, x)
	fmt.Println(">>>>>>>regist", x)
}
