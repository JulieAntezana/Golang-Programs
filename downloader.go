// Downloader

// The Challenge:

// You are making a file downloader.
// To make the downloads faster, you decide to use concurrency. Each file download will run as a separate Goroutine.

// To simulate a file download, the download() function needs to take the size of the file as an argument and count the sum of all integers from 0 to that number (inclusive).

// The given program takes three file sizes as inputs and calls the download() function as Goroutines for each file.
// Complete the program by declaring the download() function and sending the result of their operation to main() using channels. The results should be added together in main() and output.
// You can use channels to get the data from the download() functions and add them together in main().

package main
import "fmt"

// create a function with size and channel parameters
func download(s int, ch chan int) {
    var sum int
    sum = 0
    // Count the sum of all integers from 0 to size number inclusive
    for i:=0; i<=s; i++ {
        sum += i
    }
    // Send the the result of operation to main
    ch<-sum
}
func main() {
    // Make channels for each download file
    ch1 := make(chan int)
    ch2 := make(chan int)
    ch3 := make(chan int)
    // Take the file sizes from input 
    var s1, s2, s3 int
    fmt.Scanln(&s1)
    fmt.Scanln(&s2)
    fmt.Scanln(&s3)
    // Call the download() function as Goroutines for each file
    go download(s1, ch1)
    go download(s2, ch2)
    go download(s3, ch3)

    //output the sum of all results
    fmt. Println(<-ch2 + <-ch1 + <-ch3)
}
