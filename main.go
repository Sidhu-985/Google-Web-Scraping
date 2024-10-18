package main

import "fmt"

type point struct{
	x,y int
}

func main() {
// 	intArr := [...]int32{1,2,3}
// 	fmt.Println(intArr)

// 	var intSlice []int32 = []int32{2,3,4}
// 	intSlice = append(intSlice,5)
// 	fmt.Printf("The length is %v and capacity = %v\n",len(intSlice),cap(intSlice))

// 	var myMap2 = map[string]int64{"Sidhu":23,"Sfu":231}
// 	fmt.Println(myMap2["Sidhu"])

// 	var num,ok = myMap2["Sfu"]

// 	if ok{
// 		fmt.Println(num)
// 	}

// 	for name,age := range myMap2{
// 		fmt.Printf("Name: %v, Age: %v\n",name,age)
// 	}

// 	for i,number:= range intArr{
// 		fmt.Printf("Index: %v, Number:%v\n",i,number)
// 	}

// 	for s:=0;s<len(myMap2);s++{
// 		fmt.Println(s)
// 	}
// 

//   var myString = "rèsumè"
//   var indexed = myString[0]

//   fmt.Printf("%v,%T\n",indexed,indexed)

//   for i,v := range myString{
// 	fmt.Println(i,v)
//   }

//   p:=point(1)

//   fmt.Println(p)

x := 0
switch x {
case 0:
	fmt.Println(69)
case 1, 2:
    fmt.Println("Multiple matches")
case 42:   // Don't "fall through".
    fmt.Println("reached")
case 43:
    fmt.Println("Unreached")
default:
    fmt.Println("Optional")
}

}

// package main

// import(
// 	"fmt"
// )

// type Marks struct{
// 	x,y int32
// }

// func (m Marks) Total() int32{
// 	return m.x+m.y
// }

// func main(){

// 	var m,n int32
// 	fmt.Scan(&m,&n)

// 	Mark := Marks{m,n}

// 	fmt.Println(Mark.Total())
// }

// package main

// import "fmt"

// func main(){
// 	var p *int32
// 	var i int32
// 	p=&i

// 	fmt.Println(p)
// 	fmt.Println(i)
// }
