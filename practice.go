package main

import "fmt"

// Variadic Functions

func f(v ...int) {
	res := 20
	for _, a := range v {
		res -= a
	}
	fmt.Println(res)
}

func sum(nums ...int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	fmt.Println(total)
}

func main() {

	// Practice Variadic Functions

	// s := []int{1, 2, 4, 6, 8}
	// s[2] = s[1]
	// s[3] = s[2]+s[0]
	// fmt.Println(s[4] % s[3])

	// Variadic function that calculates
	// and outputs the sum of its arguments:

	// total = sum(2, 4, 6)

	// Variadic functions are functions that can be called with any number of arguments.
	// Example: fmt.Println() has any number of comma-
	// separated arguments output separated by spaces.
	// fmt.Println("a", "b", "c")

	// Variadic Function Slices
	// s: []int{42, 33, 22, 11}
	// sum(s...)
	// v := []int{8, 5, 3}
	// f(v...)
	//Answer: 4

	// Maps Practice
	// Maps are also called dictionaries, associative arrays, or hash tables.

	// What is the output of this code?

	// m := map[int]int{
	// 8: 42,
	// 2: 6,
	// 4: 9,
	// 5: 3}
	// delete(m, 2)
	// fmt.Println(m[4]-m[5])
	// Answer: 6

	// Delete an element from the map
	// m := map[string]int {
	// "James": 42,
	// "Amy": 24}
	// delete(m, "James")
	// m["Bob"] = 8
	// Initialize a map using this syntax. If the requested key does not exist in the map, a zero value will be returned.
	// m := map[string]int {
	// "James": 42,
	// "Amy": 24}
	// fmt.Println(m["Amy"])

	// Map a string key to an integer value

	// m := make(map[string]int)
	// m["James"] = 42
	// m["Amy"] = 24
	// fmt.Println(m["James"])

	// Practice Range
	//  res := 0
	//  nums := [3]int{2, 4, 6}
	//  for i, v := range nums {
	//      if i%2==0 {
	//          res += v
	//      }
	//  }
	//  fmt.Println(res)

	// iterate over characters of a string to
	// output unicode code points of the chars.
	// Use Printf() function to output actual
	// characters.
	// x := "hello"
	// for _, c := range x {
	//     //fmt.Println(c)
	//     fmt.Printf("%c \n", c)
	//     //fmt.Println(c)
	// }
	// loop that iterates over the values of
	// the array 'nums' and calculates their sum.
	// sum := 0
	// for _, v := range nums {
	//     sum += v
	// }

	// Range for loop
	// a := make([]int, 5)
	// a[1] = 2
	// a[2] = 3
	// for _, v := range a {
	// fmt.Println(v)
	// }

	// a := make([]int, 5)
	// a[1] = 2
	// a[2] = 3
	// for i, v := range a {
	//     fmt.Println(i, v)
	// }

	// Practice Slices
	// make() function

	// a := make([]int, 5)
	// a = append(a, 8)
	// fmt.Println(a)

	// Slices
	// a := [5]int{1, 2, 4, 6, 8}
	// var s []int = a[1:3]
	// fmt.Println(s[1])

	// s[0] = 8
	// fmt.Println(a)

	// Practice Arrays
	// arr := [5]int{1, 2, 4, 6, 8}
	// fmt.Println(arr[2])

	// var a [5]int
	// a[0] = 8
	// a[1] = 42

	// fmt.Println(a[1])

	// What is the output of this code? (1)

	// res := 0
	// for i:=0;i<10;i+=3 {
	//     res += i
	// }

	// fmt.Println(res)

	// Create a valid if/else statement.

	// height := 195
	// if height > 185 {
	//     fmt.Println("Yes")
	// } else {
	//     fmt.Println("No")
	// }

	// What is the output of this code? (1)

	// x :=7

	// res:=0

	// switch {
	//     case x>8:
	//         res++
	//     case x>3 && x<10:
	//         res++
	//     case X>5:
	//         res++
	// }
	// fmt.Println(res)

	// Creat a loop that outputs the numbers 9 to 1.

	// for i:=9; i>0; i-- {
	//     fmt.Println(i)
	// }

	//  Calculate the sum of numbers from 1 to 1000.
	// sum := 1
	// res := 0
	// for sum <= 1000 {
	//     res += sum
	//     sum++
	// }
	// fmt.Println(res)

	//      package main

	//      import "fmt"

	//      func main() {

	// var i, j, side int

	// fmt.Print("Enter Any Side of a Square = ")
	// fmt.Scanln(&side)

	// fmt.Println("Hollow Square Star Pattern")
	// for i = 0; i < side; i++ {
	//     for j = 0; j < side; j++ {
	//         if i == 0 || i == side-1 || j == 0 || j == side-1 {
	//             fmt.Print("* ")
	//         } else {
	//             fmt.Print("  ")
	//         }
	//     }
	//     fmt.Println()
	// }

	// sum:=0

	// for i:=1; i<=3; i++ {
	//     sum += i
	// }
	// fmt.Println(sum)

	// for i:=0; i<5; i++ {
	//     fmt.Println(i)
	// }

	// Problem
	// Temperature Checker

	// You are writing a program for a temperature checking system at an airport.
	// The system measures the body temperature of a person and needs to output "Allowed" if it is in the normal range, or "Fever" if it is higher than normal.
	// Normal: up to 99.5 °F
	// Fever: > 99.5 °F

	// The program should take the temperature as a float from input and output the corresponding message.
	// Sample Input:
	// 101.3

	// Sample Output:
	// Fever
	// Use an if/else statement to make the decision.

	// package main

	// import "fmt"

	// func main() {

	// var input float32

	// fmt.Scanln(&input)
	// if input <= 99.5 {
	//     fmt.Println("Allowed")
	// }   else {
	//     fmt.Println("Fever")
	// }

	// }

	// x := 2

	// switch {
	//     case x>0 && x<10:
	//         fmt.Println("something")
	//     case x > 10:
	//         fmt.Println("something else")
	// }

	// num := 3

	// switch num {
	// case 1:
	//     fmt.Println("One")
	// case 2:
	//     fmt.Println("Two")
	// case 3:
	//     fmt.Println("Three")
	// default:
	//     fmt.Println(num)
	// }

	// num := 3

	// if num == 1 {
	//     fmt.Println("One")
	// } else if num == 2 {
	//     fmt.Println("Two")
	// } else if num == 3 {
	//     fmt.Println("Three")
	// } else {
	//     fmt.Println(num)
	// }

	//     if x := 42; x > 18 {
	//     fmt.Println("Allowed")
	// }   else {
	//     fmt.Println("Not allowed")
	// }

	// x := 14

	// if x > 18 {
	//     fmt.Println("Allowed")
	// }   else {
	//     fmt.Println("Not allowed")
	// }

	// x := 42

	// if x > 18 {
	//     fmt.Println("Allowed")
	// }

	// var input1 int
	// var input2 int

	// fmt.Scanln(&input1)
	// fmt.Scanln(&input2)

	// fmt.Println(input1 + input2)

	// var input string
	// fmt.Scanln(&input)
	// fmt.Println(input)

	// age := 15

	// res := (age > 18) || age == 0

	// fmt.Println(res)

	// x :=42
	// y := 8

	// // Logical AND
	// fmt.Println(x != y && x<=y)

	// // Logical OR
	// fmt.Println(x != y || x<=y)

	// // Logical NOT
	// fmt.Println(!(x>y))

	// x :=42
	// y := 8

	// fmt.Println(x == y)
	// fmt.Println(x != y)
	// fmt.Println(x > y)
	// fmt.Println(x < y)

	// x := "Hello "
	// y := "world!"

	// fmt.Println(x+y)

	// x :=42
	// y := 8

	// x += y
	// fmt.Println(x)

	// x *= y
	// fmt.Println(x)

	// // Addition
	// res := x + y
	// fmt.Println(res)

	// // Subtraction
	// res = x - y
	// fmt.Println(res)

	// // Multitplication
	// res = x * y
	// fmt.Println(res)

	// // Division
	// res = x / y
	// fmt.Println(res)

	// // Modulus, results in the remainder of the division
	// res = x % y
	// fmt.Println(res)

	// fmt.Println(name)
	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(online)
	// fmt.Println(pi)
	fmt.Println("Happy coding!")
}
