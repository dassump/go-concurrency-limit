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
	fmt.Println("š Queue:", len(queue))
	fmt.Println("š„ Concurrency:", concurrency)
	fmt.Println("ā± Duration: up to", duration, "milliseconds")
	fmt.Println()
}
```
```shell
š Queue: 78
š„ Concurrency: 8
ā± Duration: up to 5000 milliseconds
```

## Function work
Fake long-running job.
Prints info when the job starts and ends.

Increment the counter value when done.

```go
func work(item string, duration int) {
	fmt.Println("\tš Job start:", item)
	defer fmt.Println("\tš Job end:", item)

	time.Sleep(time.Duration(duration) * time.Millisecond)
	done++
}
```
```shell
	š Job start: c3
	š Job end: c3
```

## Function main

### Status
Print the running status of the queue every one second.

```go
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Printf("š„ Concurrency: %d | Queue: %d | Done: %d š„\n", runtime.NumGoroutine()-2, len(queue), done)
		}
	}()
```

```shell
š„ Concurrency: 8 | Queue: 78 | Done: 0 š„
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
š Queue: 78
š„ Concurrency: 8
ā± Duration: up to 5000 milliseconds

	š Job start: c2
	š Job start: a2
	š Job start: a3
	š Job start: b3
	š Job start: c1
	š Job start: b1
	š Job start: b2
	š Job start: a1
š„ Concurrency: 8 | Queue: 78 | Done: 0 š„
	š Job end: c1
	š Job start: c3
	š Job end: c3
	š Job start: d1
	š Job end: a2
	š Job start: d2
š„ Concurrency: 8 | Queue: 78 | Done: 3 š„
	š Job end: b3
	š Job start: d3
	š Job end: b1
	š Job start: e1
	š Job end: d2
	š Job start: e2
	š Job end: e2
	š Job start: e3
	š Job end: c2
	š Job start: f1
š„ Concurrency: 8 | Queue: 78 | Done: 8 š„
	š Job end: a1
	š Job start: f2
š„ Concurrency: 8 | Queue: 78 | Done: 9 š„
	š Job end: a3
	š Job start: f3
	š Job end: f2
	š Job start: g1
	š Job end: b2
	š Job start: g2
š„ Concurrency: 8 | Queue: 78 | Done: 12 š„
	š Job end: d1
	š Job start: g3
	š Job end: f3
	š Job start: h1
	š Job end: g3
	š Job start: h2
	š Job end: d3
	š Job start: h3
	š Job end: e1
	š Job start: i1
	š Job end: h1
	š Job start: i2
š„ Concurrency: 8 | Queue: 78 | Done: 18 š„
	š Job end: f1
	š Job start: i3
	š Job end: h3
	š Job start: j1
š„ Concurrency: 8 | Queue: 78 | Done: 20 š„
	š Job end: h2
	š Job start: j2
	š Job end: e3
	š Job start: j3
	š Job end: g1
	š Job start: k1
š„ Concurrency: 8 | Queue: 78 | Done: 23 š„
	š Job end: k1
	š Job start: k2
	š Job end: k2
	š Job start: k3
	š Job end: g2
	š Job start: l1
	š Job end: i1
	š Job start: l2
š„ Concurrency: 8 | Queue: 78 | Done: 27 š„
	š Job end: l2
	š Job start: l3
	š Job end: i3
	š Job start: m1
	š Job end: l3
	š Job start: m2
	š Job end: j1
	š Job start: m3
	š Job end: j2
	š Job start: n1
	š Job end: n1
	š Job start: n2
š„ Concurrency: 8 | Queue: 78 | Done: 33 š„
	š Job end: m2
	š Job start: n3
	š Job end: l1
	š Job start: o1
	š Job end: j3
	š Job start: o2
	š Job end: o1
	š Job start: o3
	š Job end: k3
	š Job start: p1
	š Job end: i2
	š Job start: p2
	š Job end: o2
	š Job start: p3
š„ Concurrency: 8 | Queue: 78 | Done: 40 š„
	š Job end: m1
	š Job start: q1
	š Job end: m3
	š Job start: q2
	š Job end: n2
	š Job start: q3
š„ Concurrency: 8 | Queue: 78 | Done: 43 š„
	š Job end: q1
	š Job start: r1
	š Job end: q2
	š Job start: r2
	š Job end: o3
	š Job start: r3
š„ Concurrency: 8 | Queue: 78 | Done: 46 š„
	š Job end: n3
	š Job start: s1
	š Job end: q3
	š Job start: s2
š„ Concurrency: 8 | Queue: 78 | Done: 48 š„
	š Job end: r2
	š Job start: s3
	š Job end: p2
	š Job start: t1
š„ Concurrency: 8 | Queue: 78 | Done: 50 š„
	š Job end: p1
	š Job start: t2
	š Job end: r3
	š Job start: t3
	š Job end: p3
	š Job start: u1
š„ Concurrency: 8 | Queue: 78 | Done: 53 š„
	š Job end: r1
	š Job start: u2
	š Job end: s1
	š Job start: u3
	š Job end: s2
	š Job start: v1
š„ Concurrency: 8 | Queue: 78 | Done: 56 š„
	š Job end: v1
	š Job start: v2
	š Job end: s3
	š Job start: v3
	š Job end: u1
	š Job start: w1
š„ Concurrency: 8 | Queue: 78 | Done: 59 š„
	š Job end: w1
	š Job start: w2
	š Job end: u3
	š Job start: w3
š„ Concurrency: 8 | Queue: 78 | Done: 61 š„
	š Job end: t1
	š Job start: x1
	š Job end: t3
	š Job start: x2
	š Job end: u2
	š Job start: x3
	š Job end: t2
	š Job start: y1
š„ Concurrency: 8 | Queue: 78 | Done: 65 š„
	š Job end: x1
	š Job start: y2
	š Job end: v2
	š Job start: y3
	š Job end: w2
	š Job start: z1
	š Job end: z1
	š Job start: z2
	š Job end: y3
	š Job start: z3
š„ Concurrency: 8 | Queue: 78 | Done: 70 š„
	š Job end: x2
š„ Concurrency: 7 | Queue: 78 | Done: 71 š„
	š Job end: z2
	š Job end: v3
š„ Concurrency: 5 | Queue: 78 | Done: 73 š„
	š Job end: z3
	š Job end: w3
	š Job end: y1
	š Job end: x3
š„ Concurrency: 1 | Queue: 78 | Done: 77 š„
	š Job end: y2
```
