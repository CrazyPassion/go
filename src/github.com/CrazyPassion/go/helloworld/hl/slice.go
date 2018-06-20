package helloworld

import "fmt"

func SliceandArrary() {
	//arrary
	months := [...]string{1: "January", 2: "Feb", 3: "March", 4: "Apr", 5: "May", 6: "June", 7: "July", 8: "Aug", 9: "Sep", 10: "Oct", 11: "Nov", 12: "December"}
	summer := months[4:7]
	fmt.Println(summer)
	endlessSummer := summer[:5] // extend a slice (within capacity)
	fmt.Println(endlessSummer)  // "[June July August September October]"
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Reverse() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)
	s := []int{0, 1, 2, 3, 4, 5}
	reverse(s)
	fmt.Println(s)
	reverse(s[:2])
	fmt.Println(s)
}

func Append() {
	var runes []rune
	var test = []rune("test rune")
	for _, v := range "Hello world!" {
		runes = append(runes, v)
	}
	fmt.Printf("%q\n", runes)
	fmt.Printf("%q\n", test)
}
