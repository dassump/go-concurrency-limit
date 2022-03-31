# go-concurrency-limit

## Variables

```go
var (
	concurrency = runtime.NumCPU()
	duration    = 5000
	queue       = []string{"a1", "a2", "a3", "b1", "b2", "b3", "c1", "c2", "c3", "d1", "d2", "d3", "e1", "e2", "e3", "f1", "f2", "f3", "g1", "g2", "g3", "h1", "h2", "h3", "i1", "i2", "i3", "j1", "j2", "j3", "k1", "k2", "k3", "l1", "l2", "l3", "m1", "m2", "m3", "n1", "n2", "n3", "o1", "o2", "o3", "p1", "p2", "p3", "q1", "q2", "q3", "r1", "r2", "r3", "s1", "s2", "s3", "t1", "t2", "t3", "u1", "u2", "u3", "v1", "v2", "v3", "w1", "w2", "w3", "x1", "x2", "x3", "y1", "y2", "y3", "z1", "z2", "z3"}
	done        = 0
)
```

### concurrency
Limit the maximum number of running goroutines.
Default value is the number of CPUs.

### duration
Sleep time used in fake long-running job.
Up to 5000 milliseconds.

### queue
Fake work queue, used for illustration only.

### done
Number of completed jobs in the queue.


## Function init
Print information banner.

```go
func init() {
	fmt.Println()
	fmt.Println("📋 Queue:", len(queue))
	fmt.Println("🔥 Concurrency:", concurrency)
	fmt.Println("⏱ Duration: up to", duration, "milliseconds")
	fmt.Println()
}
```
```shell
📋 Queue: 78
🔥 Concurrency: 8
⏱ Duration: up to 5000 milliseconds
```

## Function work
Fake long-running job.
Prints info when the job starts and ends.

Increment the counter value when done.

```go
func work(item string, duration int) {
	fmt.Println("\t🚀 Job start:", item)
	defer fmt.Println("\t🛑 Job end:", item)

	time.Sleep(time.Duration(duration) * time.Millisecond)
	done++
}
```
```shell
	🚀 Job start: c3
	🛑 Job end: c3
```

## Function main

### Status
Print the running status of the queue every one second.

```go
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Printf("🔥 Concurrency: %d | Queue: %d | Done: %d 🔥\n", runtime.NumGoroutine()-2, len(queue), done)
		}
	}()
```

```shell
🔥 Concurrency: 8 | Queue: 78 | Done: 0 🔥
```

### Golang magic
* The buffered channel is used to control the number of running goroutines.
  When the limit is reached, it becomes a blocking operation until some buffer space is freed. 
* WaitGroup is used to wait for all items in the queue to complete.
* Goroutine is called to run concurrently queued items. When the job is complete, the WaitGroup and a space in the channel's buffer is freed.

```go
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
```

### Sample output
```shell
📋 Queue: 78
🔥 Concurrency: 8
⏱ Duration: up to 5000 milliseconds

	🚀 Job start: c2
	🚀 Job start: a2
	🚀 Job start: a3
	🚀 Job start: b3
	🚀 Job start: c1
	🚀 Job start: b1
	🚀 Job start: b2
	🚀 Job start: a1
🔥 Concurrency: 8 | Queue: 78 | Done: 0 🔥
	🛑 Job end: c1
	🚀 Job start: c3
	🛑 Job end: c3
	🚀 Job start: d1
	🛑 Job end: a2
	🚀 Job start: d2
🔥 Concurrency: 8 | Queue: 78 | Done: 3 🔥
	🛑 Job end: b3
	🚀 Job start: d3
	🛑 Job end: b1
	🚀 Job start: e1
	🛑 Job end: d2
	🚀 Job start: e2
	🛑 Job end: e2
	🚀 Job start: e3
	🛑 Job end: c2
	🚀 Job start: f1
🔥 Concurrency: 8 | Queue: 78 | Done: 8 🔥
	🛑 Job end: a1
	🚀 Job start: f2
🔥 Concurrency: 8 | Queue: 78 | Done: 9 🔥
	🛑 Job end: a3
	🚀 Job start: f3
	🛑 Job end: f2
	🚀 Job start: g1
	🛑 Job end: b2
	🚀 Job start: g2
🔥 Concurrency: 8 | Queue: 78 | Done: 12 🔥
	🛑 Job end: d1
	🚀 Job start: g3
	🛑 Job end: f3
	🚀 Job start: h1
	🛑 Job end: g3
	🚀 Job start: h2
	🛑 Job end: d3
	🚀 Job start: h3
	🛑 Job end: e1
	🚀 Job start: i1
	🛑 Job end: h1
	🚀 Job start: i2
🔥 Concurrency: 8 | Queue: 78 | Done: 18 🔥
	🛑 Job end: f1
	🚀 Job start: i3
	🛑 Job end: h3
	🚀 Job start: j1
🔥 Concurrency: 8 | Queue: 78 | Done: 20 🔥
	🛑 Job end: h2
	🚀 Job start: j2
	🛑 Job end: e3
	🚀 Job start: j3
	🛑 Job end: g1
	🚀 Job start: k1
🔥 Concurrency: 8 | Queue: 78 | Done: 23 🔥
	🛑 Job end: k1
	🚀 Job start: k2
	🛑 Job end: k2
	🚀 Job start: k3
	🛑 Job end: g2
	🚀 Job start: l1
	🛑 Job end: i1
	🚀 Job start: l2
🔥 Concurrency: 8 | Queue: 78 | Done: 27 🔥
	🛑 Job end: l2
	🚀 Job start: l3
	🛑 Job end: i3
	🚀 Job start: m1
	🛑 Job end: l3
	🚀 Job start: m2
	🛑 Job end: j1
	🚀 Job start: m3
	🛑 Job end: j2
	🚀 Job start: n1
	🛑 Job end: n1
	🚀 Job start: n2
🔥 Concurrency: 8 | Queue: 78 | Done: 33 🔥
	🛑 Job end: m2
	🚀 Job start: n3
	🛑 Job end: l1
	🚀 Job start: o1
	🛑 Job end: j3
	🚀 Job start: o2
	🛑 Job end: o1
	🚀 Job start: o3
	🛑 Job end: k3
	🚀 Job start: p1
	🛑 Job end: i2
	🚀 Job start: p2
	🛑 Job end: o2
	🚀 Job start: p3
🔥 Concurrency: 8 | Queue: 78 | Done: 40 🔥
	🛑 Job end: m1
	🚀 Job start: q1
	🛑 Job end: m3
	🚀 Job start: q2
	🛑 Job end: n2
	🚀 Job start: q3
🔥 Concurrency: 8 | Queue: 78 | Done: 43 🔥
	🛑 Job end: q1
	🚀 Job start: r1
	🛑 Job end: q2
	🚀 Job start: r2
	🛑 Job end: o3
	🚀 Job start: r3
🔥 Concurrency: 8 | Queue: 78 | Done: 46 🔥
	🛑 Job end: n3
	🚀 Job start: s1
	🛑 Job end: q3
	🚀 Job start: s2
🔥 Concurrency: 8 | Queue: 78 | Done: 48 🔥
	🛑 Job end: r2
	🚀 Job start: s3
	🛑 Job end: p2
	🚀 Job start: t1
🔥 Concurrency: 8 | Queue: 78 | Done: 50 🔥
	🛑 Job end: p1
	🚀 Job start: t2
	🛑 Job end: r3
	🚀 Job start: t3
	🛑 Job end: p3
	🚀 Job start: u1
🔥 Concurrency: 8 | Queue: 78 | Done: 53 🔥
	🛑 Job end: r1
	🚀 Job start: u2
	🛑 Job end: s1
	🚀 Job start: u3
	🛑 Job end: s2
	🚀 Job start: v1
🔥 Concurrency: 8 | Queue: 78 | Done: 56 🔥
	🛑 Job end: v1
	🚀 Job start: v2
	🛑 Job end: s3
	🚀 Job start: v3
	🛑 Job end: u1
	🚀 Job start: w1
🔥 Concurrency: 8 | Queue: 78 | Done: 59 🔥
	🛑 Job end: w1
	🚀 Job start: w2
	🛑 Job end: u3
	🚀 Job start: w3
🔥 Concurrency: 8 | Queue: 78 | Done: 61 🔥
	🛑 Job end: t1
	🚀 Job start: x1
	🛑 Job end: t3
	🚀 Job start: x2
	🛑 Job end: u2
	🚀 Job start: x3
	🛑 Job end: t2
	🚀 Job start: y1
🔥 Concurrency: 8 | Queue: 78 | Done: 65 🔥
	🛑 Job end: x1
	🚀 Job start: y2
	🛑 Job end: v2
	🚀 Job start: y3
	🛑 Job end: w2
	🚀 Job start: z1
	🛑 Job end: z1
	🚀 Job start: z2
	🛑 Job end: y3
	🚀 Job start: z3
🔥 Concurrency: 8 | Queue: 78 | Done: 70 🔥
	🛑 Job end: x2
🔥 Concurrency: 7 | Queue: 78 | Done: 71 🔥
	🛑 Job end: z2
	🛑 Job end: v3
🔥 Concurrency: 5 | Queue: 78 | Done: 73 🔥
	🛑 Job end: z3
	🛑 Job end: w3
	🛑 Job end: y1
	🛑 Job end: x3
🔥 Concurrency: 1 | Queue: 78 | Done: 77 🔥
	🛑 Job end: y2
```
