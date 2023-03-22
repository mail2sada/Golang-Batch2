# Go-Building Blocks
1. ## Data Types
    Data types specify the type of data that a valid Go variable or constant can hold. In Go language, the types are divided into four categories as follows: 
    1. Basic type: 
    2. Aggregate type: 
    3. Reference type: 
    4. Interface type  

    Let's explore  one by one...

    Data Types | Basic | Aggregate | Reference | Interface | 
    --- | --- | --- | --- |--- |
    1 | Numbers | Array | Pointers | interface | 
    2 | strings | structs | Silice |  | 
    3 | booleans |        | maps |  | 
    4 |         |         | functions |  | 
    5 |         |         | channels |  | 
    

    1. Basic Data types.

        Number types are of 3 types

        1. Integers (Whole numbers +ve and -ve)
            Below list of data types are supported in Integers

    Integer Types | Description |
    --- | --- | 
    int8 | 8-bit signed integer | 
    int16 | 16-bit signed integer | 
    int32 | 32-bit signed integer |
    int64 | 64-bit signed integer | 
    uint8 | 8-bit unsigned integer| 
    uint16 | 16-bit unsigned integer | 
    uint32 | 32-bit unsigned integer | 
    uint64 | 64-bit unsigned integer | 
    int | Both int and uint contain same size, either 32 or 64 bit. | 
    uint | Both int and uint contain same size, either 32 or 64 bit. | 
    rune | It is a synonym of int32 and also represent Unicode code points. | 
    byte | It is a synonym of uint8. | 
    uintptr | It is an unsigned integer type. Its width is not defined, but its can hold all the bits of a pointer value.| 

        2. Floating Point (Franctional numbers)
        Go supports 32 and 64 bit floats

    Data Type| 	Description | 
     --- | --- | 
    float32|	32-bit IEEE 754 floating-point number|
    float64|	64-bit IEEE 754 floating-point number|


        3. Complex Numbers
        Go has support for 64 and 128 bit complex numbers.
    Data Type| 	Description | 
     --- | --- |     
    complex64|	Complex numbers which contain float32 as a real and imaginary component.|
    complex128| Complex numbers which contain float64 as a real and imaginary component.|

    ###

    2. strings
    The string data type represents a sequence of Unicode code points. Or in other words, we can say a string is a sequence of immutable bytes, means once a string is created you cannot change that string. A string may contain arbitrary data, including bytes with zero value in the human-readable form.

    3. booleans
    The boolean data type represents only one bit of information either true or false. The values of type boolean are not converted implicitly or explicitly to any other type. 

    Aggregate, Reference, Interface types will be discussed in further section

2. ## Variables
    A typical program uses various values that may change during its execution. variable can store temporarily in the Random Access Memory of the computer and these values in this part of memory vary throughout the execution and hence another term for this came which is known as Variables. So basically, a Variable is a placeholder of the information which can be changed at runtime. And variables allow to Retrieve and Manipulate the stored information.

    1. Delcaring variables

    var keyword is used to specify a variable.

            Syntax 1:
            var variable_name type = expression

            Example:
            var firstInteger int = 10
            var unisignedInt uint = 20

            Syntax 2:

            var variable_name = expression
            type is identified by the compiler at the compile time
            // here myString type will be evaliated string
            var myString = "This is my new string"
            // integer type will be int
            var integar = 100
            // myFloatingPoint will be evaluated to float64
            var myFloatingPoint = 25.45


            Example:
            package main

            import (
	            "fmt"
            )

            func main() {
	            // var variable_name type = expression
	            var myInt int = 10
	            // var is keyword,
	            //myInt is identifier/variable
	            //name and int is data type

	            fmt.Println("Demo: Variable declaration")

	            fmt.Println("Value of myInt is ", myInt)
	            // Printing value of myInt using Println function of fmt package

	            var Integer = 500

	            fmt.Printf("Type of my string is %T\n", Integer)
	            fmt.Println("Value of Integer is ", Integer)

	            var greet = "Welcome to go training..."

	            fmt.Printf("Type of greet is %T", greet)

	            fmt.Println("Value of greet is ", greet)

            }

            https://github.com/mail2sada/Golang-Batch2/blob/main/02_BuildingBlocks/01_Variables/var.go

    1. Declaring multiple variables.

        Multiple variable can be declared as below.

            syntax:
            var var1, var2, var3 type = exp1, exp2, exp3

            Example: 
            var v1, v2, v3 uint8 = 10, 11, 12

            syntax:
            var var1, var2, var3 type

            Expample:
            var v1, v2, v3 int
            v1 = 10
            v2 = -10
            v3 = 1000

            var var1, var2, var3 = exp1, exp2, exp3

            Example:

            var myInt, myFloat, myString = 10, 34.22, "Hello everyone!!!"

            Example:

            package main

            import "fmt"

            func main() {
	            fmt.Println("Demo: declarining multiple variables")

	            var v1, v2, v3 uint = 10, 20, 30

	            fmt.Println("Value of v1 is", v1, "\nValure of v2 ", v2, "\nValue of v3 ", v3)

	            var t1, t2, t3 int

	            t1 = 10

	            t2 = 20

	            t3 = -100

	            fmt.Println("Value of t1 is", t1, "\nValure of t2 ", t2, "\nValue of t3 ", t3)

	            var myInt, myFloat, myString, myBool = 55, 36.82, "Hello everyone", true

	            fmt.Printf("Type of myInt %T\nType of myFloat %T\nType of myString %T\nType of myBool %T", myInt, myFloat, myString, myBool)

	            fmt.Println("Value of myInt, myFloat, myString and myBool are", myInt, myFloat, myString, myBool)

            }       



3. ## Constants
4. ## Rune in Golang
5. ## Operators in Golang
6. ## Scope of Variables
7. ## Type Casting
8. ## var Keyword in Golang
9. ## Short Declaration Operator(:=)
10. ## var keyword vs short declaration operator
11. ## Decision Making Statements
12. ## Loops in Golang
13. ## Loop Control Statements
14. ## Switch Statement in Go