package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var expected int

	if len(os.Args) <= 1 {
		expected = 19690720
	} else {
		expected, _ = strconv.Atoi(os.Args[1])
	}

	inputString := "1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,5,19,23,2,10,23,27,1,27,5,31,2,9,31,35,1,35,5,39,2," +
		"6,39,43,1,43,5,47,2,47,10,51,2,51,6,55,1,5,55,59,2,10,59,63,1,63,6,67,2,67,6,71,1,71,5,75,1,13,75,79,1,6," +
		"79,83,2,83,13,87,1,87,6,91,1,10,91,95,1,95,9,99,2,99,13,103,1,103,6,107,2,107,6,111,1,111,2,115,1,115,13,0," +
		"99,2,0,14,0"

	found := make(chan int)
	failed := make(chan struct{})
	var wg sync.WaitGroup

	for noun := 0; noun <= 99; noun++ {
		wg.Add(1)
		for verb := 0; verb <= 99; verb++ {
			wg.Add(1)
			go func(noun, verb, expected int) {
				ic := IntCode{}
				ic.InitiateFromString(inputString)
				ic.SetNounVerb(noun, verb)
				result, _ := ic.Run()
				if result == expected {
					found <- (noun * 100) + verb
				}
				wg.Done()
			}(noun, verb, expected)
		}
	}

	go func(wg sync.WaitGroup, failed chan struct{}) {
		wg.Wait()
		failed <- struct{}{}
	}(wg, failed)

	select {
	case result := <-found:
		fmt.Println(result)
	case <-failed:
		_, _ = fmt.Fprintln(os.Stderr, "failed")
	}
}

type IntCode struct {
	sequence []int
	debug    bool
	log      log.Logger
}

func (ic *IntCode) SetNounVerb(noun, verb int) {
	ic.sequence[1] = noun
	ic.sequence[2] = verb
}

func (ic *IntCode) Run() (int, error) {
	for i := 0; i < len(ic.sequence); i += 4 {
		if ic.debug {
			ic.log.Printf("index: %d\n", i)
			ic.log.Printf("instruction set: %d\n", i/4)
		}

		opcode := ic.sequence[i]
		if ic.debug {
			ic.log.Printf("opcode: %d\n", opcode)
		}

		if opcode == 99 {
			if ic.debug {
				ic.log.Println()
			}
			break
		}

		positionOne := ic.sequence[i+1]
		if ic.debug {
			ic.log.Printf("position one: %d\n", positionOne)
		}

		positionTwo := ic.sequence[i+2]
		if ic.debug {
			ic.log.Printf("position two: %d\n", positionTwo)
		}

		valueOne := ic.sequence[positionOne]
		if ic.debug {
			ic.log.Printf("value one: %d\n", valueOne)
		}

		valueTwo := ic.sequence[positionTwo]
		if ic.debug {
			ic.log.Printf("value two: %d\n", valueTwo)
		}

		positionToUpdate := ic.sequence[i+3]
		if ic.debug {
			ic.log.Printf("position to update: %d\n", positionToUpdate)
		}

		if opcode == 1 {
			total := valueOne + valueTwo
			ic.sequence[positionToUpdate] = total
			if ic.debug {
				ic.log.Println("operation: addition")
				ic.log.Printf("setting position %d to %d\n\n", positionToUpdate, total)
			}
			continue
		}

		if opcode == 2 {
			total := valueOne * valueTwo
			ic.sequence[positionToUpdate] = total
			if ic.debug {
				ic.log.Println("operation: multiplication")
				ic.log.Printf("setting position %d to %d\n\n", positionToUpdate, total)
			}
			continue
		}

		return 0, fmt.Errorf("something went wrong")
	}

	return ic.sequence[0], nil
}

func (ic *IntCode) EnableDebug() {
	ic.debug = true
}

func (ic *IntCode) InitiateFromString(s string) {
	values := strings.Split(s, ",")
	length := len(values)
	valuesNumeric := make([]int, length)

	for index, str := range values {
		intVal, _ := strconv.Atoi(str)
		valuesNumeric[index] = intVal
	}

	ic.log.SetOutput(os.Stdout)
	ic.sequence = valuesNumeric
}
