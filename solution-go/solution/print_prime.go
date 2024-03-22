package solution

import (
	"fmt"
	"math/big"
)

func PrintPrime() {
	for i := 1; i <= 100; i++ {
		if n := big.NewInt(int64(i)); n.ProbablyPrime(0) {
			fmt.Printf("Bilangan prima: %d\n", n)
		} else if i%9 == 0 {
			fmt.Printf("Kelipatan 9 ke-%d\n", i/9)
		}
	}
}
