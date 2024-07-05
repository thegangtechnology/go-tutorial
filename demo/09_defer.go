package demo

import (
	"bufio"
	"fmt"
	"os"
)

func Cat(fname string) {
	// golang read file line by line
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close() // there are some subtlety
	//defer func() {
	//	err := f.Close()
	//	if err != nil {
	//		fmt.Println("close error", err)
	//	}
	//}()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func DeferSubtlety() {
	x := 1
	//this is why I do anonymous function call
	defer fmt.Println(x)
	x++
	fmt.Println("hello", x)
}
