# Golang-Batch2.  
### Overview
1. Go Features
1. Advantages
2. When to use go?
1. Installation (Windows & Mac)
2. Hello World in Golang
3. Golang code structures
1. Identifiers in Golang
2. Keywords in Golang

### Go-Building Blocks
1. Data Types
2. Variables
3. Constants
4. Rune in Golang
5. Operators in Golang
6. Scope of Variables
7. Type Conversion
9. Short Declaration Operator(:=)
11. Decision Making Statements
12. Loops in Golang
13. Loop Control Statements
14. Switch Statement in Go


### Array and Slice
1. Arrays in Golang, 
2. Copying an Array into Another Array in Golang, 
3. Passing an Array to a Function in Golang, 
4. Slices in Golang
5. Slice Composite Literal
6. Copying one Slice into another Slice
7. Passing a Slice to Function
8. Comparing two Slices in Golang
9. Checking the Equality of Slices in Golang
10. Sorting a Slice in Golang
11. Trimming a Slice in Golang
12. Splitting a Slice in Golang,
13. Slice Sort, 
14. Reverse, 
15. Search Functions

### Strings 
1. Strings in Golang
2. Different ways to compare Strings
3. Different ways to concatenate two strings
4. Trimming a String in Golang
5. Splitting a String in Golang
6. Check if the given characters are present in String
7. Repeating a String for Specific Number of Times
8. Finding the index value of specified string
9. Counting the Number of Repeated Characters in String

### Structures
1. Overview
2. Structure Equality
3. Nested Structure
4. Anonymous Structure and Fields
5. Promoted, Fields in Structure
6. Function as a Field in Structure

### Pointers
1. What is pointer?
2. Pointer to a variable
3. Pointers to a Function
4. Returning Pointer from a Function
5. Pointer to an Array as Function Argument
6. Pointer to Struct
7. Comparing Pointers
8. Finding the Capacity of the Pointer
9. Finding the Length of the Pointer

### Maps Overview
1. Allowed Key types in a Map
2. Allowed Value types in a Map
3. Creating a Map
4. Using the map[<key_type>]<value_type>
5. Using Make
6. Map Operations
7. Zero Value
8. Maps are referenced data types
9. Iterate over a map
10. Maps are not safe for concurrent use
11. Slice Sort
12. Reverse and Search Functions


### Functions & Methods
1. What are the Functions?
2. Variadic Function
3. Anonymous Function
4. main and init function
5. Function Arguments
6. Function Returning Multiple Values
7. Named Return Values
8. Blank Identifier
9. Defer
10. Methods
11. Methods With Same Name
12. Promoted Methods in Structure

### Interfaces
1. Multiple Interfaces
2. Embedding Interfaces
3. Polymorphism Using Interfaces
4. Methods in connection with interfaces

### Error Handling
1. error interface
2. Advantages of using error as a type
3. Different ways of creating an error
4. Using errors.New(“some_error_message”)
5. Using fmt.Errorf(“error is %s” “some_error_message”) 
6. Creating Custom error, 
7. Ignoring errors, 
8. Wrapping of error, 
9. Unwrap an error, 
10. Check if two error are equal
11. Using the equality operator (==)
12. Using the Is function of errors package
13. Get the underlying error from an error represented by the error interface
14. Using the .({type}) assert, 
15. Using the As function of errors package, 
16. Runtime Error Panic
17. Calling the panic function explicitly
18. Panic with defer
19. Recover in golang
20. Panic/Recover and Goroutine
21. Printing stack trace
22. Return value of the function when panic is recovered


### Date and Time
1. Create a new time
2. Using time.Now()
3. Using time.Date()
4. Understanding Duration
5. Add or Subtract to a time
6. Add to time, Subtract to time
7. Time Parsing/Formatting
8. Time Parse Example
9. Time Formatting Example
10. Time Diff
11. Time Conversion
12. Convert time between
13. different timezones
14. Timers
15. Tickers


### Goroutines – Concurrency
1. Overview of Goroutine
2. Main Goroutine 
3. Goroutine vs Thread, 
4. Channel in Golang, 
    1. Unbuffered channel
    2. Buffer channel 
    3. Channel direction    
    4. Channel Ownership
5. Unidirectional Channel in Golang, 
6. Synchronization in Golang
    1. Sync package overview
    2. Waitgroups
    3. Mutex
    4. Conditional Variables
    5. Atomic
    6. Sync Once
    7. Sync Pool
7. Concurrency Patterns
    1. Piplines
    2. Fan out & Fan in
    3. Cancelling go-routines


7. Solving Dining Philosopher problem

### Context Interface
1. Creating New Context
2. Context Tree
3. Deriving From Context
4. Best Practices

### Packages and Modules
1. Packages
2. Modules
3. Exported and Unexported Names
4. Nested Packages
5. Alias in importing
6. Init functions  
7. Package Naming Convention
8. Types of Modules
9. Package vs Module
10. Do a go get
11. go mod command 
12. Direct vs Indirect Dependencies in go.mod file


### Files and Directories
1. Overview of os package
2. Understanding Files and Directories
3. Filesystem package
4. Creating files and directories
5. Working with files
6. File seek
7. io/fs/ioutil/bufio package

### JSON/XML
1. Overview of Encoding package
2. Json/XML Marshling, 
3. Json/XML Unmarshaling, 

### Networking
1. Working with tcp
2. Working with udp
3. http package
    1. Request, 
    2. Response, 
    3. Pair of API signature and its handler, 
    4. Mux, 
    5. Listener,
    6. Using server’s ListenAndServe function
    7. Using http's ListenAndServe function

### Important resources and tutorials


Section 1: Go Programming languages 

https://www.geeksforgeeks.org/golang/?ref=lbp  

Deep dive data structures in Golang (Arrays, Slices, Maps, Structs, Interfaces) – Practice programs 

Integrate pProfs for programs you write - https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/ 

https://www.golangprograms.com/basic-programs.html (Execute every single program here) 

https://golangbyexample.com/golang-comprehensive-tutorial/ 

https://golongwithgolang.com/when-should-you-use-pointers-in-go 

https://golongwithgolang.com/thread-safety-in-golang 

https://golongwithgolang.com/the-glorious-nethttp-go-module-1 

https://www.geeksforgeeks.org/how-to-use-go-with-mysql/ 

Section 2: Go Libraries 

Goroutines in Golang 

https://www.golangprograms.com/goroutines.html 

https://www.digitalocean.com/community/tutorials/how-to-run-multiple-functions-concurrently-in-go 

https://www.golangprograms.com/go-language/channels.html 

https://www.golangprograms.com/go-language/concurrency.html 

https://blog.logrocket.com/concurrency-patterns-golang-waitgroups-goroutines/#:~:text=Golang%20provides%20goroutines%20to%20support,of%20stack%20space%20to%20initialize. 

GORM in Golang 

https://gorm.io/docs/index.html 

https://www.mindbowser.com/golang-go-with-gorm/ 

https://towardsdev.com/introduction-to-orm-using-gorm-in-golang-d1936a45ffdb 

https://golangdocs.com/gorm-golang-orm-package 

https://www.youtube.com/watch?v=dpx6hpr-wE8 

Go Gin 

https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin 

https://blog.logrocket.com/building-microservices-go-gin/ 

https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/ 

Golang Mux Router 

https://gowebexamples.com/routes-using-gorilla-mux/ 

https://golangdocs.com/golang-mux-router 

https://medium.com/geekculture/develop-rest-apis-in-go-using-gorilla-mux-5869b2179a18 

https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql 

https://www.golangprograms.com/goroutines-and-channels-example.html 

 

Section 3: Caching and Data streaming tools 

Caching 

https://www.geeksforgeeks.org/what-is-the-caching-mechanism/ 

https://medium.datadriveninvestor.com/all-things-caching-use-cases-benefits-strategies-choosing-a-caching-technology-exploring-fa6c1f2e93aa 

 

 

Golang and caching mechanism 

https://kislayverma.com/software-architecture/architecture-patterns-caching-part-1/ 

https://www.mailgun.com/blog/it-and-engineering/golangs-superior-cache-solution-memcached-redis/ 

https://github.com/mailgun/groupcache 

https://stackoverflow.com/questions/31389484/how-does-groupcache-in-go-compare-to-redis-and-memcached 

https://medium.com/@shreyanshsinha/thundering-herd-problem-cache-stampede-187a7b2f6cf7 

 

 

Redis (Prerequisite: Well versed with Data structures)  

https://www.youtube.com/watch?v=Qu5gX2uOaL8 

https://tutorialedge.net/golang/go-redis-tutorial/ 

https://github.com/gomodule/redigo 

 

KAFKA  

https://www.tutorialspoint.com/apache_kafka/apache_kafka_introduction.htm 

https://www.sohamkamani.com/golang/working-with-kafka/ 

https://github.com/confluentinc/confluent-kafka-go 

Videos: 

https://www.youtube.com/watch?v=jFfo23yIWac 

https://youtu.be/YS4e4q9oBaU 


