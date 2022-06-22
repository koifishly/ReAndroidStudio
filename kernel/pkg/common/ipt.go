package common

import "fmt"

func ScanString(info string) (str string) {
	fmt.Print(info)
	fmt.Scanln(&str)
	return
}

func ScanInt(info string) (n int) {
	fmt.Print(info)
	fmt.Scanln(&n)
	return
}
