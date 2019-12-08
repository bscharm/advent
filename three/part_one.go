package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
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

	crosses := wireOne.findCrosses(wireTwo)
	shortest := findShortest(crosses)
	fmt.Println(shortest)

}

type wire struct {
	instructions []string
	currentPoint point
	visited      map[point]bool
}

type point struct {
	x int
	y int
}

func (w wire) resolve() {
	for _, instruction := range w.instructions {
		w.followInstruction(instruction)
	}
}

func (w *wire) followInstruction(instruction string) {
	direction := instruction[:1]
	distance, _ := strconv.Atoi(instruction[1:])
	currentPoint := w.currentPoint

	switch direction {
	case "R":
		for i := 0; i < distance; i++ {
			newX := w.currentPoint.x + 1
			newPoint := point{newX, currentPoint.y}
			w.visited[newPoint] = true
			w.currentPoint = newPoint
		}
	case "L":
		for i := 0; i < distance; i++ {
			newX := w.currentPoint.x - 1
			newPoint := point{newX, currentPoint.y}
			w.visited[newPoint] = true
			w.currentPoint = newPoint
		}
	case "U":
		for i := 0; i < distance; i++ {
			newY := w.currentPoint.y + 1
			newPoint := point{currentPoint.x, newY}
			w.visited[newPoint] = true
			w.currentPoint = newPoint
		}
	case "D":
		for i := 0; i < distance; i++ {
			newY := w.currentPoint.y - 1
			newPoint := point{currentPoint.x, newY}
			w.visited[newPoint] = true
			w.currentPoint = newPoint
		}

	}
}

func findShortest(points []point) int {
	distance := math.MaxInt32
	for _, point := range points {
		total := absolute(point.x) + absolute(point.y)
		if total < distance {
			distance = total
		}
	}
	return distance
}

func absolute(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (w wire) findCrosses(ow wire) []point {
	var crosses []point
	for point := range w.visited {
		isCenter := point.x == 0 && point.y == 0
		if isCenter {
			continue
		}

		if ow.visited[point] {
			crosses = append(crosses, point)
		}
	}
	return crosses
}

func newWire(s string) wire {
	instructions := strings.Split(s, ",")
	startingPoint := point{0, 0}
	visited := make(map[point]bool)
	visited[startingPoint] = true
	return wire{
		instructions,
		startingPoint,
		visited,
	}
}
