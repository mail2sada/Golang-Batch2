# Overview

1. ## Introduction

Also called Golang, Go is an open-source, statically typed,  and compiled programming language designed by Rob Pike, Robert Griesemer, and Ken Thompson. The language, that appeared in the market in 2009, was designed with an intention to enhance programming productivity in the era of networked machines, multicore, and huge codebases. Something for which the Google team picked the best characteristics of the popular languages.
      
    1. Static typing and runtime efficiency of C++.
    2. Usability and Readability of Python and JavaScript.
    3. Object Oriented Programming (OOPs) concept of Smalltalk. (https://en.wikipedia.org/wiki/Smalltalk / https://squeak.org/)
    4. Concurrency element of Newsqueak. (https://www.wikidata.org/wiki/Q262003)

The language has just entered into its version 1.21 But, has gained a huge momentum in the market – bringing it several steps ahead in Go vs Rust discussion and similar comparisons. It has entered into the list of Top 10 Programming language by IEEE Spectrum and become the fourth most active languages on GitHub.

1. ## Go Features
    1. #### Open-Source
    The foremost characteristic of Golang programming language is that it is open-source. That means, anyone can download and experiment with the code to bring better codes into picture and fix related bugs.

    2. #### Static Typing
    Go is a statically typed programming language and works with a mechanism that makes it possible to compile code accurately while taking care of type conversions and compatibility level. This gives developers freedom from challenges associated with dynamically typed languages.

    3. #### Concurrency Support
    One of the prime characteristics of go programming language is its concurrency support. Golang, unlike other programming languages, offers easier and trackable concurrency options. This makes it easier for app developers to complete requests at a faster pace, free up allocated resources and network earlier, and much more.

    4. #### Powerful Standard Library
    This programming language also comes loaded with a robust standard library. This libraries offer ample components that gives developers an escape from turning towards third party packages anymore.

    5. #### Powerful tool set
        1. Gofmt: It automatically formats your Go code, which eventually brings a major impact on readability.
        2. Gorun: This tool is used to add a ‘bang line’ in the source code to run it, or run a similar sode code file explicitly. It is often used by Go developers when experimenting with codes written in Python.
        3. Goget: The Goget tool downloads libraries from GitHub and save it to your GoPath so that you can easily import the libraries in your app project.
        4. Godoc: The tool parses Go source code, including comments and creates a documentation in HTML or plain text format. The documentation made is tightly coupled with codes it documents and can be easily navigated with one click.

    6. #### Testing Capabilities
    Go language also offers an opportunity to write unit tests along with writing the app codes. Besides, it avails support to understand code coverage, benchmark tests, and write example codes to create your own code documentation.
    7. #### Garbage Collection
    The programming language also offers exceptional power of garbage collection. Meaning, developers need not worry about freeing up pointers or the situation associated with dangling pointer. Because of these characteristics, the language has gained attention of various popular brands.

2. ## Advantages
    1. #### Easy to use
    One of the biggest benefits of Golang language is that its syntax is similar to that of C and C++. There are not many complex functions to learn and implement. Besides, the documentation is simple and can be used quickly.
    A ripple effect of which is that the number of developers taking an interest in working with Go language for building native and cross-platform applications are increasing.

    2. #### Cross-Platform Development Opportunities
    Another advantage of going with this language is that multiple types of applications can be developed in Go. This includes software and mobility solutions for Windows, Unix, Linux, BSD versions, and mobile platforms.

    3. #### Faster Compilation and Execution
    In Golang development environment, there is no Virtual Machine. The code is directly compiled to machine code, which makes the compilation process faster and more effective. Besides, the compiler produces only one executable file after compilation. This file does not suffer from any kind of dependency and can be uploaded and run anywhere. This makes the code execution process faster.

    4. #### Scalable
    Another feature that makes Golang outshine other programming languages is that it is highly scalable in nature. It enables top android mobile app developers to handle multiple tasks at the same time, especially channels and goroutines.

    5. #### Time Saving
    With features like automatic declaration of variables, latency free garbage collection, and faster compile time, it saves consideration development time; giving developers enough time to be productive.
3. ## Drawbacks
    1. #### Still Developing
    Though the popularity graph of Go programming language is growing and it holds a promising future in the market, it is still in its nascent stage. This is making it tough for the language to beat the competition with popular names like Java.

    2. #### Too Simple
    Programming languages like Swift and Haskell are difficult to learn. But, at the same time, they offer a myriad of opportunities to perform abstractions and other complex processes. Something that is not possible in the case of Golang because of its ultra-simple design.
    So, the most primary characteristic turns out to be one of the major disadvantages of golang language.

    3. #### Absence of GUI Library
    Another con of go programming language is that it does not offer native support for GUI libraries. This implies app development companies have to connect a library to their application manually, rather than using native solutions like that in the case of Java or Python.

    4. #### No Specific Niche
    Designed by Google with the motive to deliver endless support and solutions, the language has characteristics that goes well with all. However, the language has still not found a single niche to conquer. While it is facing competition from JavaScript in frontend development world, it is lagging behind Python in the world of data visualization and analysis. Because of this, many developers are still hesitating from investing in this language for a particular purpose.

    5. #### Poor Error Handling
    In the case of Go language, a function is required to return error if any error is expected. Developers have to write a huge number of ‘if’ blocks to perform error handling process efficiently and effectively. Something that is one of the golang coding challenges when they have lost track to the error that can further result in missing out of some imperative error handling logic.

3. ## Where go is used?
    Go is used for a wide range of purposes, including:

    1. #### Web development: 
    Go has a built-in package for creating HTTP servers and clients, making it a good choice for building web applications and microservices.

    2. #### Networking: 
    Go has strong support for concurrent programming and networking, making it a good choice for building networked systems and servers.

    3. #### Distributed systems: 
    Go's simplicity and efficiency make it well-suited for building distributed systems that need to handle a large number of concurrent connections.

    4. #### Data processing: 
    Go's support for concurrency and parallelism makes it well-suited for data processing tasks that can be parallelized.

    5. #### System programming: 
    Go can be used to build low-level systems programs, such as operating system kernels and device drivers.

4. ## Installation (Windows & Mac)
    Go has 2 components, go compiler and editor.
    Go compiler can be downloaded from the official site. (https://go.dev/doc/install)

    We can use any of the editors Visual Studio Code, Notepad, etc. Visal Studio Code comes with lot of extensions and it is easy to use.
    https://code.visualstudio.com/docs/setup/mac

    On successful installation of visual studio code, need to install go extension for visual studio code this will enable auto complete, auto format, etc.

    VSCode extension: Rich Go language support for Visual Studio Code

5. ## Hello World in Golang

    helloworld is best introduction to any programming language, Lets see below code snippet
    <i><b>
    // helloworld.go
    // Few things to note
    /*
    Every go file has extension .go
    every go file belongs to a package
    */

        package main // this go source belongs to main package.

        import "fmt" // fmt package is imported, so that it functionalities can be used

        // func is keyword, main is name of the function, () it is not taking any arguments

        func main() {

	        //fmt package has a function with the name Println

	        // it can take any number of arguments and arguments can be of any type.

	        fmt.Println("Welcome to Mavenir Go lang training")  

        }
    <b></i>
    https://github.com/mail2sada/Golang-Batch2/blob/main/01_Overview/helloworld.go

    Lets analyze source code

            package main

    <b><i>package</i></b> key word defines the package where the source code belongs. Every go source must belong to a package. helloword.go belongs to main package

            import "fmt"

    <b><i>import</i></b> key word specifies which external packages are getting used. We can import external package, this enables code reusablity. We are importing fmt package

            func main()

    func key word is used to define a function, we are defining main function. main() function in main package is the entry point of execution.  Functions are enclosed under curly braces.

            fmt.Println("Welcome to Mavenir Go lang training")
    
    fmt.Println is a function defined in fmt package this will print the specified data. Here wre are printing "Welcome to Mavenir Go lang training" 

6. ## Golang code structures
7. ## Identifiers in Golang
8. ## Keywords in Golang