package main

import (
	"fmt"
	"sort"
)

type Fruit struct {
	Name string
	X, Y float64
}

type neighbor struct {
	Name     string
	Distance float64
}

func FindNearestFruit(fruits []Fruit, x, y float64, k int) string {
	if k <= 0 || k > len(fruits) {
		return fmt.Sprintf("Invalid k, use 1..%d", len(fruits))
	}

	neighbors := make([]neighbor, len(fruits))
	for i, f := range fruits {
		neighbors[i] = neighbor{f.Name, (x-f.X)*(x-f.X) + (y-f.Y)*(y-f.Y)}
	}

	sort.Slice(neighbors, func(i, j int) bool {
		return neighbors[i].Distance < neighbors[j].Distance
	})

	counts := map[string]int{}
	for i := 0; i < k; i++ {
		counts[neighbors[i].Name]++
	}

	maxCount := 0
	var leaders []string
	for name, c := range counts {
		if c > maxCount {
			leaders = []string{name}
			maxCount = c
		} else if c == maxCount {
			leaders = append(leaders, name)
		}
	}

	if len(leaders) == 1 {
		return fmt.Sprintf("Your fruit with x: %.2f, y: %.2f is probably %s", x, y, leaders[0])
	}
	return "It's 50/50 bro"
}

func main() {
	fruits := []Fruit{
		{"orange", 1, 1},
		{"orange", 2, 2},
		{"grapefruit", 7, 7},
		{"grapefruit", 8, 8},
	}

	fmt.Println(FindNearestFruit(fruits, 3, 3, 3))
	fmt.Println(FindNearestFruit(fruits, 6, 6, 2))
	fmt.Println(FindNearestFruit(fruits, 5, 5, 2))
}
