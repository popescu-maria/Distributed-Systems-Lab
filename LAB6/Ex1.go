package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func diametru(adj [][]int) int {
	n := len(adj)
	diam := 0
	for src := 0; src < n; src++ {
		dist := make([]int, n)
		for i := range dist {
			dist[i] = -1
		}
		dist[src] = 0
		queue := []int{src}
		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			for _, v := range adj[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					queue = append(queue, v)
				}
			}
		}
		for _, d := range dist {
			if d == -1 {
				return -1
			}
			if d > diam {
				diam = d
			}
		}
	}
	return diam
}

type Info struct {
	ID    int
	Value int
}

type Memory map[Info]struct{}

type Node struct {
	ID        int
	Neighbors []int
	Memory    Memory
}

func round(nodes []Node, alive []bool) {
	incoming := make([]Memory, len(nodes))
	for i := range incoming {
		incoming[i] = make(Memory)
	}

	for u := range nodes {
		if !alive[u] {
			continue
		}
		for _, v := range nodes[u].Neighbors {
			for info := range nodes[u].Memory {
				incoming[v][info] = struct{}{}
			}
		}
	}

	for v := range nodes {
		if !alive[v] {
			continue
		}
		for info := range incoming[v] {
			nodes[v].Memory[info] = struct{}{}
		}
	}
}

func main() {
	var n, s, d int
	fmt.Scan(&n, &s, &d)

	tokens := make([]int, n)
	for i := range tokens {
		fmt.Scan(&tokens[i])
	}

	adj := make([][]int, n)
	for i := range adj {
		adj[i] = make([]int, d)
		for j := range adj[i] {
			fmt.Scan(&adj[i][j])
		}
	}

	diam := diametru(adj)
	if diam == -1 {
		fmt.Println("Graful nu este tare conex")
		return
	}

	nodes := make([]Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = Node{
			ID:        i,
			Neighbors: adj[i],
			Memory:    Memory{Info{i, tokens[i]}: {}},
		}
	}

	alive := make([]bool, n)
	for i := range alive {
		alive[i] = true
	}

	crashed := 0
	t := 0
	for ; ; t++ {
		if crashed < s && rand.Intn(4) == 0 {
			v := rand.Intn(n)
			if alive[v] {
				alive[v] = false
				crashed++
			}
		}
		round(nodes, alive)
		if t > (s+1)*diam {
			break
		}
	}

	fmt.Println("Diametru:", diam)
	fmt.Println("Oprit la iteratia:", t)
	for i := range nodes {
		if !alive[i] {
			continue
		}
		var list []Info
		for info := range nodes[i].Memory {
			list = append(list, info)
		}
		sort.Slice(list, func(a, b int) bool {
			return list[a].ID < list[b].ID
		})
		for k, info := range list {
			if k > 0 {
				fmt.Print(" ")
			}
			fmt.Printf("(%d, %d)", info.ID, info.Value)
		}
		fmt.Println()
		break
	}
}
