package fibonacci

import (
	"math/big"
)

func Fib(N uint) *big.Int {
	var table []*big.Int
	table = make([]*big.Int, N+1)
	table[0] = new(big.Int).SetInt64(0)
	table[1] = new(big.Int).SetInt64(1)

	for i := uint(2); i <= N; i += 1 {
		table[i] = new(big.Int).Add(table[i-1], table[i-2])
	}

	return table[N]
}
