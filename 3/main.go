package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
)

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")
	input := getInputNumeric(bytes)

	found := make(chan int)
	failed := make(chan struct{})
	var wg sync.WaitGroup

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			wg.Add(1)

			go func(noun, verb int, input []int, wg *sync.WaitGroup) {
				defer wg.Done()

				tmp := make([]int, len(input))
				copy(tmp, input)
				tmp[1] = noun
				tmp[2] = verb

				result := Run(tmp)
				if result == 19690720 {
					found <- (noun * 100) + verb
				}
			}(noun, verb, input, &wg)
		}
	}

	go func(wg *sync.WaitGroup, failed chan struct{}) {
		wg.Wait()
		failed <- struct{}{}
	}(&wg, failed)

	select {
	case <-failed:
		fmt.Println("failed")
	case result := <-found:
		fmt.Println(result)
	}
}

func getInputNumeric(input []byte) []int {
	s := string(input)
	values := strings.Split(s, ",")
	length := len(values)
	numeric := make([]int, length)

	for index, str := range values {
		intVal, _ := strconv.Atoi(str)
		numeric[index] = intVal
	}

	return numeric
}

func Run(ic []int) int {
	for i := 0; i < len(ic); i += 4 {
		opcode := ic[i]
		if opcode == 99 {
			break
		}

		positionOne := ic[i+1]
		positionTwo := ic[i+2]
		positionToUpdate := ic[i+3]

		if len(ic) <= positionToUpdate || len(ic) <= positionOne || len(ic) <= positionTwo {
			break
		}

		valueOne := ic[positionOne]
		valueTwo := ic[positionTwo]

		switch opcode {
		case 1:
			total := valueOne + valueTwo
			ic[positionToUpdate] = total
			continue
		case 2:
			total := valueOne * valueTwo
			ic[positionToUpdate] = total
			continue
		}
	}

	return ic[0]
}
