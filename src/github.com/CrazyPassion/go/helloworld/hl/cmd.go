package helloworld

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func PrintArgs() {
	var s, sep string
	for i, args := range os.Args[1:] {
		fmt.Printf("%d :", i)
		s += sep + args
		sep = " "
	}
	fmt.Println(s)
}

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func PrintArgs2() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
