# Buffered Channels

**Buffered channels in Go are a type of channel that allows a fixed number of elements to be stored in the channel buffer before blocking any further sends.** They are **useful when you want to decouple the sender and receiver goroutines**, allowing the sender to continue execution up to the buffer capacity even **if the receiver isn't immediately ready to receive.**

**Creating a Buffered Channel**

You specify the buffer capacity when creating the channel using the `make` function:
```go
ch := make(chan int, 3) // Buffered channel with a capacity of 3
```

**How Buffered Channels Work**

- **Sending to a Buffered Channel**:
  - The `send` operation (`ch <- value`) adds the value to the channel buffer if there's space.
  - If the buffer is full, the `send` operation blocks until there's space available.

- **Receiving from a Buffered Channel**:
  - The `receive` operation (`value := <-ch`) retrieves a value from the channel buffer if it's not empty.
  - If the buffer is empty, the `receive` operation blocks until a value is available.

***Example***

```go
package main

import "fmt"

func main() {
    ch := make(chan int, 2) // Buffered channel with capacity of 2

    // Send values to the channel
    ch <- 10
    ch <- 20
    fmt.Println("Sent 10 and 20 to the channel")

    // Receive values from the channel
    fmt.Println(<-ch) // Outputs: 10
    fmt.Println(<-ch) // Outputs: 20
}
```

**Key Points**

1. **Non-blocking up to Capacity**:
   Sending to a buffered channel is non-blocking only while the buffer has space.

2. **Capacity**:
   You can query the capacity of a channel using `cap(ch)` and the current number of elements in the buffer using `len(ch)`.

3. **Deadlocks**:
   If the sender keeps sending to a full channel without a receiver, or if a receiver keeps waiting on an empty channel without a sender, it will lead to a deadlock.

4. **Closing Buffered Channels**:
   When a buffered channel is closed, receivers can still retrieve values already in the buffer. Receiving from a closed and empty channel returns the zero value of the channel's type.

***Example with Capacity Query***

```go
package main

import "fmt"

func main() {
    ch := make(chan string, 3)

    ch <- "hello"
    ch <- "world"

    fmt.Println("Capacity:", cap(ch)) // Outputs: 3
    fmt.Println("Length:", len(ch))   // Outputs: 2

    fmt.Println(<-ch) // Outputs: hello
    fmt.Println(<-ch) // Outputs: world
}
``` 

**Buffered channels are especially useful in scenarios where producers and consumers operate at different rates.**