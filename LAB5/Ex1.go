package main

import (
	"fmt"
	"math"
	"os"
)

const radius = 15
const eps = 0.00001

type Point struct{ X, Y float64 }

func readInput(filename string) ([]Point, error) {
	f, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer f.Close()

	var n int
	if _, err = fmt.Fscan(f, &n); err != nil {
		return nil, err
	}

	points := make([]Point, n)
	for i := 0; i < n; i++ {
		if _, err = fmt.Fscan(f, &points[i].X, &points[i].Y); err != nil {
			return nil, err
		}
	}

	return points, nil
}

func dist(a, b Point) float64 {
	dx := b.X - a.X
	dy := b.Y - a.Y
	d := dx*dx + dy*dy
	return d
}

func step(points []Point) []Point {
	newPoints := make([]Point, len(points))
	for i := range points {
		sumX, sumY := points[i].X, points[i].Y
		count := 1
		for j := range points {
			if j == i {
				continue
			}

			if dist(points[i], points[j]) < (radius * radius) {
				sumX += points[j].X
				sumY += points[j].Y
				count++
			}
		}
		newPoints[i] = Point{sumX / float64(count), sumY / float64(count)}
	}
	return newPoints
}

func merge(points []Point) []Point {
	var kept []Point
	for i := range points {
		duplicate := false
		for j := range kept {
			if dist(points[i], kept[j]) < eps {
				duplicate = true
				break
			}
		}
		if !duplicate {
			kept = append(kept, points[i])
		}
	}
	return kept
}

func main() {
	points, err := readInput("input.txt")

	if err != nil {
		fmt.Println("Eroare:", err)
		return
	}

	fmt.Println(points)
	const maxIter = 1000

	for t := 0; t < maxIter && len(points) > 1; t++ {
		next := merge(step(points))

		if len(next) == len(points) {
			moved := 0.0
			for i := range next {
				moved = math.Max(moved, dist(points[i], next[i]))
			}
			if moved < eps {
				points = next
				break
			}
		}
		points = next
		fmt.Printf("Iteration %d: %.2f\n", t, points)
	}
	fmt.Printf("Final: %.2f\n", points)

}
