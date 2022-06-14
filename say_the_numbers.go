package main

import "fmt"

func main() {

	// Say The Numbers

	// You are making a robot that can speak numbers.
	// Your robot should take 3 numbers in the range of 0-10 as input and output the corresponding texts in English.

	// Sample Input:
	// 8
	// 0
	// 5

	// Sample Output:
	// Eight
	// Zero
	// Five
	// Note, that each number should be output on a new line, with a capital letter.

	var input1, input2, input3 int

	//fmt.Println("Enter 3 numbers in the range of 0-10.")

	fmt.Scanln(&input1)
	fmt.Scanln(&input2)
	fmt.Scanln(&input3)

	num1 := input1

	switch {
	case num1 == 0:
		fmt.Println("Zero")
	case num1 == 1:
		fmt.Println("One")
	case num1 == 2:
		fmt.Println("Two")
	case num1 == 3:
		fmt.Println("Three")
	case num1 == 4:
		fmt.Println("Four")
	case num1 == 5:
		fmt.Println("Five")
	case num1 == 6:
		fmt.Println("Six")
	case num1 == 7:
		fmt.Println("Seven")
	case num1 == 8:
		fmt.Println("Eight")
	case num1 == 9:
		fmt.Println("Nine")
	case num1 == 10:
		fmt.Println("Ten")
	}

	num2 := input2

	switch {
	case num2 == 0:
		fmt.Println("Zero")
	case num2 == 1:
		fmt.Println("One")
	case num2 == 2:
		fmt.Println("Two")
	case num2 == 3:
		fmt.Println("Three")
	case num2 == 4:
		fmt.Println("Four")
	case num2 == 5:
		fmt.Println("Five")
	case num2 == 6:
		fmt.Println("Six")
	case num2 == 7:
		fmt.Println("Seven")
	case num2 == 8:
		fmt.Println("Eight")
	case num2 == 9:
		fmt.Println("Nine")
	case num2 == 10:
		fmt.Println("Ten")
	}

	num3 := input3

	switch {
	case num3 == 0:
		fmt.Println("Zero")
	case num3 == 1:
		fmt.Println("One")
	case num3 == 2:
		fmt.Println("Two")
	case num3 == 3:
		fmt.Println("Three")
	case num3 == 4:
		fmt.Println("Four")
	case num3 == 5:
		fmt.Println("Five")
	case num3 == 6:
		fmt.Println("Six")
	case num3 == 7:
		fmt.Println("Seven")
	case num3 == 8:
		fmt.Println("Eight")
	case num3 == 9:
		fmt.Println("Nine")
	case num3 == 10:
		fmt.Println("Ten")
	}

}

// Another Solution
// package main
// import "fmt"

// func main() {
//     var vin int
//     var i int

//     for i=0;i<3;i++ {
//         n, err := fmt.Scanln(&vin)
//         if n==0 || err!=nil {
//             break
//         }
//         switch vin {
//             case 0:fmt.Print("Zero")
//             case 1:fmt.Print("One")
//             case 2:fmt.Print("Two")
//             case 3:fmt.Print("Three")
//             case 4:fmt.Print("Four")
//             case 5:fmt.Print("Five")
//             case 6:fmt.Print("Six")
//             case 7:fmt.Print("Seven")
//             case 8:fmt.Print("Eight")
//             case 9:fmt.Print("Nine")
//             case 10:fmt.Print("Ten")
//             default:
//         }
//         if i<2 {
//             fmt.Println("")
//         }
//     }
// }
