# Cloud Design: Questions during Interview (dlcl)

### 1. Rate limit:

**You can set producer rate limit and consumer rate limit configurations.**
**Normally setting consumer rate limit with "back off" algorithm in case of high demand, so it can have enough time to scale out underlying VMs/ec2 instances.** Discarded messages that the consumer couldn't process int time can be sent to a DLQ (death letter queue) for a later process.
Other strategies include having a high speed queue with recent data, another lower speed queue for working with less recent data that can be handled in a separated time frame and, and finally, the DLQ for re-trying/processing messages lost in a slow time frame. 

### 2. Key/Idempotent:

Sync API call scenario:
```text
API call ---> WAF/ALB -----> Paymant APi(s) --¬
                        |--> Payment API    --|--> DB
                        |--> Payment API    --|
```	

**If the Payment API receives a call, but it breaks in the middle, so the client can receive the message and process it, but the client retries/API Call/ALB retries and then the message can be processed twice? How to "ensure" only one execution?**

**Having a unique key "idempotent" ensures that the execution is idempotent, executed once so if next call attempts to perform again the response should say "already processed".**

**Idempotency is a key aspect of API design, it is a concept that refers to the property of an operation where performing it multiple times produces the same result as performing it once.**

In a payment context, idempotency ensures that if a request to, say, "charge a customer" is accidentally sent multiple times (e.g., due to network retries), only one charge will be processed. This is often implemented using an **idempotency key**.

Imagine a payment API endpoint, `POST /payments`, that charges a customer's account. To ensure the charge happens only once, the client includes an **idempotency key** with each request. The payment service then checks this key to determine if the operation was already completed.

Here’s a basic flow:

1. **Client Request**: The client sends a `POST` request to `/payments` with an idempotency key, e.g., `X-Idempotency-Key: 12345`.
   
2. **First Request Handling**:
   - The payment service checks if the key `12345` has been seen before.
   - If it hasn’t, it processes the payment, stores the idempotency key, and associates it with the payment record.
   - The server responds with success (e.g., `200 OK` or `201 Created`), indicating the amount has been charged.

3. **Subsequent Requests with the Same Key**:
   - If a duplicate `POST /payments` request with `X-Idempotency-Key: 12345` is received (e.g., due to a network retry), the server checks the key.
   - Since the key `12345` is already associated with a completed transaction, the server simply returns the previous result without processing the payment again.

***Example in Code***

```http
POST /payments
Content-Type: application/json
X-Idempotency-Key: 12345

{
    "customer_id": "789",
    "amount": 100.0
}
```

If the `X-Idempotency-Key` (`12345`) is unique for each intended payment, any retries with the same key will have no effect beyond the initial transaction. This ensures only one charge per key, making the operation idempotent.

### 4. Circuit Breaker Pattern and Scaling Out/Scaling In according to the demand

### 3. Service Lock: Memcache/Redis (and Pool -see "4")

### 4. Go Pools:

**A goroutine pool is a pattern used to limit and manage the number of goroutines running concurrently in a Go program.** This is particularly **useful in scenarios where a "large number" of tasks need to be executed concurrently but creating too many goroutines could lead to excessive memory usage or CPU contention.**

Here’s a quick overview, its benefits and drawbacks, and how to implement a simple goroutine pool in Go.

**Why Use a Goroutine Pool?**

- **Resource Control: By limiting the number of concurrent goroutines**, you control the resource usage of your application.
- **Improved Efficiency**: Managing goroutines in a pool can help prevent overwhelming system resources, especially when dealing with high workloads.
- **Avoiding Goroutine Leaks: With "structured pooling", it’s easier to track goroutines and ensure they’re properly closed** or returned after execution.

***Implementation of a Simple Goroutine Pool***

In a **goroutine pool**, you generally create a **"fixed number" of workers (goroutines) that process tasks** from a shared queue (channel). Here’s an example:

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// Task represents a unit of work.
type Task struct {
	ID int
}

// Worker function that processes tasks from the job channel.
func worker(id int, jobs <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job.ID)
		time.Sleep(1 * time.Second) // Simulate work
	}
}

func main() {
	const numWorkers = 3
	const numJobs = 10

	// Create a channel to hold tasks and a WaitGroup for synchronization.
	jobs := make(chan Task, numJobs)
	var wg sync.WaitGroup

	// Start the worker pool.
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Send tasks to the job channel.
	for j := 1; j <= numJobs; j++ {
		jobs <- Task{ID: j}
	}

	close(jobs) // Close the channel to signal no more jobs
	wg.Wait()   // Wait for all workers to complete
	fmt.Println("All jobs completed.")
}
```

***Explanation***

1. **Worker Function**: Each worker receives tasks from the `jobs` channel and processes them. The `time.Sleep` function is used here to simulate some work.
2. **Job Queue**: `jobs` is a buffered channel holding tasks (`Task` objects).
3. **WaitGroup**: Used to wait for all workers to finish their tasks before the program exits.
4. **Task Submission**: We submit jobs to the `jobs` channel and close it after submitting all tasks.

***Advantages of a Goroutine Pool***

- **Limits Concurrency: "Prevents creating" an unbounded number of goroutines**.
- **Efficient Resource Usage**: Maintains a fixed number of workers, optimizing memory and CPU usage.
- **Simplifies Task Management**: Ensures all tasks are processed in a structured, manageable way.

***Disadvantages of a Goroutine Pool***

- **Increased Complexity**: Requires careful management of channels, synchronization, and error handling.
- **Limited Throughput**: Constraining the number of concurrent tasks might slow down task completion if the pool size is too small.
- **Static Pool Size: A fixed pool may either underutilize or overload system resources, especially if workload "size varies" widely. Dynamic pools are more complex but may be "more efficient"**.

**When to Use a Goroutine Pool**

- **CPU/Memory Intensive Tasks**: When each task uses significant system resources and you want to avoid overwhelming the system.
- **Rate-Limiting Concurrent Operations: If accessing a "rate-limited" resource (e.g., an API) or using a resource that has connection limits (e.g., database connections).**
- **"Bulk" Task Processing: Ideal for processing tasks like data processing, handling batch jobs**, or implementing worker-like behavior.

**A goroutine "pool" is NOT necessarily the same as a "buffered" channel**, but **buffered channels are often used to help implement goroutine pools**. Here’s how they relate and where they differ:

1. **Buffered Channels as Task Queues**: 
   - In a goroutine pool, a buffered channel can act as a queue to hold tasks or jobs that workers (goroutines) will process.
   - **When a task is sent to the buffered channel, if there’s space available in the buffer, it will be stored "without blocking" the Sender**. This can help smooth out bursts of tasks by temporarily holding them until a worker is ready to process them.

2. **Goroutine Pool Implementation with Buffered Channels**:
   - The goroutine pool typically has a fixed number of goroutines (or “workers”) that constantly receive tasks from a shared channel.
   - The channel may be **buffered to handle a "burst" of incoming tasks**. **This means that even if all workers are busy, the tasks can still queue up in the channel’s buffer instead of blocking the producer immediately.**

3. **Synchronization and Concurrency Control**:
   - **Buffered channels provide a basic form of backpressure by holding tasks temporarily**, so the producer can continue working even if all workers are busy.
   - If the buffer limit is reached, producers will block until a slot becomes available, providing a level of flow control.

4. **Examples of Unbuffered and Buffered Channel Usage**:
   - An **Unbuffered channel would "block the sender immediately" if no goroutine is available to pick up the task**. This might be appropriate if tasks should only be produced when a worker is immediately available.
   - A **Buffered channel lets tasks wait in a queue until a worker becomes available**, balancing between keeping workers busy and not overwhelming them with excessive tasks.

***Example Comparison of Buffered vs. Unbuffered Channels in a Goroutine Pool***

***Using Buffered Channel***

```go
const poolSize = 3
const taskCount = 10

func main() {
    jobs := make(chan int, taskCount) // Buffered to hold up to taskCount tasks
    var wg sync.WaitGroup

    // Start workers
    for w := 0; w < poolSize; w++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for job := range jobs {
                fmt.Printf("Worker %d processing job %d\n", id, job)
                time.Sleep(time.Second) // Simulated work
            }
        }(w)
    }

    // Submit jobs
    for j := 0; j < taskCount; j++ {
        jobs <- j
    }
    close(jobs)
    wg.Wait()
}
```

***Using Unbuffered Channel***

```go
jobs := make(chan int) // No buffer
var wg sync.WaitGroup

// Start workers
for w := 0; w < poolSize; w++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        for job := range jobs {
            fmt.Printf("Worker %d processing job %d\n", id, job)
            time.Sleep(time.Second) // Simulated work
        }
    }(w)
}

// Submit jobs in a goroutine to avoid deadlock
go func() {
    for j := 0; j < taskCount; j++ {
        jobs <- j
    }
    close(jobs)
}()
wg.Wait()
```

In summary, **a buffered channel is often used to implement a goroutine pool but is not itself a goroutine pool**. It simply provides a mechanism to handle the inflow of tasks and smooth out the workload distribution among workers.