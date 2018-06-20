package helloworld

import "fmt"

const i100 = 100

//关键字iota，这个关键字用来声明enum的时候采用，它默认开始值是0，const中每增加一行加1
const (
	x = iota // x == 0
	y = iota // y == 1
	z = iota // z == 2
	w        // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
)

const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0

const (
	h, i, j = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
)

const (
	a       = iota //a=0
	b       = "B"
	c       = iota             //c=2
	d, e, f = iota, iota, iota //d=3,e=3,f=3
	g       = iota             //g = 4
)

func helloworld() {
	var c complex64 = 5 + 5i
	//output: (5+5i)
	fmt.Printf("Value is: %v\n", c)
	fmt.Printf("Hello, world or 你好，世界 or καλημ ́ρα κóσμ or こんにちはせかい\n")
}

func modifystring() {
	var ss string = "sfasfda"
	s := "hello"
	c := []byte(s) // 将字符串 s 转换为 []byte 类型
	c[0] = 'c'
	s2 := string(c) // 再转换回 string 类型
	fmt.Printf("%s\n%s\n", s2, ss)
}

func linkstring() {
	s := "hello,"
	m := " world"
	a := s + m
	fmt.Printf("%s\n", a)
}

func modifystring2() {
	s := "hello"
	s = "c" + s[1:] // 字符串虽不能更改，但可进行切片操作
	fmt.Printf("%s\n", s)
}

func multilinestring() {
	m := `hello
    world`
	fmt.Printf("%s\n", m)
	// err := errors.New("emit macho dwarf: elf header corrupted")
	// if err != nil {
	// 	fmt.Print(err)
	// }
}
