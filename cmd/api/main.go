package main

import (
	"fmt"
	"os/user"
	"time"
)

type Point struct {
	x int
	y int
}


type Vertex struct {
	X,Y int
}

func (v Vertex) Area() int {
	return v.X * v.Y
}

func add(x int, y int) int{
	return x + y
}

//複数の値の返り値
func calc(x,y int) (plus int, minus int){
	plus = x + y
	minus = x - y
	return plus, minus
}

//可変長引数
func numbers(params ...int){
	fmt.Println(params)
}

func main() {
	fmt.Println("Hello World!")
	fmt.Println(time.Now())
	fmt.Println(user.Current())
	var i int = 1
	var f float64 = 1.2
	var s string = "test1"
	var t bool = true
	fmt.Println(i,f,s,t)

	// 配列
	const n = 2
	var a [n]int = [n]int{100,200}
	fmt.Println(a)
	var b []int = []int{1,2}
	b = append(b,3)
	fmt.Println(b)
	fmt.Println(b[0:2])

	//make cap
	x := make([]int,3,5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(x),cap(x),x)

	// c := make([]int,5)
	c := make([]int,0,5)
	for i:= 0; i<5;i++{
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)

	// map
	m := map[string]int{"apple": 100}
	fmt.Println(m)
	m["banana"] = 200
	fmt.Println(m)
	v, err := m["vsvs"]
	fmt.Println(v,err)

	//byte
	bt := []byte{72,73}
	fmt.Println(bt)
	fmt.Println(string(bt))


	//関数
	fmt.Println(add(11,2))
	plus,minus := calc(11,1)
	fmt.Println(plus,minus)
  numbers(1,2,3,4,5,6)
	l := []int{1,2,3,4}
	numbers(l...)

	// if文
	num := 4
	if num%2 == 0 {
		fmt.Println(num,"is even")
	}else{
		fmt.Println(num,"is odd")
	}

	for i:=0;i<10;i++{
		if i==3{
			continue
		}
		if i == 6 {
			break
		}
		fmt.Println(i)
	}
   
strs := []string{"python","go","java"}

for i,v := range(strs){
	fmt.Println(i,v)
}

mp := map[string]int{"apple": 200, "banana": 300}
for k,v := range(mp){
	fmt.Println(k, v)
}

//switch文
OS := "mac"
switch OS{
case "mac":
		fmt.Println("MAC!!!!")
case "windows":
		fmt.Println("Windows")
default:
		fmt.Print("Others")
}

//deferは関数の最後に実行
// stack形式で蓄積される
defer fmt.Println("Hello1")
defer fmt.Println("Hello2")
fmt.Println("Hello3")

//確実にfileを閉じれる
// file,_ := os.Open("./main.go")
// defer file.Close()
// data := make([]byte, 1000)
// file.Read(data)
// fmt.Println(string(data))

//ポインタ
var p *int = new(int)
fmt.Println(p)

//構造体

point := Point{x:1,y:2}
fmt.Println(point.x, point.y)

}