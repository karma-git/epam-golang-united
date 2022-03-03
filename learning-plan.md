
<details>
<summary>announcement</summary>

> **AliakseiBurau**
> Hi @here,
>
> We regret to inform you that our second lecture about "Memory model" on 4th of March has been postponed. We made this decision to postpone the event due to circumstances beyond our control. We will announce the new date of the event in a new communication. Stay safe!
> 
> For those of you who don't want to stop, we prepared detailed plan how you can move further.
> Best regards,
> Event Org Committee

</details>

## Specification & Tips and tricks

- https://golang.org/ref/spec - a reference manual of Go language. This is a good step after Go tour, and is generally recommended to check it from time to time

- https://golang.org/doc/effective_go.html - contains tips for writing clear and idiomatic Go code. This is a next  step after you’re done with specification and Go tour

- https://github.com/golang-standards/project-layout - recommended project structure for Golang applications

- http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/ - a collection of common problems and mistakes that most new Go developers face

- https://gobyexample.com - a collection of code examples. You can find code snippets almost for any Golang concept. IT is a good idea to compare your code to these examples if you’re not sure about the quality of your solution. 

- https://medium.com/@blanchon.vincent - an awesome author, who writes really good articles about the internals of Go

## Books

- [Go in Action 1st Edition](https://www.amazon.com/Go-Action-William-Kennedy/dp/1617291781/ref=sr_1_1?dchild=1&keywords=Go+in+Action&qid=1608287051&sr=8-1)
- [Go Programming Language](https://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing/dp/0134190440/ref=sr_1_1?dchild=1&keywords=Golang&qid=1608287028&sr=8-1)

## Learn plan

### 1) Golang basics:

- https://golang.org/doc/code.html
- https://golang.org/ref/spec
- https://golang.org/doc/effective_go.html
- https://tour.golang.org/basics/1
- https://tour.golang.org/flowcontrol/1
- https://blog.golang.org/using-go-modules
- https://github.com/golang-standards/project-layout - recommended

#### What you are expected to know to effectively consider this module complete:

- be familiar with Golang’s basic types.
- know what Constant is and how to define it.
- be able to explain how constants differ from plain variables.
- be able to use the short and long syntax for variables declaration.
- know what exported and non-exported names in Go are.
- know what the package is.
- be able to use loops, know how to use while, plain for, for range loops.
- know what If-else statement is and how to use it.
- know what switch construct is, be able to explain the order of case execution, know what the default case is.
- know what deferred function is, the execution order of deferred functions, common use cases.    

#### Questions for self-check:
- Which basic types are therein Go?
- Can int and int32 variables be compared?
- What are the zero values for int, bool, string, *string?
- What is the difference between comparing plain variables and plain variables with constant?
- How do you implement while loop in Go?
- When are deferred functions being executed?
- In what order the deferred function being executed?
- Are deferred functions called when code panics?
- What is Gomodule?
- How do you add an external dependency to the project?     

#### Tasks:

Use go modules, pick three random emojies from https://github.com/kyokomi/emoji and print them into the console

### 2) Golang data structures Arrays & Slices & Maps:

- https://tour.golang.org/moretypes/6 
- https://golang.org/pkg/sort/
- https://blog.golang.org/slices
- https://blog.golang.org/slices
- https://github.com/golang/go/wiki/SliceTricks
- https://golang.org/pkg/container/  - optional

#### What you are expected to know to effectively consider this module complete:

- know the difference between array and slice
- know the internal representation of the slice
- be able to use slicing operations on slice and array
- know comparability rules for arrays, slices & maps 
- be able to sort slice with sort package, able to sort slice of custom types
- know how to copy arrays, slices & maps

#### Questions for self-check:

- Can arrays be compared with == ?
- Can slices be compared with == ?
- Can maps be compared with == ?
- How do you sort slices with custom types?
- How can you copy a slice?
- How can you copy a map?
- How can you copy an array?
- Which types can be used as a key in a map?
- Is it safe to share a single slice & array across several Goroutines?
- What happens if a few goroutines write to a single map instance?
- Is it safe to insert values into a map that isn’t initialized with the make function?
- Is it safe to append to a slice that isn't initialized with the make function?
- Is it required to always care about the return value from the append function?
- In what order is the traversal over a map(for range loop) performed, How to traverse the map in order sorted by keys?

#### Tasks:
- https://www.hackerrank.com/challenges/array-left-rotation/problem 
- https://www.hackerrank.com/challenges/2d-array/problem 
- https://www.hackerrank.com/challenges/arrays-ds/problem
- https://www.hackerrank.com/challenges/find-the-median/problem 
- https://www.hackerrank.com/contests/codart-2-0/challenges/word-count-1

### 3) Golang strings:

- https://golang.org/ref/spec#String_types
- https://blog.golang.org/strings
- https://golang.org/pkg/strings/
- https://golang.org/pkg/strconv/
- https://golang.org/pkg/bytes/
- https://golang.org/pkg/unicode/utf8/ 
   
#### What you are expected to know to effectively consider this module complete:

- know the difference between runes and bytes
- understand how strings are encoded in Go
- understand UTF-8 encoding
- be able to iterate over string using range loop
- be familiar with strings, strconv, bytes packages

#### Questions for self-check:

- What is the type of variable x := str[0], (str is a plain Golang string)?
- If I cast a string to slice of runes by r := []rune(str), is str going to be copied or not?
- How can you get rune value of a character in string at specific index?
- How can you get the number of bytes and runes in a string? 

#### Tasks:
https://www.hackerrank.com/challenges/camelcase/problem
https://www.hackerrank.com/challenges/caesar-cipher-1/problem
https://www.hackerrank.com/challenges/pangrams/problem
https://www.hackerrank.com/challenges/string-construction/problem
   
### 4) Golang Structs & Methods & Interfaces:

- https://golang.org/doc/codewalk/functions/
- https://tour.golang.org/moretypes/1 
- https://tour.golang.org/methods/1
- https://golang.org/ref/spec#Types
- https://golang.org/ref/spec#Type_identity
- https://golang.org/ref/spec#Method_sets
- https://go101.org/article/type-embedding.html

#### What you are expected to know to effectively consider this module complete:

- know how to define a structure
- know what function and method are and what is their difference
- know what method receivers are and the difference between the value receiver and pointer receiver
- know what method set and interface are
- know how to implement an interface on structure
- know what the empty interface is
- know the difference between pointer and value receivers method sets
- know how interfaces are compared
- be able to use type alias & type cast operation   
- be able to use the type switch
- be able to use type embedding

#### Questions for self-check:

- Can I pass one function to another function?
- In case I define a method on a value receiver can I use it with a pointer receiver?
- In case I define a method on a pointer receiver, can I use it with a value receiver?
- How do you embed one structure to another?
- Is it possible to compare some arbitrary structure with an empty interface?
- Is it possible to compare variables whose types are two different interfaces?
- Is it possible to compare variables whose types are different from each other?
- What are the return values for _, _ :=  i.(string) if i variable is an empty interface and string as an internal representation?
 
#### Tasks:
Build an interface to store an arbitrary set of bytes under specific path
* Interface must contain at least four methods List, Get, Put, Delete
* Implement two version: one to store files on the local filesystem, another - on Amazon S3
* Both implementations must implement the same interface

### 5) Error handling in Golang:

- https://gobyexample.com/errors
- https://github.com/golang/go/wiki/Errors
- https://golang.org/doc/faq#exceptions
- https://blog.golang.org/error-handling-and-go
- https://blog.golang.org/go1.13-errors
- https://blog.golang.org/defer-panic-and-recover

### 6) Golang go routines & Synchronisation:

- https://tour.golang.org/concurrency/1
- https://www.youtube.com/watch?v=f6kdp27TYZs
- https://www.youtube.com/watch?v=QDDwwePbDtw
- https://golang.org/doc/codewalk/sharemem/
- https://golang.org/pkg/sync/
- https://golang.org/pkg/sync/atomic/

#### What you are expected to know to effectively consider this module complete:
- know the difference between goroutines, system threads, processes
- know how to use sync.Mutex, sync.RWMutex, sync.WaitGroup, sync.Pool, sync.Once primitives
- know what a race condition is
- be able to use sync/atomic package
 
#### Questions for self-check:
- What is the difference between a goroutine and a system thread?
- What is the difference between a process and a system thread?
- How can you get the number of logical cores on a current machine in a Go code?
- Which function can we use to set the recommended number of threads used by the Go scheduler?
- What is the difference between sync.Mutex and sync.RWMutex?
- What is the race condition? Why is it dangerous?
- What happens if the main goroutine finishes its execution, while there are still other goroutines running?
 
#### Tasks:
https://en.wikipedia.org/wiki/Dining_philosophers_problem

### 7) Golang channels:

> use materials from the previous section, split just to separate questions [medium](https://medium.com/rungo/anatomy-of-channels-in-go-concurrency-in-go-1ec336086adb)

#### What you are expected to know to effectively consider this module complete:

#### Questions for self-check:
- How are channels initialized?
- What happens if you send an element into a nil channel?
- What happens if you read an element from a nil channel?
- What happens if you send an element into a closed channel?
- What happens if you read an element from a closed channel?
- What is the difference between a non-buffered channel and a channel with a buffer for one element?
- What happens if the buffer is full on a buffered channel?   

#### Tasks:

Write a program which reads full path to a directory and returns statistic about the number of files in this directory grouped by file extension:
- the initial function should start with a file descriptor
- create a buffered channel; as a value you can use map[string]int
- create a wait group (sync.WaitGroup)
- go for each child item
  - if the file descriptor refers to a file, increment a local counter for corresponding file extension
  - if the file descriptor refers to a directory, increment the wait group counter, start a function in a goroutine(a deamon) and pass result channel as a parameter to it. The function should be called for each subitem and, depending on whether it’s a directory or a file, increment file extension counters or start another daemon
  - when child daemon is complete, decrement the wait group counter
 
- read and merge results from the channel. When all child items are processed return summarized result to the caller
 
 As a result, you will get a recursive program which can concurrently scrape directory statistic and print it to console

### 8) Golang Networking(http):

- https://golang.org/doc/articles/wiki/
- https://golang.org/pkg/net/http/
- https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol
- https://golang.org/pkg/context/
- https://benhoyt.com/writings/go-routing/
                       
#### What you are expected to know to effectively consider this module complete:
- know how to initialize, configure, and use http.Client
- know how to create requests using http.NewRequest, send them with client.Do?????
- know how to use the http.ResponseWriter and *http.Request to serve requests
- know how to set up routing using http.ServeMux
- know how to start the http or https service using http.ListenAndServe and http.ListenAndServeTLS
- be able to use http.Server
- be able to implement graceful shutdown for the webserver in Golang
- Questions for self-check:
- What the default timeout value for http.DefaultClient is
- Can you share a single http.Client instance across several goroutines?
- Is it safe to use shared variables inside http handlers without any synchronization mechanisms? Why?
   
#### Tasks
Write simple caching proxy service
<details>
<summary>details</summary>

The main purpose of the service is to proxy requests and save the responses into a cache.

Use the Host Header and path to build an URL of a target service Host Header syntax - [Doc](https://tools.ietf.org/html/rfc7230#section-5.5)

Examples:

```shell
# Request to the proxy (8080 is for non-TLS connections)
curl -H "Host: godoc.org" <http://0.0.0.0:8080/github.com/stretchr/testify/assert>
# Target URL: <http://godoc.org/github.com/stretchr/testify/assert>
# Request to the proxy (8080 is for non-TLS connections)
curl -H "Host: godoc.org:80" <http://0.0.0.0:8080/github.com/stretchr/testify/assert>
# Target URL: <http://godoc.org:80/github.com/stretchr/testify/assert>
# Request to the proxy (9443 is for TLS connections)
curl -H "Host: godoc.org" <http://0.0.0.0:9443/github.com/stretchr/testify/assert>
# Target URL: <https://godoc.org/github.com/stretchr/testify/assert>
# Request to the proxy (9443 is for TLS connections)
curl -H "Host: godoc.org:443" <http://0.0.0.0:9443/github.com/stretchr/testify/assert>
# Target URL: <https://godoc.org:443/github.com/stretchr/testify/assert>
```
**This implies the proxy server should be capable of receiving both HTTP & HTTPS requests.**

- Use Cache-Control Header to control the lifetime of cached requests or to bypass caching. [Doc](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cache-Control)
- Service should only cache safe HTTP methods. See the link to Wikipedia to find out what a safe method is. In case the request isn't cacheable just proxy request
- The cache should be implemented  purely in-memory, i.e. should exist only while the service is running. There's no persistence of cached data in any persistent storage.

</details>

### 9) Golang testing:

- https://golang.org/pkg/testing/
- https://dave.cheney.net/2019/05/07/prefer-table-driven-tests 
- https://github.com/golang/go/wiki/TableDrivenTests
- https://github.com/golang/mock    
- https://www.digitalocean.com/community/tutorials/customizing-go-binaries-with-build-tags 
- https://blog.golang.org/cover
 
#### What you are expected to know to effectively consider this module complete:
- know how to write tests, general naming, required parameters
- know how to run tests for package/project, run tests with specific prefix/suffix
- know how to build code-coverage and visualize code-coverage with the standard set of tools
- know what table testing is and how to apply it to writing tests with the same scenario but different input parameters

#### Questions for self-check:
- Which type of testing package we can use to implement benchmarks?
- Which type of testing package we can use if we need to do initialization before running tests(setup database, create files with random data)?
- How can you run only a subset of tests in a package?
- How can you run benchmarks for packages?
- How build flags used to separate unit/integration/e2e tests    

#### Tasks:
Pick up proxy service from http section
* cover the functions with unit tests as much as  possible
* refer to https://golang.org/pkg/net/http/httptest/ for integration testing
* write a few benchmarks for places where you are not sure about their performance or deciding on several different approaches for a solution
* write a few integration tests with testing.M (build the binary, that behaves like a client using http.Client)   

### 10) Golang Debug/Profiling:

https://golang.org/doc/diagnostics.html 

#### What you are expected to know to effectively consider this module complete:
- know how to debug Go application with Goland IDE or VS Code
- know how to debug race condition
- be able to profile application with net/http/pprof application and go tool pprof

#### Questions for self-check:
- Which tools does Go provide to find race conditions?
- How do you set up net/http/pprof handlers with multiplexers other than default?
- How do you get top functions by allocation in go tool pprof interface?
- How do you get top functions by CPU consumed in go tool pprof interface?
- How can you visualize graphs of memory/cpu/allocation in go tool pprof?     

#### Tasks:
Pick up the proxy service from the http programming section
* try to debug any race condition(in case there is none add it and confirm that you able to find it with the current test suite)
* add integration with net/http/pprof package and find the most expensive places for memory and cpu, try to improve performance
