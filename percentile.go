package main

import "sort"

func calcPercentile(operations []*operation, percentile int) int64 {
	operationAmount := getAmounts(operations)
	sort.Slice(operationAmount, func(i, j int) bool { return operationAmount[i] < operationAmount[j] })
	return operationAmount[len(operationAmount)*percentile/100-1]
}

func getAmounts(operations []*operation) []int64 {
	ammounts := make([]int64, len(operations))
	for i, oper := range operations {
		ammounts[i] = oper.amount
	}
	return ammounts
}

//TODO For a parallel sort to work fast we need to parallelize the merge
func parallelSort(values []int64) []int64 {
	ch := make(chan []int64)
	go mergeSort(values[0:len(values)/2-1], values[len(values)/2:len(values)-1], ch)
	return <-ch
}

func mergeSort(left []int64, right []int64, ch chan []int64) {
	chaL := make(chan []int64)
	chaR := make(chan []int64)

	if len(left) > 1 {
		go mergeSort(left[0:len(left)/2-1], left[len(left)/2:len(left)-1], chaL)
	} else {
		go func(c chan []int64) {
			c <- left
		}(chaR)
	}

	if len(right) > 1 {
		go mergeSort(right[0:len(right)/2-1], right[len(right)/2:len(right)-1], chaR)
	} else {
		go func(c chan []int64) {
			c <- right
		}(chaL)
	}

	l, r := <-chaL, <-chaR
	ch <- merge(l, r)
}

func merge(left []int64, right []int64) []int64 {
	var res []int64
	for _, l := range left {
		for _, r := range right {
			if r < l {
				res = append(res, r)
			}
		}
		res = append(res, l)
	}

	return res
}
