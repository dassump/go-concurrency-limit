package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var (
	concurrency = runtime.NumCPU()
	duration    = 5000
	queue       = []string{"a1", "a2", "a3", "b1", "b2", "b3", "c1", "c2", "c3", "d1", "d2", "d3", "e1", "e2", "e3", "f1", "f2", "f3", "g1", "g2", "g3", "h1", "h2", "h3", "i1", "i2", "i3", "j1", "j2", "j3", "k1", "k2", "k3", "l1", "l2", "l3", "m1", "m2", "m3", "n1", "n2", "n3", "o1", "o2", "o3", "p1", "p2", "p3", "q1", "q2", "q3", "r1", "r2", "r3", "s1", "s2", "s3", "t1", "t2", "t3", "u1", "u2", "u3", "v1", "v2", "v3", "w1", "w2", "w3", "x1", "x2", "x3", "y1", "y2", "y3", "z1", "z2", "z3"}
	done        = 0
)

func init() {
	fmt.Println()
	fmt.Println("📋 Queue:", len(queue))
	fmt.Println("🔥 Concurrency:", concurrency)
	fmt.Println("⏱ Duration: up to", duration, "milliseconds")
	fmt.Println()
}

func work(item string, duration int) {
	fmt.Println("\t🚀 Job start:", item)
	defer fmt.Println("\t🛑 Job end:", item)

	time.Sleep(time.Duration(duration) * time.Millisecond)
	done++
}

func main() {
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Printf("🔥 Concurrency: %d | Queue: %d | Done: %d 🔥\n", runtime.NumGoroutine()-2, len(queue), done)
		}
	}()

	var channel = make(chan bool, concurrency)
	var wait sync.WaitGroup

	for key, item := range queue {
		channel <- true
		wait.Add(1)

		go func(key int, item string) {
			defer wait.Done()

			work(item, rand.Intn(duration))

			<-channel
		}(key, item)
	}

	close(channel)
	wait.Wait()
}
