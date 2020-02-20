package atest

import "fmt"

type X int

var XX []X

func init() {
	XX = append(XX, 1000)
}

func Regist(x X) {
	XX = append(XX, x)
	fmt.Println(">>>>>>>regist", x)
}
