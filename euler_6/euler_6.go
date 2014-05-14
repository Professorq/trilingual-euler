
package main

func main() {
    sum_of_squares := 0
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
        sum_of_squares += i * i
    }
    square_of_sum := sum * sum
    print(square_of_sum - sum_of_squares)
}
