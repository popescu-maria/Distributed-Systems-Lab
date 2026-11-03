package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Clock struct{ offset time.Duration }

func (c *Clock) Now() time.Time { return time.Now().Add(c.offset) }

func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	theta = (t2.Sub(t1) + t3.Sub(t4)) / 2
	delta = t4.Sub(t1) - t3.Sub(t2)
	return
}

func measure(c_clock, s_clock *Clock) (theta, delta time.Duration) {
	t1 := c_clock.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t2 := s_clock.Now()

	time.Sleep(5*time.Millisecond + time.Duration(rand.Intn(20))*time.Millisecond)
	t3 := s_clock.Now()

	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t4 := c_clock.Now()
	return estimate(t1, t2, t3, t4)
}

func main() {
	server := &Clock{offset: 0}
	client := &Clock{offset: -2 * time.Second}
	trueOffset := server.offset - client.offset

	// (a)
	theta, delta := measure(client, server)
	fmt.Println("(a) theta:", theta, "delta:", delta,
		"error:", trueOffset-theta)

	// (b) 
	minDelta := time.Hour
	minTheta := time.Hour
	var thetaAtMinDelta time.Duration

	for i := 1; i <= 10; i++ {
		theta, delta := measure(client, server)
		fmt.Printf("  #%d theta: %v delta: %v error: %v\n",
			i, theta, delta, trueOffset-theta)

		if delta < minDelta {
			minDelta = delta
			thetaAtMinDelta = theta
		}
		if theta < minTheta {
			minTheta = theta
		}
	}

	fmt.Println("(b) min delta:", minDelta)
	fmt.Println("    theta at min delta:", thetaAtMinDelta,
		"error:", trueOffset-thetaAtMinDelta)
	fmt.Println("    min theta (literal):", minTheta,
		"error:", trueOffset-minTheta)
}