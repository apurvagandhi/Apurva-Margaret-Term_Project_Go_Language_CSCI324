package main

import (
	"fmt"
	"strconv"
)

//Enumarated Constants
const (
	A = iota
	B = iota
	C = iota
)
const (
	A1 = iota
	B1
	C1
)

const (
	_ = iota
	A2
	B2
	C2
)
const (
	_ = iota + 5
	A3
	B3
	C3
)

//Variable declaration at package level
var outsideMain int = 41

//You can declare variables in blocks so you don’t have to keep typing
//the keyword var before each one
var (
	actorName string = "Anne"
	number    int    = 3
	season    int    = 11
)

func main() {

	//Print Statement
	fmt.Println("Hello, World!")

	/*******Variable Declaration in GO******/
	//useful when you want to declare then do stuff before assigning it
	var a int
	a = 42

	//if go cant guess the type on its own
	var b int = 43

	//fast and good unless you want variable to be specific type i.e. 20 to be a float
	c := 44

	var n bool = true

	n1 := 1 == 1

	var n2 int

	var n3 uint16 = 42
	n6 := strconv.Itoa(n2)
	//Checking Variable Declaration
	fmt.Println("Zeroth way to declare variable: ", a)
	fmt.Println("First way to declare variable: ", a)
	fmt.Println("second way to declare variable: ", b)
	fmt.Printf("%v, %T", c, c)
	fmt.Println(actorName)
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", n1, n1)
	fmt.Printf("%v, %T\n", n2, n2)
	fmt.Printf("%v, %T\n", n6, n6)

	fmt.Printf("%v, %T\n", n3, n3)
	n4 := 10 //1010
	n5 := 3  //0011
	fmt.Println("n4 is ", n4)
	fmt.Println("n5 is ", n5)
	fmt.Println("n4 + n5 =", n4+n5)
	fmt.Println("n4 - n5 =", n4-n5)
	fmt.Println("n4 * n5 =", n4*n5)
	fmt.Println("n4 / n5 =", n4/n5)
	fmt.Println("n4 % n5 =", n4%n5)

	//Bit operators
	fmt.Println("\nBit operators")
	fmt.Println("n4 & n5 =", n4&n5) //0010
	fmt.Println("n4 | n5 =", n4|n5) //1011
	fmt.Println("n4 ^ n5 =", n4^n5) //1001
	fmt.Println("n4 &^ n5 =", n4&^n5)

	//Bit shift
	fmt.Println("\nBit operators")
	fmt.Println("n4 << 2 =", n4<<2)
	fmt.Println("n4 >> 2 =", n4>>2)

	var n7 float32 = 3.14
	fmt.Println("\nFloat Type")
	fmt.Printf("%v, %T\n", n7, n7)

	var n8 complex64 = 1 + 2i
	var n9 complex64 = complex(5, 12)
	fmt.Println("\nComplex Type")
	fmt.Printf("%v, %T\n", n8, n8)
	fmt.Printf("%v, %T\n", real(n8), real(n8))
	fmt.Printf("%v, %T\n", imag(n8), imag(n8))
	fmt.Printf("%v, %T\n", n9, n9)

	//String
	fmt.Println("\n String Type")
	s := "this is a string"
	fmt.Printf("\n%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", s[2], s[2])
	fmt.Printf("%v, %T\n", string(s[2]), s[2])

	//String concatination
	fmt.Println("\nString Concatination")
	S3 := "this is a string "
	S2 := "this is another string"
	S4 := S3 + S2
	fmt.Printf("\n%v, %T\n", S4, S4)

	fmt.Println("\nByte Slice")
	S5 := []byte(S3)
	fmt.Printf("\n%v, %T\n", S5, S5)

	//Declaration of constant
	const myCost int = 42
	const myCost1 bool = true
	const myCost2 float32 = 42.32
	const myCost3 string = "foo"

	fmt.Println("\nConstant Declaration and Practice")
	fmt.Printf("\n%v, %T\n", myCost, myCost)
	fmt.Printf("\n%v, %T\n", myCost1, myCost1)
	fmt.Printf("\n%v, %T\n", myCost2, myCost2)
	fmt.Printf("\n%v, %T\n", myCost3, myCost3)

	//printing enumarated Constants
	fmt.Println("Printing Enumarated Constants")
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

	fmt.Println(A1)
	fmt.Println(B1)
	fmt.Println(C1)

	fmt.Println("ignoring first 0")
	fmt.Println(A2)
	fmt.Println(B2)
	fmt.Println(C2)

	fmt.Println("Adding +5 to iota and ignoring first 0")
	fmt.Println(A3)
	fmt.Println(B3)
	fmt.Println(C3)

}
