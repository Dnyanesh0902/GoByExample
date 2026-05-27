# Go By Example - Complete Go Learning Repository 🚀

A comprehensive Go (Golang) learning repository built by practicing concepts from **Go by Example** with hands-on coding, interview preparation, and backend development concepts.

This repository contains:
- Practical Go programs
- Concurrency examples
- File handling examples
- JSON/XML processing
- HTTP server examples
- Goroutines & channels
- Testing & benchmarking
- Interview Questions & Answers

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

### Interview Question

## Q: What is the entry point of Go program?

### Answer:
The execution of Go program starts from:

```go
func main()
```

---

## Variables

### Key Concepts

```go
var name string
age := 25
```

### Interview Questions

## Q: Difference between `var` and `:=` ?

### Answer:

| var | := |
|---|---|
| Explicit declaration | Short declaration |
| Can be used globally | Only inside functions |
| Requires datatype sometimes | Type inferred automatically |

---

## Constants

### Interview Question

## Q: What is constant in Go?

### Answer:
Constants are immutable values declared using:

```go
const
```

Their values cannot be changed after declaration.

---

## Loops

### Interview Question

## Q: Which loop exists in Go?

### Answer:
Go supports only one loop:

```go
for
```

It replaces:
- while
- do-while

---

# 2️⃣ Collections

## Arrays

### Interview Question

## Q: What is array?

### Answer:
Array is a fixed-size collection of elements.

Example:

```go
var nums [5]int
```

---

## Slices

### Key Concepts
- append()
- make()
- slicing

### Interview Questions

## Q: Difference between Array and Slice?

| Array | Slice |
|---|---|
| Fixed Size | Dynamic Size |
| Value Type | Reference Type |
| Less Flexible | Flexible |

---

## Maps

### Interview Question

## Q: What is map in Go?

### Answer:
Map stores key-value pairs.

Example:

```go
m := map[string]int{
    "age": 25,
}
```

Maps are unordered collections.

---

# 3️⃣ Functions

## Multiple Return Values

### Interview Question

## Q: Why multiple return values are important in Go?

### Answer:
Go commonly returns:

```go
value, error
```

This improves explicit error handling.

---

## Variadic Functions

### Interview Question

## Q: What is variadic function?

### Answer:
Function accepting variable number of arguments.

Example:

```go
func sum(nums ...int)
```

---

## Closures

### Interview Question

## Q: What is closure in Go?

### Answer:
Closure is a function that remembers variables from outer scope.

---

# 4️⃣ Pointers & Memory

## Pointers

### Interview Question

## Q: What is pointer in Go?

### Answer:
Pointer stores memory address of variable.

Example:

```go
x := 10
p := &x
```

Go supports pointers but does not support pointer arithmetic.

---

# 5️⃣ Structs & Interfaces

## Structs

### Interview Question

## Q: What is struct?

### Answer:
Struct is a custom datatype used to group related fields.

Example:

```go
type User struct {
    Name string
    Age int
}
```

---

## Interfaces

### Interview Question

## Q: What is interface in Go?

### Answer:
Interface defines behavior using method signatures.

Example:

```go
type Shape interface {
    Area() float64
}
```

Go supports implicit interface implementation.

---

## Struct Embedding

### Interview Question

## Q: Does Go support inheritance?

### Answer:
No.

Go uses:
- composition
- struct embedding

Example:

```go
type Admin struct {
    User
}
```

---

## Generics

### Interview Question

## Q: What are generics in Go?

### Answer:
Generics allow reusable type-safe code.

Introduced in:
```text
Go 1.18
```

---

# 6️⃣ Error Handling

## Errors

### Interview Question

## Q: Why Go avoids try-catch?

### Answer:
Go uses explicit error handling for better readability and simplicity.

Example:

```go
if err != nil {
    return err
}
```

---

## Custom Errors

### Interview Question

## Q: How to create custom error?

### Answer:

```go
type MyError struct {}

func (e MyError) Error() string {
    return "custom error"
}
```

Custom errors implement:

```go
Error() string
```

---

## Defer

### Interview Question

## Q: What is defer?

### Answer:
`defer` delays execution until surrounding function returns.

Commonly used for:
- closing files
- database cleanup
- unlocking mutex

Example:

```go
defer file.Close()
```

---

## Panic & Recover

### Interview Question

## Q: Difference between panic and recover?

### Answer:

| Panic | Recover |
|---|---|
| Stops execution | Handles panic |
| Used in critical failures | Prevents crash |

---

# 7️⃣ Concurrency

## Goroutines

### Interview Question

## Q: What is goroutine?

### Answer:
Goroutine is a lightweight thread managed by Go runtime.

Example:

```go
go myFunction()
```

---

## Channels

### Interview Question

## Q: What is channel?

### Answer:
Channels are used for communication between goroutines.

Example:

```go
ch := make(chan int)
```

Channels help avoid race conditions.

---

## WaitGroups

### Interview Question

## Q: What is WaitGroup?

### Answer:
WaitGroup waits for multiple goroutines to complete execution.

Package:

```go
sync.WaitGroup
```

---

## Mutex

### Interview Question

## Q: Why mutex is used?

### Answer:
Mutex protects shared resources and avoids race conditions.

Example:

```go
var mu sync.Mutex
```

---

# 8️⃣ JSON & XML

## JSON

### Interview Question

## Q: What is Marshal and Unmarshal?

### Answer:

| Function | Purpose |
|---|---|
| Marshal | Go Struct → JSON |
| Unmarshal | JSON → Go Struct |

---

## XML

### Interview Question

## Q: Why struct tags are used in XML/JSON?

### Answer:
Struct tags define mapping between struct fields and external data formats.

Example:

```go
Name string `json:"name"`
```

---

# 9️⃣ File Handling

## Reading Files

### Interview Question

## Q: Which package is used for file handling?

### Answer:
Common packages:
- os
- io
- bufio

Example:

```go
os.ReadFile()
```

---

# 🔟 Testing & Benchmarking

## Testing

### Interview Question

## Q: Naming convention for Go test files?

### Answer:

```text
*_test.go
```

---

## Benchmarking

### Interview Question

## Q: How to run benchmark test?

### Answer:

```bash
go test -bench=.
```

---

# 1️⃣1️⃣ Networking

## HTTP Server

### Interview Question

## Q: Which package is used to build APIs in Go?

### Answer:

```go
net/http
```

---

## Context

### Interview Question

## Q: Why context package is important?

### Answer:
Context manages:
- timeout
- cancellation
- deadlines
- request lifecycle

Used heavily in:
- APIs
- microservices
- database operations

---

# 1️⃣2️⃣ System Programming

## Spawning Processes

### Interview Question

## Q: Which package executes shell commands?

### Answer:

```go
os/exec
```

---

# 🛠 Technologies Used

- Go (Golang)
- Go Standard Library
- VS Code
- Git & GitHub

---

# 📂 Repository Structure

```text
GoByExample/
│
├── 01-HelloWorld
├── 02-Values
├── 03-Variables
├── 04-Constants
├── 05-For
├── 06-If_Else
├── 07-Switch
├── 08-Arrays
├── 09-Slices
├── 10-Maps
├── 11-Functions
├── 12-Closures
├── 13-Recursion
├── 14-Pointers
├── 15-strings_runs
├── 16-Structs
├── 17-Methods
├── 18-Interface
├── 19-Enums
├── 20-Struct_Embedding
├── 21-Generics
├── 22-Range_Over_Iterators
├── 23-Errors
├── 24-Custom_errors
├── 25-Routines
├── 26-Channels
├── 27-Channel-Buffering
├── 28-Channel-Synchronization
├── 29-Channel-Directions
├── 30-Select
├── 31-Timeouts
├── 32-Non-Blocking-Channel-Operations
├── 33-Closing-Channel
├── 34-Range-over-Channels
├── 35-Timers
├── 36-Tickers
├── 37-Worker-Pools
├── 38-WaitGroups
├── 39-Rate-Limiting
├── 40-Atomic-Counters
├── 41-Mutexes
├── 42-Stateful-Goroutines
├── 43-Sorting
├── 44-Sorting-by-functions
├── 45-Panic
├── 46-Defer
├── 47-Recover
├── 48-String-Functions
├── 49-String-Formatting
├── 50-Text-Templates
├── 51-Regular-Expressions
├── 52-JSON
├── 53-XML
├── 54-Time
├── 55-Epoch
├── 56-Time-Formatting-Parsing
├── 57-Random-Numbers
├── 58-Number-Parsing
├── 59-URL-Parsing
├── 60-SHA256-Hashes
├── 61-Base64-Encoding
├── 62-Reading-files
├── 63-Writing-files
├── 64-Line-Filters
├── 65-File-Path
├── 66-Directories
├── 67-Temporary-Files-And-Directories
├── 68-Embed-Directive
├── 69-Testing-and-benchmarking
├── 70-Command-Line-Arguments
├── 71-Command-Line-Flags
├── 72-Command-Line-SubCommands
├── 73-Environment-Variables
├── 74-Logging
├── 75-HTTP-Client
├── 76-HTTP-Server
├── 77-TCP-Server
├── 78-Context
├── 79-Spawning-Processes
├── 80-Execing-Processes
├── 81-Signals
├── 82-Exit
│
├── README.md
├── go.mod
├── go.sum
│
├── subdir/
└── exit/
```
---
# ▶️ Run Project

Run example:

```bash
go run main.go
```

Run tests:

```bash
go test -v
```

Run benchmark:

```bash
go test -bench=.
```

---

# 🎯 Learning Goals

- Learn Go from Basic to Advanced
- Build Strong Backend Development Skills
- Understand Concurrency Internals
- Prepare for Go Backend Interviews
- Build Production-Ready APIs & Microservices

---

# 👨‍💻 Author

## Dnyanesh Kokate

- GitHub: https://github.com/Dnyanesh0902/GoByExample
- LinkedIn: https://www.linkedin.com/in/dnyaneshwar-kokate-04a12b258/
- Portfolio : https://dnyanesh.miracledevelopers.in

---

# ⭐ Support

If you found this repository useful, give it a ⭐ on GitHub.