package main

import (
	"fmt"
	"math"
)

// 声明一个新的类型
type person struct {
	name string
	age  int
}

// 比较两个人的年龄，返回年龄大的那个人，并且返回年龄差
// struct也是传值的
func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age { // 比较p1和p2这两个人的年龄
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func initstruct() {
	var tom person

	// 赋值初始化
	tom.name, tom.age = "Tom", 18

	// 两个字段都写清楚的初始化
	bob := person{age: 25, name: "Bob"}

	// 按照struct定义顺序初始化值
	paul := person{"Paul", 43}

	tb_Older, tb_diff := Older(tom, bob)
	tp_Older, tp_diff := Older(tom, paul)
	bp_Older, bp_diff := Older(bob, paul)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, bob.name, tb_Older.name, tb_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, paul.name, tp_Older.name, tp_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		bob.name, paul.name, bp_Older.name, bp_diff)
}

type Skills []string

type Human struct {
	name   string
	age    int
	////anoymousstruct
	//weight int
	phone string
}

type Student struct {
	Human      // 匿名字段，struct
	////anoymousstruct
	//Skills     // 匿名字段，自定义的类型string slice
	//int        // 内置类型作为匿名字段
	//speciality string
	school string
}

//func anoymousstruct() {
//	// 初始化学生Jane
//	jane := Student{Human: Human{"Jane", 35, 100}, speciality: "Biology"}
//	// 现在我们来访问相应的字段
//	fmt.Println("Her name is ", jane.name)
//	fmt.Println("Her age is ", jane.age)
//	fmt.Println("Her weight is ", jane.weight)
//	fmt.Println("Her speciality is ", jane.speciality)
//	// 我们来修改他的skill技能字段
//	jane.Skills = []string{"anatomy"}
//	fmt.Println("Her skills are ", jane.Skills)
//	fmt.Println("She acquired two new ones ")
//	jane.Skills = append(jane.Skills, "physics", "golang")
//	fmt.Println("Her skills now are ", jane.Skills)
//	// 修改匿名内置类型字段
//	jane.int = 3
//	fmt.Println("Her preferred number is", jane.int)
//}

type Rectangle struct {
	width,height float64
}

type Circle struct{
	radius float64
}

func (r Rectangle) area()(area float64)  {
	area = r.width * r.height
	return area
}

func (r Circle) area()(float64)  {
	return r.radius * r.radius * math.Pi
}

func methodf()  {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())
}

const(
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color Color
}

type BoxList []Box //a slice of boxes

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color) {
	b.color = c
}

func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PaintItBlack() {
	for i, _ := range bl {
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string {"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

//自定义结构的method继承f
func methodf2() {
	boxes := BoxList {
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
	fmt.Println("The color of the last one is",boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.BiggestColor().String())

	fmt.Println("Let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is", boxes[1].color.String())

	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())
}

type Employee struct {
	Human //匿名字段
	company string
}

//在human上面定义了一个method
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//method inherit
func methodf3() {
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

	mark.SayHi()
	sam.SayHi()
}

func main() {
	methodf3()
}
