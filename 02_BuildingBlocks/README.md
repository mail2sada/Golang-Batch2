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

    https://github.com/mail2sada/Golang-Batch2/blob/main/02_BuildingBlocks/01_Variables/01_var.go

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

    https://github.com/mail2sada/Golang-Batch2/blob/main/02_BuildingBlocks/01_Variables/02_var.go

3. ## Constants
    As the topic suggest constant means a identifier whose value is not goiing to change.

            Syntax:

                const const_name type = expression
                Example:
                const pi float64 = 3.14159
            Syntax:
                const const1, const2, const3 type = expr1, expr2, expr
                Exmaple:
                const second, minute, hour = 1, 1*60, 1 * 60 * 60
             Syntax:
                const cnst1, cnst2, cnst3 = exp1, exp2, exp3
                Example:
                const pi, second, greeting = 3.14159, 1000, "Wecome to all!!!"

            Example:
                package main

                import "fmt"

                func main() {

	                fmt.Println("Demo: Constant declaration")

	                const pi float64 = 3.14159

	                fmt.Printf("Type of pi is %T", pi)

	                const milliSecond = 1
	                const second = 1000 * milliSecond
	                const minute = 60 * second
	                const hour = 60 * minute

	                const greeting = "Welcome to go training..."

	                fmt.Println("Value of milliSecond:", milliSecond)
	                fmt.Println("Value of second:", second)
	                fmt.Println("Value of minute:", minute)
	                fmt.Println("Value of hour:", hour)
	
	                fmt.Println(greeting)
                }   
    https://github.com/mail2sada/Golang-Batch2/blob/main/02_BuildingBlocks/02_Constant/constant.go

4. ## Rune in Golang

    In the past, we only had one character set, and that was known as ASCII (American Standard Code for Information Interchange). There, we used 7 bits to represent 128 characters, including upper and lowercase English letters, digits, and a variety of punctuations and device-control characters. Due to this character limitation, the majority of the population is not able to use their custom writing systems. To solve this problem, Unicode was invented. Unicode is a superset of ASCII that contains all the characters present in today’s world writing system. It includes accents, diacritical marks, control codes like tab and carriage return, and assigns each character a standard number called “Unicode Code Point”, or in Go language, a “Rune”. The Rune type is an alias of int32. Important Points:

    Always remember, a string is a sequence of bytes and not of a Rune. A string may contain Unicode text encoded in UTF-8. But, the Go source code encodes as UTF-8, therefore, no need to encode the string in UTF-8.

    UTF-8 encodes all the Unicode in the range of 1 to 4 bytes, where 1 byte is used for ASCII and the rest for the Rune.
    ASCII contains a total of 256 elements and out of which, 128 are characters and 0-127 are identified as code points. Here, code point refers to the element which represents a single value.

    1. Rune Literal
        It represents a Rune constant, where an integer value recognizes a Unicode code point. In Go language, a Rune Literal is expressed as one or more characters enclosed in single quotes like ‘g’, ‘\t’, etc. In between single quotes, you are allowed to place any character except a newline and an unescaped single quote. Here, these single-quoted characters themselves represent the Unicode value of the given character and multi-character sequences with a backslash (at the beginning of the multi-character sequence) encode values in a different format. In Rune Literals, all the sequences that start with a backslash are illegal, only the following single-character escapes represent special values when you use them with a backslash: 

    Character | Unicode | Description |
    --- | --- | --- |
    \a	| U+0007	| Alert or Bell | 
    \b	| U+0008	| backspace | 
    \f	| U+000C	| form feed | 
    \n	| U+000A	| line feed or newline | 
    \r	| U+000D	| carriage return | 
    \t	| U+0009	| horizontal tab | 
    \v	| U+000b	| vertical tab | 
    \\	| U+005c	| backslash | 
    \’	| U+0027	| single-quote | 
    \”	| U+0022	| double quote(legal only in string literals) |     


    Example
            package main

            import "fmt"

            func main() {
	            fmt.Println("Demo: Rune...")

	            var rune0 rune = '♛'
	            var rune1, rune2, rune3, rune4 rune = '♠', '♧', '♡', '♬'

	            fmt.Println(rune0, rune1, rune2, rune3, rune4)

	            fmt.Printf("%c %c %c %c %c", rune0, rune1, rune2, rune3, rune4)
        } 
    
5. ## Operators in Golang
    Operators are the foundation of any programming language. Thus the functionality of the Go language is incomplete without the use of operators. Operators allow us to perform different kinds of operations on operands. In the Go language, operators Can be categorized based on their different functionality:

    1. Arithmetic Operators
    2. Relational Operators
    3. Logical Operators
    4. Bitwise Operators
    5. Assignment Operators
    6. Misc Operators

    ### Arithmetic Operators
    These are used to perform arithmetic/mathematical operations on operands in Go language: 

    1. Addition: The ‘+’ operator adds two operands. For example, x+y.
    2. Subtraction: The ‘-‘ operator subtracts two operands. For example, x-y.
    3. Multiplication: The ‘*’ operator multiplies two operands. For example, x*y.
    4. Division: The ‘/’ operator divides the first operand by the second. For example, x/y.
    5. Modulus: The ‘%’ operator returns the remainder when the first operand is divided by the second. For example, x%y.

        package main

        import "fmt"

        func main() {
	        fmt.Println("Demo: Operators...")

	        var a int = 500

	        var b int = 100

	        fmt.Println("Value of a is ", a, "\nValue of b is ", b)

	        var result = a + b

	        fmt.Println("Result of a, b addition is ", result)

	        result = a - b

	        fmt.Println("Result of a, b substraction is ", result)

	        result = a * b

	        fmt.Println("Result of a, b multiplication is ", result)

	        result = a / b

	        fmt.Println("Result of a, b division is ", result)

	        result = a % b

	        fmt.Println("Result of a, b remainder is ", result)

        }

    ###  Relational Operators
    Relational operators are used for the comparison of two values. Let’s see them one by one:

    1. ‘=='(Equal To) operator checks whether the two given operands are equal or not. If so, it returns true. Otherwise, it returns false. For example, 5==5 will return true.
    2. ‘!='(Not Equal To) operator checks whether the two given operands are equal or not. If not, it returns true. Otherwise, it returns false. It is the exact boolean complement of the ‘==’ operator. For example, 5!=5 will return false.
    3. ‘>'(Greater Than)operator checks whether the first operand is greater than the second operand. If so, it returns true. Otherwise, it returns false. For example, 6>5 will return true.
    4. ‘<‘(Less Than)operator checks whether the first operand is lesser than the second operand. If so, it returns true. Otherwise, it returns false. For example, 6<5 will return false.
    5. ‘>='(Greater Than Equal To)operator checks whether the first operand is greater than or equal to the second operand. If so, it returns true. Otherwise, it returns false. For example, 5>=5 will return true.
    6. ‘<='(Less Than Equal To)operator checks whether the first operand is lesser than or equal to the second operand. If so, it returns true. Otherwise, it returns false. For example, 5<=5 will also return true.

    Example:
        package main

        import "fmt"

        func main() {
	        fmt.Println("Demo relational operators...")

	        var a = 50

	        var b = 60

	        fmt.Println("Value of a is ", a, "\nValue of b is ", b)

	        var result = a == b

	        fmt.Println("Result of  a == b is", result)

	        result = a != b

	        fmt.Println("Result of  a != b is", result)

	        result = a > b

	        fmt.Println("Result of  a > b is", result)

	        result = a < b

	        fmt.Println("Result of  a < b is", result)

	        result = a >= b

	        fmt.Println("Result of  a >= b is", result)

	        result = a <= b

	        fmt.Println("Result of  a <= b is", result)

        }

    ### Logical Operators

    Logical operators are used to combine two or more conditions/constraints or to complement the evaluation of the original condition in consideration.  
    1. Logical AND: The ‘&&’ operator returns true when both the conditions in consideration are satisfied. Otherwise it returns false. For example, a && b returns true when both a and b are true (i.e. non-zero).
    2. Logical OR: The ‘||’ operator returns true when one (or both) of the conditions in consideration is satisfied. Otherwise it returns false. For example, a || b returns true if one of a or b is true (i.e. non-zero). Of course, it returns true when both a and b are true.
    3. Logical NOT: The ‘!’ operator returns true the condition in consideration is not satisfied. Otherwise it returns false. For example, !a returns true if a is false, i.e. when a=0.

        package main

        import "fmt"

        func main() {

	        fmt.Println("Demo logical operators")

	        var a = 100

	        var b = 100

	        var c = 200

	        fmt.Println("Value of a is ", a, "\n Value of b is ", b, "\nValue of c is ", c)

	        var result = (a == b) && (a == c)

	        // here first condition is true and second condition is false, and operation should result in false
	        fmt.Println("Result of (a == b) && (a == c) ", result)

	        result = (a == b) || (a == c)

	        // here first condition is true and second condition is false, or operation should result in true
	        fmt.Println("Result of (a == b) || (a == c) ", result)

	        result = !((a == b) && (a == c))

	        // here first condition is true and second condition is false, and operation should result in false, however a not operation on whole expresion should result in true
	        fmt.Println("Result of !((a == b) && (a == c)) ", result)
        }


    ### Bitwise Operators
    Go Supports 6 bitwise operators, which work at bit level or used to perform bit by bit operations. below are the bitwise operators 

    1. & (bitwise AND): Takes two numbers as operands and does AND on every bit of two numbers. The result of AND is 1 only if both bits are 1.
    2. | (bitwise OR): Takes two numbers as operands and does OR on every bit of two numbers. The result of OR is 1 any of the two bits is 1.
    3. ^ (bitwise XOR): Takes two numbers as operands and does XOR on every bit of two numbers. The result of XOR is 1 if the two bits are different.
    4. << (left shift): Takes two numbers, left shifts the bits of the first operand, the second operand decides the number of places to shift.
    5. >> (right shift): Takes two numbers, right shifts the bits of the first operand, the second operand decides the number of places to shift.
    6. &^ (AND NOT): This is a bit clear operator.

    ### Assignment Operators
    Assignment operators are used to assigning a value to a variable. The left side operand of the assignment operator is a variable and right side operand of the assignment operator is a value. The value on the right side must be of the same data-type of the variable on the left side otherwise the compiler will raise an error. Different types of assignment operators are shown below:

    1. “=”(Simple Assignment): This is the simplest assignment operator. This operator is used to assign the value on the right to the variable on the left.
    2. “+=”(Add Assignment): This operator is a combination of ‘+’ and ‘=’ operators. This operator first adds the current value of the variable on left to the value on the right and then assigns the result to the variable on the left.
    3. “-=”(Subtract Assignment): This operator is a combination of ‘-‘ and ‘=’ operators. This operator first subtracts the current value of the variable on left from the value on the right and then assigns the result to the variable on the left.
    4. “*=”(Multiply Assignment): This operator is a combination of ‘*’ and ‘=’ operators. This operator first multiplies the current value of the variable on left to the value on the right and then assigns the result to the variable on the left.
    4. “/=”(Division Assignment): This operator is a combination of ‘/’ and ‘=’ operators. This operator first divides the current value of the variable on left by the value on the right and then assigns the result to the variable on the left.
    6. “%=”(Modulus Assignment): This operator is a combination of ‘%’ and ‘=’ operators. This operator first modulo the current value of the variable on left by the value on the right and then assigns the result to the variable on the left.
    7. “&=”(Bitwise AND Assignment): This operator is a combination of ‘&’ and ‘=’ operators. This operator first “Bitwise AND” the current value of the variable on the left by the value on the right and then assigns the result to the variable on the left.
    8. “^=”(Bitwise Exclusive OR): This operator is a combination of ‘^’ and ‘=’ operators. This operator first “Bitwise Exclusive OR” the current value of the variable on left by the value on the right and then assigns the result to the variable on the left.
    9. “|=”(Bitwise Inclusive OR): This operator is a combination of ‘|’ and ‘=’ operators. This operator first “Bitwise Inclusive OR” the current value of the variable on left by the value on the right and then assigns the result to the variable on the left.
    10. “<<=”(Left shift AND assignment operator): This operator is a combination of ‘<<’ and ‘=’ operators. This operator first “Left shift AND” the current value of the variable on left by the value on the right and then assigns the result to the variable on the left.
    11. “>>=”(Right shift AND assignment operator): This operator is a combination of ‘>>’ and ‘=’ operators. This operator first “Right shift AND” the current value of the variable on left by the value on the right and then assigns the result to the variable on the left.

    ### Misc Operators
    Below are the list of Misc operators, we will discuss this operators in Pointer and Channel topics.
    1. &: This operator returns the address of the variable.
    2. *: This operator provides pointer to a variable.
    3. <-:The name of this operator is receive. It is used to receive a value from the channel.

6. ## Scope of Variables
7. ## Type Casting
9. ## Short Declaration Operator(:=)
10. ## var keyword vs short declaration operator
11. ## Decision Making Statements
12. ## Loops in Golang
13. ## Loop Control Statements
14. ## Switch Statement in Go