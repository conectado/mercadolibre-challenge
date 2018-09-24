package main

import (
	"fmt"
)

func main() {
	_, pago, cobro, descuento, inversion, movements := parseFile("./movements.log")
	pagoAvg := sumAverage(pago)
	cobroAvg := sumAverage(cobro)
	descuentoAvg := sumAverage(descuento)
	inversionAvg := sumAverage(inversion)
	fmt.Printf("Pago average: %s\n", pagoAvg.FloatString(6))
	fmt.Printf("Cobro average: %s\n", cobroAvg.FloatString(6))
	fmt.Printf("Descuento average: %s\n", descuentoAvg.FloatString(6))
	fmt.Printf("Inversión average: %s\n", inversionAvg.FloatString(6))
	user, value := mapMax(movements)
	fmt.Printf("User with most movements is %s and has %d movements\n", user, value)
}

func mapMax(val map[string]int) (string, int) {
	currUser := ""
	curr := 0
	for key, value := range val {
		if value >= curr {
			curr = value
			currUser = key
		}
	}

	return currUser, curr
}

type operation struct {
	user   string
	kind   string
	amount int64
}
