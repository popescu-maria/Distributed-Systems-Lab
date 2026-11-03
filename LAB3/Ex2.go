package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
)

type LamportClock struct{ counter int }

func (l *LamportClock) Tick() { l.counter++ }

func (l *LamportClock) Send() { l.Tick() }

func (l *LamportClock) Receive(ts int) {
	if ts > l.counter {
		l.counter = ts
	}
	l.Tick()
}

type Message struct{ sender, ts int }

type Event struct {
	process          int
	internal_counter int
	kind             string
}

func less(a, b Event) bool {
	if a.internal_counter != b.internal_counter {
		return a.internal_counter < b.internal_counter
	}
	return a.process < b.process
}

func runProcess(id, rounds int, out chan<- Message, in <-chan Message,
	events chan<- Event, wg *sync.WaitGroup) {
	defer wg.Done()
	clock := &LamportClock{}

	for r := 0; r < rounds; r++ {
		n := 1 + rand.Intn(3)
		for i := 0; i < n; i++ {
			clock.Tick()
			events <- Event{id, clock.counter, "internal"}
		}

		clock.Send()
		events <- Event{id, clock.counter, "send"}
		out <- Message{id, clock.counter}

		msg := <-in
		clock.Receive(msg.ts)
		events <- Event{id, clock.counter, "recv"}
	}
}

func main() {
	const rounds = 2

	ch12 := make(chan Message, rounds)
	ch23 := make(chan Message, rounds)
	ch31 := make(chan Message, rounds)
	events := make(chan Event, 100)

	var wg sync.WaitGroup
	wg.Add(3)
	go runProcess(1, rounds, ch12, ch31, events, &wg)
	go runProcess(2, rounds, ch23, ch12, events, &wg)
	go runProcess(3, rounds, ch31, ch23, events, &wg)

	wg.Wait()
	close(events)

	var all []Event
	for e := range events {
		all = append(all, e)
	}

	sort.Slice(all, func(i, j int) bool { return less(all[i], all[j]) })

	for _, e := range all {
		fmt.Printf("C=%2d  P%d  %s\n", e.internal_counter, e.process, e.kind)
	}
}