package solutions

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func SolveTwo() {
	f, err := os.Open("three/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	wireOneInput, _ := reader.ReadString('\n')
	wireTwoInput, _ := reader.ReadString('\n')

	wireOne := newWire(wireOneInput)
	wireTwo := newWire(wireTwoInput)

	wireOne.resolve()
	wireTwo.resolve()

	shortest := wireOne.findShortestCross(wireTwo)
	fmt.Println(shortest)
}

func (w wire) findShortestCross(ow wire) int {
	shortest := math.MaxInt32

	for point, step := range w.visited {
		isCenter := point.x == 0 && point.y == 0
		if isCenter {
			continue
		}


		if minStep, ok := ow.visited[point]; ok {
			if (minStep + step) < shortest {
				shortest = minStep + step
			}
		}
	}
	return shortest
}
