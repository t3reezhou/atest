package atest

type X int

var XX []X

func init() {
	XX = append(XX, 1000)
}

func Regist(x X) {
	XX = append(XX, x)
}
