
package primes

import (
    "math"
)

func FindNth(nth int) (prime int) {
    if nth == 1 {
        prime = 2
        return
    }
    prime = 3       // Start with 2,3 already known.
    count := 2
    for count < nth {
        prime_root := int(math.Sqrt(float64(prime)))
        // min(j) * min(k) == root(prev) * root(prev)
        // max(j) * max(j) == prev       * root(prev)
        // min(j) * max(k) == root(prev) * root(prev)
        // max(j) * min(k) == prev       * 1
        // range of i,j == [prev: prev * root(prev)]
        next_slice := prime * (prime - 1)
        composites := make([]bool, next_slice)
        for j := prime_root; j <= prime; j++ {
            for k := prime / j; k <= prime * prime_root; k++ {
                switch x := j*k; {
                    case x < prime: continue
                    case x > prime * prime: break
                    case x % 2 == 0: continue
                    default: composites[(x - prime)/2] = true
                }
            }
        }
        last_index := 0
        for index, found := range composites {
            if !found && !IsPrime(prime + 2 * (index - last_index)) {
                count++
                prime += 2 * (index - last_index)
                last_index = index
            }
            if count == nth {
                return
            }
        }
    }
    return
}

func IsPrime(number int) bool {
    root := int(math.Sqrt(float64(number)))
    for i := 2; i <= root; i++ {
        if number % i == 0 {
            return true
        }
    }
    return false
}
