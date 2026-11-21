package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Message struct {
	SenderID int
	Round    int
	Value    int
}

type Node struct {
	id          int
	x           int
	s           int
	n           int
	isByzantine bool
	inbox       chan Message
}

func majority(votes []int) (int, int) {
	ones := 0
	for _, v := range votes {
		if v == 1 {
			ones++
		}
	}
	zeros := len(votes) - ones

	if ones >= len(votes)/2+1 {
		return 1, ones
	}
	return 0, zeros
}

func run(nd *Node, nodes []Node, wg *sync.WaitGroup) {
	defer wg.Done()
	var stash []Message

	for t := 0; t <= nd.s+1; t++ {
		// Runda 1
		for j := range nodes {
			v := nd.x
			if nd.isByzantine {
				v = rand.Intn(2)
			}
			nodes[j].inbox <- Message{nd.id, 1, v}
		}

		votes := []int{}
		got := make([]bool, nd.n)
		tie := -1
		var nextStash []Message

		handle := func(m Message) {
			if m.Round == 2 {
				tie = m.Value
			} else if !got[m.SenderID] {
				got[m.SenderID] = true
				votes = append(votes, m.Value)
			} else {
				nextStash = append(nextStash, m)
			}
		}

		for _, m := range stash {
			handle(m)
		}
		for len(votes) < nd.n {
			handle(<-nd.inbox)
		}

		maj, mult := majority(votes)

		// Runda 2
		if nd.id == t {
			for j := range nodes {
				v := maj
				if nd.isByzantine {
					v = rand.Intn(2)
				}
				nodes[j].inbox <- Message{nd.id, 2, v}
			}
		}
		for tie == -1 {
			handle(<-nd.inbox)
		}

		if 2*mult > nd.n+2*nd.s {
			nd.x = maj
		} else {
			nd.x = tie
		}
		stash = nextStash
	}
}

func main() {
	var n, s int
	fmt.Scan(&n, &s)

	nodes := make([]Node, n)
	for i := range nodes {
		nodes[i] = Node{id: i, s: s, n: n, inbox: make(chan Message, 3*n)}
		fmt.Scan(&nodes[i].x)
	}

	perm := rand.Perm(n)
	for k := 0; k < s; k++ {
		nodes[perm[k]].isByzantine = true
	}

	var wg sync.WaitGroup
	for i := range nodes {
		wg.Add(1)
		go run(&nodes[i], nodes, &wg)
	}
	wg.Wait()

	for i := range nodes {
		if !nodes[i].isByzantine {
			fmt.Println(nodes[i].x)
			break
		}
	}
}
