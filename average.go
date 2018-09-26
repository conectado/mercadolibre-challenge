package main

import (
	"math/big"
	"sync"
)

const prec = 200

func sumAverage(operations []*operation) *big.Rat {
	switch n := len(operations); n {
	case 0:
		return createFloat(0)
	case 1:
		return createFloat(operations[0].amount)
	default:
		return parallelAverage(operations)
	}
}

func createFloat(a int64) *big.Rat {
	return new(big.Rat).SetInt64(a)
}

func parallelAverage(operations []*operation) *big.Rat {
	wg := new(sync.WaitGroup)
	var partial []*big.Rat
	n := len(operations)
	var iterationLimit int

	if n%2 == 0 {
		iterationLimit = n / 2
	} else {
		iterationLimit = (n - 1) / 2
	}

	partial = make([]*big.Rat, iterationLimit)
	wg.Add(iterationLimit)
	for i := 0; i < iterationLimit; i++ {
		go func(index int) {
			partial[index] = div(add(createFloat(operations[index*2].amount), createFloat(operations[index*2+1].amount)), createFloat(int64(n)))
			wg.Done()
		}(i)
	}

	wg.Wait()
	if n%2 != 0 {
		partial[len(partial)-1] =
			add(partial[len(partial)-1], div(createFloat(operations[len(operations)-1].amount), createFloat(int64(len(operations)))))
	}
	return parallelSum(partial)
}

func add(a *big.Rat, b *big.Rat) *big.Rat {
	return a.Add(a, b)
}

func div(a *big.Rat, b *big.Rat) *big.Rat {
	return a.Quo(a, b)
}

func parallelSum(values []*big.Rat) *big.Rat {
	if len(values) == 1 {
		return values[0]
	} else {
		var partial []*big.Rat
		wg := new(sync.WaitGroup)
		n := len(values)
		var iterationLimit int

		if n%2 == 0 {
			iterationLimit = n / 2
		} else {
			iterationLimit = (n - 1) / 2
		}

		partial = make([]*big.Rat, iterationLimit)
		wg.Add(iterationLimit)
		for i := 0; i < iterationLimit; i++ {
			go func(index int) {
				partial[index] = add(values[index*2], values[index*2+1])
				wg.Done()
			}(i)
		}

		wg.Wait()
		if n%2 != 0 {
			partial[len(partial)-1] = add(partial[len(partial)-1], values[len(values)-1])
		}

		return parallelSum(partial)
	}
}
