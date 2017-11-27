package main

import (
	"fmt"
)

func main() {
	control()
}

func multiArray() {
	var arr [10]int                                 // 声明了一个int类型的数组
	arr[0] = 42                                     // 数组下标是从0开始的
	arr[1] = 13                                     // 赋值操作
	fmt.Printf("The first element is %d\n", arr[0]) // 获取数据，返回42
	fmt.Printf("The last element is %d\n", arr[9])  //返回未赋值的最后一个元素，默认返回0

	a := [3]int{1, 2, 3} // 声明了一个长度为3的int数组

	b := [10]int{1, 2, 3} // 声明了一个长度为10的int数组，其中前三个元素初始化为1、2、3，其它默认为0

	c := [...]int{4, 5, 6} // 可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度
	fmt.Printf("The first element of a, b,c,%d %d %d\n", a[0], b[0], c[0])

	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	easyArray := [2][4]int{{1, 2, 3, 5}, {5, 6, 7, 8}}
	fmt.Printf("The first element of multi-arry,%d %d\n", doubleArray[0], easyArray[0])
}

func slicefun() {
	// 声明一个数组
	var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// 声明两个slice
	var aSlice, bSlice []byte

	// 演示一些简便操作
	aSlice = array[:3] // 等价于aSlice = array[0:3] aSlice包含元素: a,b,c
	aSlice = array[5:] // 等价于aSlice = array[5:10] aSlice包含元素: f,g,h,i,j
	aSlice = array[:]  // 等价于aSlice = array[0:10] 这样aSlice包含了全部的元素

	// 从slice中获取slice
	// slice是引用类型，所以当引用改变其中元素的值时，其它的所有引用都会改变该值，例如上面的aSlice和bSlice，如果修改了aSlice中元素的值，那么bSlice相对应的值也会改变。
	aSlice = array[3:7]  // aSlice包含元素: d,e,f,g，len=4，cap=7
	bSlice = aSlice[1:3] // bSlice 包含aSlice[1], aSlice[2] 也就是含有: e,f
	bSlice = aSlice[:3]  // bSlice 包含 aSlice[0], aSlice[1], aSlice[2] 也就是含有: d,e,f
	bSlice = aSlice[0:5] // 对slice的slice可以在cap范围内扩展，此时bSlice包含：d,e,f,g,h
	bSlice = aSlice[:]   // bSlice包含所有aSlice的元素: d,e,f,g
	aSlice[0] = 0
	fmt.Printf("The first element of aSlice:%d bSlice:%d，array:%d\n", aSlice[0], bSlice[0], array)
	bSlice[0] = 'A'
	fmt.Printf("The first element of aSlice:%d bSlice:%d，array:%d\n", aSlice[0], bSlice[0], array)
}

func mapfunc() {
	// 声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
	// var numbers map[string]int
	// 另一种map的声明方式
	numbers := make(map[string]int)
	numbers["one"] = 1  //赋值
	numbers["ten"] = 10 //赋值
	numbers["three"] = 3

	fmt.Println("第三个数字是: ", numbers["three"]) // 读取数据
	// 打印出来如:第三个数字是: 3

	// delete
	// 初始化一个字典
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}
	delete(rating, "C") // 删除key为C的元素

	var testmap = rating["C"]
	fmt.Println("testmap:", testmap)

	//map也是引用
	m := make(map[string]string)
	m["Hello"] = "Bonjour"
	m1 := m
	m1["Hello"] = "Salut" // 现在m["hello"]的值已经是Salut了
	fmt.Println("m1[\"Hello\"]:", m["Hello"])
}

func control() {
	if x := 5; x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}

	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("sum is equal to ", sum)

	sum1 := 1
	for sum1 < 1000 {
		sum1 += sum1
	}
	fmt.Println("sum is equal to ", sum1)

	m := map[string]int{"cc":5}
	for k,v:=range m {
		fmt.Println("map's key:",k)
		fmt.Println("map's val:",v)
	}
	for _, v := range m{
		fmt.Println("map's val:", v)
	}
}
