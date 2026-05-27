# Go By Example - Complete Learning Repository 🚀

This repository contains practical Go programs implemented by following the concepts from **Go by Example**.  
Each topic includes hands-on examples, syntax practice, and interview-oriented explanations.

Reference Topics Covered from PDF: :contentReference[oaicite:0]{index=0}

---

# 📚 Topics Covered

---

# 1️⃣ Go Fundamentals

## Hello World
Basic Go program structure.

### Key Concepts
- `package main`
- `func main()`
- `fmt.Println()`

### Interview Point
- Entry point of Go program is `main()`.

---

## Values
Understanding:
- strings
- integers
- booleans
- arithmetic operations

### Interview Point
Go is a statically typed language.

---

## Variables
Variable declaration using:
```go
var name string
age := 25
```

### Interview Point
- `:=` works only inside functions.
- Go automatically infers datatype.

---

## Constants
Immutable values using `const`.

### Interview Point
Constants cannot be changed after declaration.

---

## For Loop
Go supports only one loop:
```go
for
```

### Interview Point
Go removed complexity of multiple loop types.

---

## If / Else
Conditional execution.

### Interview Point
Go does not use brackets `()` in conditions.

---

## Switch
Cleaner multiple condition handling.

### Interview Point
Go switch automatically breaks.

---

# 2️⃣ Collections

## Arrays
Fixed-size collection.

### Interview Point
Arrays are value types.

---

## Slices
Dynamic arrays.

### Key Concepts
- append()
- slicing
- make()

### Interview Point
Slices are reference-based and used more than arrays.

---

## Maps
Key-value data structure.

### Interview Point
Maps are unordered collections.

---

## Range
Iteration over:
- arrays
- slices
- maps
- strings

### Interview Point
`range` returns:
- index/key
- value

---

# 3️⃣ Functions

## Functions
Reusable logic blocks.

---

## Multiple Return Values

```go
func vals() (int, int)
```

### Interview Point
Go commonly returns:
```go
value, error
```

---

## Variadic Functions

```go
func sum(nums ...int)
```

### Interview Point
Accepts variable number of arguments.

---

## Closures
Functions inside functions.

### Interview Point
Closures remember outer variable values.

---

## Recursion
Function calling itself.

### Interview Point
Used in:
- tree traversal
- factorial
- DFS algorithms

---

# 4️⃣ Memory & Types

## Pointers

### Key Concepts
- `&` → address
- `*` → dereference

### Interview Point
Go has pointers but no pointer arithmetic.

---

## Strings and Runes

### Interview Point
Strings are UTF-8 encoded.

Rune represents Unicode character.

---

# 5️⃣ Structs & OOP

## Structs
Custom data types.

### Interview Point
Structs are similar to classes but without inheritance.

---

## Methods
Functions attached to structs.

### Interview Point
Methods improve code organization.

---

## Interfaces
Defines behavior.

### Interview Point
Implicit implementation is supported in Go.

---

## Enums
Created using:
```go
iota
```

### Interview Point
Go does not have native enum keyword.

---

## Struct Embedding
Alternative to inheritance.

### Interview Point
Go prefers composition over inheritance.

---

## Generics
Reusable type-safe functions.

### Interview Point
Introduced in Go 1.18.

---

# 6️⃣ Error Handling

## Errors
Explicit error handling.

### Interview Point
Go avoids try-catch mechanism.

---

## Custom Errors

```go
type MyError struct{}
```

### Interview Point
Custom errors implement:
```go
Error() string
```

---

## Panic
Stops execution.

### Interview Point
Used only for critical failures.

---

## Recover
Recovers from panic.

### Interview Point
Usually used with `defer`.

---

## Defer
Executes function before return.

### Common Uses
- file close
- DB connection close

---

# 7️⃣ Concurrency

## Goroutines
Lightweight threads.

### Interview Point
Managed by Go runtime.

---

## Channels
Communication between goroutines.

### Interview Point
Channels prevent shared memory issues.

---

## Buffered Channels
Channels with capacity.

---

## Channel Synchronization
Used for goroutine coordination.

---

## Select
Waits on multiple channels.

### Interview Point
Used in concurrent systems.

---

## Timeouts
Prevent blocking operations.

---

## Non-Blocking Channels
Avoid waiting indefinitely.

---

## Closing Channels
Signal completion.

---

## Worker Pools
Concurrent job processing.

### Interview Point
Frequently asked in backend interviews.

---

## WaitGroups
Wait for goroutines.

### Interview Point
From:
```go
sync.WaitGroup
```

---

## Mutexes
Prevent race conditions.

### Interview Point
Protect shared resources.

---

## Atomic Counters
Thread-safe counters.

### Interview Point
Faster than mutex for simple operations.

---

# 8️⃣ Sorting & Utilities

## Sorting
Sort slices.

---

## Sorting by Functions
Custom sorting logic.

### Interview Point
Implemented using:
```go
sort.Slice()
```

---

# 9️⃣ String Processing

## String Functions
Using:
```go
strings package
```

---

## String Formatting
Using:
```go
fmt.Sprintf()
```

### Interview Point
`Printf` formatting commonly asked.

---

## Regular Expressions

### Package
```go
regexp
```

---

# 🔟 JSON & XML

## JSON

### Key Concepts
- Marshal
- Unmarshal

### Interview Point
Most important for APIs.

---

## XML

### Interview Point
Uses struct tags.

---

# 1️⃣1️⃣ Time & Parsing

## Time
Using:
```go
time package
```

---

## Epoch
Unix timestamps.

---

## Time Formatting / Parsing

### Interview Point
Go uses reference time:
```go
2006-01-02
```

---

# 1️⃣2️⃣ Security & Encoding

## SHA256 Hashes

### Interview Point
Used for hashing passwords/data.

---

## Base64 Encoding

### Interview Point
Used in APIs and tokens.

---

# 1️⃣3️⃣ File Handling

## Reading Files

```go
os.ReadFile()
```

---

## Writing Files

```go
os.WriteFile()
```

---

## Directories
Directory management.

---

## Temporary Files
Used in testing and processing.

---

# 1️⃣4️⃣ Testing

## Testing

```bash
go test
```

### Interview Point
Test files:
```text
*_test.go
```

---

## Benchmarking

```bash
go test -bench=.
```

### Interview Point
Measures performance.

---

# 1️⃣5️⃣ Command Line

## Command-Line Arguments

```go
os.Args
```

---

## Command-Line Flags

```go
flag package
```

---

## Environment Variables

```go
os.Getenv()
```

---

# 1️⃣6️⃣ Networking

## HTTP Client
Calling APIs.

---

## HTTP Server
Building APIs.

### Interview Point
Most important for backend roles.

---

## TCP Server
Socket programming.

---

## Context
Request lifecycle management.

### Interview Point
Used for:
- timeout
- cancellation
- tracing

---

# 1️⃣7️⃣ System Programming

## Spawning Processes

```go
os/exec
```

---

## Signals
Handle OS signals.

---

## Exit
Terminate program.

---

# 🛠 Technologies Used

- Go (Golang)
- Go Standard Library
- VS Code
- Git & GitHub

---

# 📂 Repository Structure

```text
Go-Learning/
│
├── 01-Hello-World
├── 02-Values
├── 03-Variables
├── ...
├── 69-Testing-and-Benchmarking
│
├── go.mod
├── go.sum
└── README.md
```

---

# ▶️ Run Examples

Run program:

```bash
go run main.go
```

Run tests:

```bash
go test -v
```

Run benchmarks:

```bash
go test -bench=.
```

---

# 🎯 Goal

- Learn Go from Basic to Advanced
- Build Backend Skills
- Learn Concurrency
- Prepare for Go Interviews
- Build Production-Level Applications

---
# 🎯 Go Interview Questions & Answers

---

# 1️⃣ What is Go Language?

## Answer:
Go (Golang) is an open-source programming language developed by Google.

It is designed for:
- simplicity
- concurrency
- high performance
- scalable backend systems

Go is widely used in:
- Backend APIs
- Microservices
- Cloud Computing
- DevOps tools

---

# 2️⃣ Why Go is Faster?

## Answer:
Go is faster because:
- Compiled language
- Lightweight goroutines
- Efficient memory management
- Built-in concurrency support

---

# 3️⃣ What is Goroutine?

## Answer:
A goroutine is a lightweight thread managed by the Go runtime.

Example:

```go
go myFunction()
```

Used for:
- concurrency
- parallel task execution

---

# 4️⃣ What is Channel in Go?

## Answer:
Channels are used for communication between goroutines.

Example:

```go
ch := make(chan int)
```

Channels help prevent:
- race conditions
- shared memory issues

---

# 5️⃣ Difference Between Goroutine and Thread?

| Goroutine | Thread |
|---|---|
| Lightweight | Heavyweight |
| Managed by Go runtime | Managed by OS |
| Faster | Slower |
| Low memory | High memory |

---

# 6️⃣ What is Slice in Go?

## Answer:
Slice is a dynamic and flexible view over arrays.

Example:

```go
nums := []int{1,2,3}
```

Features:
- dynamic size
- append support
- commonly used collection type

---

# 7️⃣ Difference Between Array and Slice?

| Array | Slice |
|---|---|
| Fixed Size | Dynamic Size |
| Value Type | Reference Type |
| Less Flexible | Flexible |

---

# 8️⃣ What is Interface in Go?

## Answer:
Interface defines behavior using method signatures.

Example:

```go
type Shape interface {
    Area() float64
}
```

Go supports implicit interface implementation.

---

# 9️⃣ What is Struct in Go?

## Answer:
Struct is a custom datatype used to group related fields.

Example:

```go
type User struct {
    Name string
    Age int
}
```

---

# 🔟 What is Struct Embedding?

## Answer:
Struct embedding is Go’s alternative to inheritance.

Example:

```go
type Admin struct {
    User
}
```

Go prefers:
- composition
- embedding
instead of inheritance.

---

# 1️⃣1️⃣ What is Pointer?

## Answer:
Pointer stores memory address of variable.

Example:

```go
var x int = 10
var p *int = &x
```

Go supports pointers but not pointer arithmetic.

---

# 1️⃣2️⃣ What is Defer?

## Answer:
`defer` delays execution until function returns.

Commonly used for:
- closing files
- database cleanup
- unlocking mutex

Example:

```go
defer file.Close()
```

---

# 1️⃣3️⃣ What is Panic and Recover?

## Answer:

### Panic
Stops normal execution of program.

### Recover
Used to recover from panic.

Example:

```go
defer func() {
    recover()
}()
```

---

# 1️⃣4️⃣ What is Error Handling in Go?

## Answer:
Go uses explicit error handling.

Example:

```go
value, err := myFunc()
if err != nil {
    return err
}
```

Go avoids try-catch mechanism.

---

# 1️⃣5️⃣ What is Custom Error?

## Answer:
Custom error is user-defined error type.

Example:

```go
type MyError struct {}

func (e MyError) Error() string {
    return "custom error"
}
```

---

# 1️⃣6️⃣ What is WaitGroup?

## Answer:
WaitGroup waits for goroutines to finish execution.

Package:
```go
sync
```

Example:

```go
var wg sync.WaitGroup
```

---

# 1️⃣7️⃣ What is Mutex?

## Answer:
Mutex prevents multiple goroutines from accessing shared resources simultaneously.

Example:

```go
var mu sync.Mutex
```

Used to avoid:
- race conditions

---

# 1️⃣8️⃣ What is JSON Marshal and Unmarshal?

## Answer:

### Marshal
Convert Go struct → JSON

### Unmarshal
Convert JSON → Go struct

Example:

```go
json.Marshal()
json.Unmarshal()
```

---

# 1️⃣9️⃣ What is Context in Go?

## Answer:
Context manages:
- request timeout
- cancellation
- deadlines

Commonly used in:
- APIs
- microservices
- database operations

---

# 2️⃣0️⃣ Why Go is Used in Backend Development?

## Answer:
Go is popular for backend because:
- fast performance
- concurrency support
- scalability
- simple syntax
- efficient memory usage

Used by:
- Google
- Uber
- Docker
- Kubernetes
- Netflix
---

# 👨‍💻 Author

## Dnyanesh Kokate

- GitHub: https://github.com/Dnyanesh0902/GoByExample.git
- LinkedIn: https://www.linkedin.com/in/dnyaneshwar-kokate-04a12b258/

---

# ⭐ Support

If you found this repository useful, give it a ⭐ on GitHub.