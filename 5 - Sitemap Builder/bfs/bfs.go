package main

import "fmt"

func bfs(url string, maxDepth int) Set {
	depth := 0
	visited := NewSet()
	queue := NewSet()
	queue.Push(url)

	for len(queue) > 0 {
		curr := queue.Pop()
		if visited.Contains(curr) {
			continue
		}
		visited.Push(curr)

		for _, link := range parseUrl(curr) {
			queue.Push(link)
		}

		depth++

		fmt.Println("queue: ", queue)
		fmt.Println("visited: ", visited)
	}

	return visited
}

type Set []string

func NewSet() Set {
	return make([]string, 0, 10)
}

func (s *Set) Push(value string) {
	*s = append(*s, value)
}

func (s *Set) Pop() string {
	if len(*s) > 0 {
		val := (*s)[0]
		*s = (*s)[1:]

		return val
	}

	return ""
}

func (s *Set) Contains(value string) bool {
	for i := range *s {
		if (*s)[i] == value {
			return true
		}
	}

	return false
}

func (s *Set) Extend(a Set) {

}
