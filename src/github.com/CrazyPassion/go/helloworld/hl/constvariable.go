package helloworld

import (
	"fmt"
)

const (
	_ = 1 << (iota * 10)
	KB
	MB
	GB
	TB
)

func Const() {
	fmt.Println(GB / KB)
	fmt.Println(GB / MB)
}
