
package primes

import(
    "testing"
)

func TestFind6th(t *testing.T) {
    const nth, expected = 6, 13
    result := FindNth(nth)
    if result != expected {
        t.Errorf("%vth prime is %v, not %v. Found %v",
                 nth, expected, result, result)
    }
}

func TestFind20th(t *testing.T) {
    const nth, expected = 20, 71
    result := FindNth(nth)
    if result != expected {
        t.Errorf("%vth prime is %v, not %v. Found %v",
                 nth, expected, result, result)
    }
}
