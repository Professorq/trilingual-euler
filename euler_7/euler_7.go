
// the test-prime algorithm would be 
// 1) for each number
// 2)   for each smaller_number <= Sqrt(number)
// 3)       test number % smaller_number == 0
// n^2

// Instead of testing each number to see if it's prime
// let's try to find all non-prime numbers

// the 'fill non-prime algorithm' is:
// for len(found_primes) < 10002:
//      for i := Sq
//          
// This is O(n^2), but probably a lot closer to n

package main

import (
    "fmt"
    "./primes"
    "os"
    "strconv"
)

func main() {
    number, err := strconv.Atoi(os.Args[1])
    if err != nil {
        panic(err)
    }
    result := primes.FindNth(number)
    fmt.Println(result)
}
