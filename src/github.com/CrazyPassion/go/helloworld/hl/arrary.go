package helloworld

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func Initial() {
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	for _, v := range q {
		fmt.Println(v)
	}
	for _, v := range r {
		fmt.Println(v)
	}

	// 数组的长度是数组类型的一个组成部分，因此[3]int和[4]int是两种不同的数组类型
	s := [...]int{1, 2, 3, 4}
	fmt.Printf("%T, len:%d\n", s, len(s)) // "[4]int"

	// 最后一个元素被初始化为-1，其它元素都是用0初始化
	nr := [...]int{99: -1}
	for i, v := range nr {
		if i > 95 {
			fmt.Println(v)
		}
	}
}

const (
	USD = iota
	EUR
	GBP
	RMB
)

var symbol = [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}

func Money() {
	fmt.Println(RMB, symbol[RMB])
}

const (
	SHA256 = iota
	SHA384
	SHA512
)

func Sha(shatype int) {
	fmt.Println(SHA256, SHA384, SHA512)
	var sha1, sha2 [32]uint8
	switch shatype {
	case SHA256:
		sha1 = sha256.Sum256([]byte{'X'})
		sha2 = sha256.Sum256([]byte("X"))
	case SHA384:
		sha1 = sha256.Sum256([]byte{'X'})
		sha2 = sha256.Sum256([]byte("X"))
	case SHA512:
		sha1 = sha512.Sum512_256([]byte{'X'})
		sha2 = sha512.Sum512_256([]byte("Z"))
	}
	fmt.Printf("%x\n%x\n%t\n%T\n", sha1, sha2, sha1 == sha2, sha1)
	fmt.Println(shadiffPopcount(sha1, sha2))
	zero(&sha1)
	zero(&sha2)
	fmt.Printf("%x\n%x\n%t\n%T\n", sha1, sha2, sha1 == sha2, sha1)
	fmt.Println(shadiffPopcount(sha1, sha2))
}
func zero(ptr *[32]byte) {
	*ptr = [32]byte{}
}
func shadiffPopcount(sha1, sha2 [32]uint8) int {
	count := 0
	for i, v := range sha1 {
		shadiff := v ^ sha2[i]
		// fmt.Printf("diff:%x,count:%d\n", shadiff, popcountByte(shadiff))
		count += popcountByte(shadiff)
	}
	return count
}

func popcountByte(x uint8) int {
	count := 0
	for x != 0 {
		count++
		x &= (x - 1)
	}
	return count
}
