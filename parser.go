package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseFile(fileName string) (
	all []*operation,
	pago []*operation,
	cobro []*operation,
	descuento []*operation,
	inversion []*operation,
	movements map[string]int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	movements = make(map[string]int)
	scanner := bufio.NewScanner(file)
	// TODO paralelize prior 2
	// TODO return multiple values per operations
	for scanner.Scan() {
		if res, err := parseLine(scanner.Text()); err == nil {
			all = append(all, res)
			movements[res.user] += 1
			switch res.kind {
			case "pago":
				pago = append(pago, res)
			case "cobro":
				cobro = append(cobro, res)
			case "descuento":
				descuento = append(descuento, res)
			case "inversión":
				inversion = append(inversion, res)
			}

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func parseLine(line string) (*operation, error) {
	var params [3]string
	for i, val := range strings.Split(line, "] ") {
		partial := strings.Split(val, ":")
		if strings.TrimPrefix(partial[0], "[") == "user" ||
			strings.TrimPrefix(partial[0], "[") == "ammount" ||
			strings.TrimPrefix(partial[0], "[") == "type" {
			params[i] = strings.TrimSuffix(partial[1], "]")
		}
	}
	return createOperation(params[0], params[1], params[2])
}

func createOperation(user string, kind string, amount string) (*operation, error) {
	if user == "" || kind == "" || amount == "" {
		return nil, errors.New("Not an operation")
	}
	op := new(operation)
	op.user = user
	op.kind = kind
	partialAmount, _ := strconv.Atoi(amount)
	op.amount = int64(partialAmount)
	return op, nil
}
