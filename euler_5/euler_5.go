
package main

import (
    "fmt"
)

func main() {
    CommonMultiple := 1
    for i := 1; i <= 20 ; i++ {
        if CommonMultiple % i == 0 {
            continue
        } else {
            for _, factor := range Factors(i) {
                for CommonMultiple % factor == 0 {
                    CommonMultiple /= factor
                }
            }
            CommonMultiple *= i
        }
    }
    fmt.Println(CommonMultiple)
}

func Factors(number int) (factors []int) {
    for i := 2; i <= number/2; i++ {
        if number % i == 0 {
            factors = append(factors, i)
        }
    }
    return
}
