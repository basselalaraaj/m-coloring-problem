package main

import "fmt"

func main() {
	fmt.Println("run the command `go test`")
}

func graphColoringBFS(graph [][]int, n int, m int) []int {
	colors := make([]int, n)
	queue := make([]int, 1)
	visited := make(map[int]struct{})

	for i := range colors {
		colors[i] = 1
	}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		_, ok := visited[item]
		if ok {
			continue
		}

		visited[item] = struct{}{}
		graphItem := graph[item]

		for i, c := range graphItem {
			if item == i || c == 0 {
				continue
			}

			if colors[item] == colors[i] {
				colors[i] += 1
			}

			queue = append(queue, i)

			if colors[i] > m {
				return []int{}
			}
		}
	}

	return colors
}
