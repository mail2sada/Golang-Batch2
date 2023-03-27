# Strings 
-   ## Strings in Golang
    -  ### It is a sequence of variable-width characters
    -  ### Every character is represented by one or more bytes using UTF-8 Encoding
    -  ### Strings are the immutable chain of arbitrary bytes, including bytes with zero values
    -  ### String is a read-only slice of bytes
    -  ### Bytes of the strings can be represented in the Unicode text using UTF-8 encoding

    -   ## String Literals
        -   ### Using double quotes(“”)
        -   ### String literals are created using double-quotes(“”)
        -   ### This type of string support escape character Using backticks(“)
        -   ### String literals are created using backticks(“) and known as raw literals
        -   ### Raw literals do not support escape characters
        -   ### Can span multiple lines
        -   ### Can contain any character except backtick
        -   ### Used for writing multiple line message, in the regular expressions, and in HTML.
-   ## Different ways to compare Strings
-   ## Different ways to concatenate two strings
-   ## Trimming a String in Golang
-   ## Splitting a String in Golang
-   ## Check if the given characters are present in String
-   ## Repeating a String for Specific Number of Times
-   ## Finding the index value of specified string
-   ## Counting the Number of Repeated Characters in String

<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>

# Structures
-   ## Overview
-   ## Structure Equality
-   ## Nested Structure
-   ## Anonymous Structure and Fields
-   ## Promoted, Fields in Structure
-   ## Function as a Field in Structure

<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>

# Pointers
-   ## What is pointer?
-   ## Pointer to a variable
-   ## Pointers to a Function
-   ## Returning Pointer from a Function
-   ## Pointer to an Array as Function Argument
-   ## Pointer to Struct
-   ## Comparing Pointers
-   ## Finding the Capacity of the Pointer
-   ## Finding the Length of the Pointer

<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>

# Maps Overview
-   ## Allowed Key types in a Map
-   ## Allowed Value types in a Map
-   ## Creating a Map
-   ## Using the map[<key_type>]<value_type>
-   ## Using Make
-   ## Map Operations
-   ## Zero Value
-   ## Maps are referenced data types
-   ## Iterate over a map
-   ## Maps are not safe for concurrent use
-   ## Sort by keys and sort by value
-   ## Reverse and Search Functions

<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>


# Functions & Methods
-   ## What are the Functions?
-   ## Variadic Function
-   ## Anonymous Function
-   ## main and init function
-   ## Function Arguments
-   ## Function Returning Multiple Values
-   ## Named Return Values
-   ## Blank Identifier
-   ## Defer
-   ## Methods
-   ## Methods With Same Name
-   ## Promoted Methods in Structure

<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>


# Interfaces
-   ## Multiple Interfaces
-   ## Embedding Interfaces
-   ## Polymorphism Using Interfaces
-   ## Methods in connection with interfaces

<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>

# Error Handling
-   ## error interface
-   ## Advantages of using error as a type
-   ## Different ways of creating an error
-   ## Using errors.New(“some_error_message”)
-   ## Using fmt.Errorf(“error is %s” “some_error_message”) 
-   ## Creating Custom error, 
-   ## Ignoring errors, 
-   ## Wrapping of error, 
-   ## Unwrap an error, 
-   ## Check if two error are equal
-   ## Using the equality operator (==)
-   ## Using the Is function of errors package
-   ## Get the underlying error from an error represented by the error interface
-   ## Using the .({type}) assert, 
-   ## Using the As function of errors package, 
-   ## Runtime Error Panic
-   ## Calling the panic function explicitly
-   ## Panic with defer
-   ## Recover in golang
-   ## Panic/Recover and Goroutine
-   ## Printing stack trace
-   ## Return value of the function when panic is recovered

<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>


# Date and Time
-   ## Create a new time
-   ## Using time.Now()
-   ## Using time.Date()
-   ## Understanding Duration
-   ## Add or Subtract to a time
-   ## Add to time, Subtract to time
-   ## Time Parsing/Formatting
-   ## Time Parse Example
-   ## Time Formatting Example
-   ## Time Diff
-   ## Time Conversion
-   ## Convert time between
-   ## different timezones
-   ## Timers
-   ## Tickers


<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>


# Goroutines – Concurrency
-   ## Overview of Goroutine
-   ## Main Goroutine 
-   ## Goroutine vs Thread, 
-   ## Channel in Golang, 
    -   ### Unbuffered channel
    -   ### Buffer channel 
    -   ### Channel direction    
    -   ### Channel Ownership
-   ## Unidirectional Channel in Golang, 
-   ## Synchronization in Golang
    -   ### Sync package overview
    -   ### Waitgroups
    -   ### Mutex
    -   ### Conditional Variables
    -   ### Atomic
    -   ### Sync Once
    -   ### Sync Pool
-   ## Concurrency Patterns
    -   ### Piplines
    -   ### Fan out & Fan in
    -   ### Cancelling go-routines



<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
<br/>
