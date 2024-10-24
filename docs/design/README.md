# GoF Design Patterns

Here’s a brief overview of the **Gang of Four (GoF) Design Patterns** categorized by **Creational**, **Structural**, and **Behavioral** in a table format:

| **Category**      | **Pattern**       | **Brief Description**                                                                 |
|-------------------|-------------------|---------------------------------------------------------------------------------------|
| **Creational**     | **Factory Method**| Creates objects via a method, deferring instantiation to subclasses.                  |
|                   | **Abstract Factory**| Provides an interface for creating families of related objects without specifying their concrete classes. |
|                   | **Builder**        | Separates the construction of complex objects from their representation.              |
|                   | **Prototype**      | Creates new objects by copying an existing object (prototype).                        |
|                   | **Singleton**      | Ensures a class has only one instance and provides a global point of access to it.    |
| **Structural**     | **Adapter**        | Converts one interface to another that a client expects, allowing incompatible interfaces to work together. |
|                   | **Bridge**         | Decouples an abstraction from its implementation, allowing both to vary independently.|
|                   | **Composite**      | Composes objects into tree-like structures to represent part-whole hierarchies.       |
|                   | **Decorator**      | Adds behavior to objects dynamically without affecting other objects of the same class.|
|                   | **Facade**         | Provides a simplified interface to a complex subsystem.                              |
|                   | **Flyweight**      | Shares objects to support large numbers of fine-grained objects efficiently.          |
|                   | **Proxy**          | Provides a surrogate or placeholder to control access to an object.                   |
| **Behavioral**     | **Chain of Responsibility** | Passes requests along a chain of handlers, each having the chance to process the request.|
|                   | **Command**        | Encapsulates a request as an object, allowing parameterization and queuing of requests.|
|                   | **Interpreter**    | Defines a representation for a language’s grammar and an interpreter to handle parsing.|
|                   | **Iterator**       | Provides a way to access elements of a collection sequentially without exposing its underlying representation. |
|                   | **Mediator**       | Facilitates communication between objects without them being directly dependent on each other. |
|                   | **Memento**        | Captures and restores an object’s internal state without violating encapsulation.     |
|                   | **Observer**       | Defines a dependency between objects so that when one changes state, others are notified automatically. |
|                   | **State**          | Allows an object to alter its behavior when its internal state changes.               |
|                   | **Strategy**       | Defines a family of algorithms and makes them interchangeable.                       |
|                   | **Template Method**| Defines the skeleton of an algorithm, with steps deferred to subclasses.              |
|                   | **Visitor**        | Adds new operations to objects without modifying their structure.                    |



---
# Common Design Patterns in GO

In Go (Golang), common design patterns focus on simplicity and leveraging Go's features like concurrency, interfaces, and idiomatic constructs. Here's a brief overview of some common Go patterns:

| **Pattern**             | **Description**                                                                 |
|-------------------------|---------------------------------------------------------------------------------|
| **Singleton**           | Ensures a type has only one instance with global access **(typically through `sync.Once`)**. |
| **Factory**             | A function that returns instances of interfaces or types without exposing their concrete implementation. |
| **Builder**             | Used to incrementally build complex objects, often with chaining methods.        |
| **Adapter**             | Converts an interface into another, allowing compatibility between otherwise incompatible types. |
| **Decorator**           | Dynamically adds functionality to objects by wrapping them. Often done by embedding types. |
| **Proxy**               | Acts as a placeholder or mediator to control access to an object, adding lazy initialization or access control. |
| **Observer**            | Implements **pub-sub (publish-subscribe) mechanisms where objects notify subscribers of events (commonly done via channels).** |
| **Strategy**            | Defines a family of algorithms that are interchangeable, often by using function types or interfaces. |
| **Command**             | Encapsulates a request as an object, allowing delayed or conditional execution (can be implemented using function types). |
| **Chain of Responsibility** | Passes a request along a chain of handlers until one handles it, commonly used with middleware in web frameworks like Gin. |
| **State**               | Manages state-specific behavior in objects by delegating to different methods or types based on state. |
| **Template Method**     | Defines the steps of an algorithm, with some steps implemented in the base and others delegated to extensions. |
| **Repository**          | Provides a clean API to access data sources like databases, abstracting the persistence logic. |
| **Concurrency Patterns**| Patterns leveraging Go's goroutines and channels for concurrent execution:                                        |
| - **Worker Pool**       | Spawns a fixed number of workers (goroutines) to process tasks from a channel.                                       |
| - **Pipeline**          | Chains stages of processing where the output of one stage becomes the input of another, using channels to pass data. |
| - **Fan-out/Fan-in**    | A concurrency pattern where multiple goroutines perform tasks in parallel (fan-out), and their results are aggregated (fan-in). |

These patterns help Go developers write clean, maintainable, and scalable code while taking full advantage of Go's strengths like goroutines, channels, and interfaces.

In Go, the most recognized and commonly used design patterns often revolve around simplicity, concurrency, and idiomatic use of Go's features like interfaces and channels. Here's a list of patterns frequently seen in Go codebases:

## Examples

### 1. **Singleton**
   - **Description:** Ensures a type has only one instance, often used for shared resources like database connections. In Go, this is implemented using `sync.Once` to ensure thread-safe initialization.
   - **Example:**
     ```go
     var once sync.Once
     var instance *Singleton

     func GetInstance() *Singleton {
         once.Do(func() {
             instance = &Singleton{}
         })
         return instance
     }
     ```

### 2. **Factory**
   - **Description:** Provides a way to create objects without exposing the creation logic. It's typically used with interfaces.
   - **Example:**
     ```go
     type Animal interface {
         Speak() string
     }

     type Dog struct{}
     func (d Dog) Speak() string { return "Woof" }

     type Cat struct{}
     func (c Cat) Speak() string { return "Meow" }

     func AnimalFactory(a string) Animal {
         if a == "dog" {
             return Dog{}
         }
         return Cat{}
     }
     ```

### 3. **Adapter**
   - **Description:** Converts one interface into another, allowing two incompatible interfaces to work together. Useful when integrating third-party packages.
   - **Example:**
     ```go
     type LegacyPrinter interface {
         Print(s string) string
     }

     type ModernPrinter struct{}
     func (mp *ModernPrinter) PrintFormatted(s string) string {
         return "[Formatted]" + s
     }

     type PrinterAdapter struct {
         ModernPrinter *ModernPrinter
     }

     func (pa *PrinterAdapter) Print(s string) string {
         return pa.ModernPrinter.PrintFormatted(s)
     }
     ```

### 4. **Decorator**
   - **Description:** Adds new functionality to an existing object without modifying its structure. In Go, this can be done by wrapping functions or embedding structs.
   - **Example:**
     ```go
     type Notifier interface {
         Send(message string)
     }

     type EmailNotifier struct{}
     func (e EmailNotifier) Send(message string) {
         fmt.Println("Email:", message)
     }

     type SlackNotifier struct {
         Notifier Notifier
     }
     func (s SlackNotifier) Send(message string) {
         s.Notifier.Send(message)
         fmt.Println("Slack:", message)
     }

     email := EmailNotifier{}
     notifier := SlackNotifier{Notifier: email}
     notifier.Send("Hello World")
     ```

### 5. **Strategy**
   - **Description:** Defines a family of interchangeable algorithms. Go uses function types or interfaces for this pattern.
   - **Example:**
     ```go
     type PaymentStrategy interface {
         Pay(amount int) string
     }

     type CreditCard struct{}
     func (c CreditCard) Pay(amount int) string {
         return fmt.Sprintf("Paid %d with Credit Card", amount)
     }

     type PayPal struct{}
     func (p PayPal) Pay(amount int) string {
         return fmt.Sprintf("Paid %d with PayPal", amount)
     }

     func Checkout(strategy PaymentStrategy, amount int) {
         fmt.Println(strategy.Pay(amount))
     }
     ```

### 6. **Repository**
   - **Description:** Abstracts data storage access, making code cleaner and more maintainable by hiding the persistence logic. This is often used in conjunction with interfaces.
   - **Example:**
     ```go
     type User struct {
         ID   int
         Name string
     }

     type UserRepository interface {
         GetUser(id int) (*User, error)
         SaveUser(user *User) error
     }

     type InMemoryUserRepo struct {
         users map[int]*User
     }

     func (r *InMemoryUserRepo) GetUser(id int) (*User, error) {
         return r.users[id], nil
     }
     func (r *InMemoryUserRepo) SaveUser(user *User) error {
         r.users[user.ID] = user
         return nil
     }
     ```

### 7. **Command**
   - **Description:** Encapsulates a request as an object, useful for handling operations such as queuing, logging, or undoable actions. In Go, this is often done via function types.
   - **Example:**
     ```go
     type Command func()

     func ExecuteCommand(cmd Command) {
         cmd()
     }

     func main() {
         command := func() { fmt.Println("Command Executed") }
         ExecuteCommand(command)
     }
     ```

### 8. **Observer**
   - **Description:** Implements the publish-subscribe pattern, where one object (subject) notifies others (observers) of state changes. In Go, this can be done using channels.
   - **Example:**
     ```go
     type Observer interface {
         Update(string)
     }

     type Subject struct {
         observers []Observer
     }

     func (s *Subject) Register(observer Observer) {
         s.observers = append(s.observers, observer)
     }

     func (s *Subject) Notify(message string) {
         for _, observer := range s.observers {
             observer.Update(message)
         }
     }

     type ConcreteObserver struct{}
     func (co *ConcreteObserver) Update(message string) {
         fmt.Println("Received:", message)
     }

     func main() {
         subject := &Subject{}
         observer := &ConcreteObserver{}

         subject.Register(observer)
         subject.Notify("Hello Observers!")
     }
     ```

### 9. **Pipeline (Concurrency Pattern)**
   - **Description:** Breaks down a process into stages where the output of one stage is the input to the next, passing data through channels.
   - **Example:**
     ```go
     func stage1(in <-chan int) <-chan int {
         out := make(chan int)
         go func() {
             for v := range in {
                 out <- v * 2
             }
             close(out)
         }()
         return out
     }

     func stage2(in <-chan int) <-chan int {
         out := make(chan int)
         go func() {
             for v := range in {
                 out <- v + 1
             }
             close(out)
         }()
         return out
     }

     func main() {
         numbers := make(chan int)
         go func() {
             for i := 0; i < 5; i++ {
                 numbers <- i
             }
             close(numbers)
         }()

         result := stage2(stage1(numbers))

         for v := range result {
             fmt.Println(v)
         }
     }
     ```

### 10. **Worker Pool (Concurrency Pattern)**
   - **Description:** Distributes work across a fixed number of goroutines (workers) that process tasks from a shared queue.
   - **Example:**
     ```go
     func worker(id int, jobs <-chan int, results chan<- int) {
         for job := range jobs {
             fmt.Printf("Worker %d started job %d\n", id, job)
             time.Sleep(time.Second) // Simulate work
             results <- job * 2
         }
     }

     func main() {
         jobs := make(chan int, 5)
         results := make(chan int, 5)

         for w := 1; w <= 3; w++ {
             go worker(w, jobs, results)
         }

         for j := 1; j <= 5; j++ {
             jobs <- j
         }
         close(jobs)

         for r := 1; r <= 5; r++ {
             fmt.Println("Result:", <-results)
         }
     }
     ```

These patterns are commonly used in Go projects to solve problems efficiently, while keeping the code simple and idiomatic.