package demo

import "fmt"

func RunWithAddLog(f func() int) {
	fmt.Println("start")
	n := f()
	fmt.Println("done")
	fmt.Println("result", n)
}
