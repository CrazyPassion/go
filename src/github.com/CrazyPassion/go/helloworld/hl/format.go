package helloworld

import "fmt"

//Printf print
func Printf() {
	fmt.Printf("%[2]d %[1]d\n", 11, 22)
	fmt.Printf("%[3]*.[2]*[1]f\n", 12.0, 2, 6)
	fmt.Printf("%d %d %#[2]x %#[1]x\n", 16, 17)
}
