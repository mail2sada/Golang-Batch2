### Array and Slice

1. Arrays in Golang
    An array is a data structure that consists of multuple values of homogeneous data types. In other words we can call it a special variable that can hold more than one values of same type. The values in the array are called elements or items. Array holds the specific number of elements, it can not grow or shrink. Different Data types can be handled in arrays such as Integer, string, boolean, or any other user defined data type. Index of array starts with 0 and last element is (size -1).  

    1. Declaring array
        An array can be declared using a var keyword or short declaration operator(:=).

        Syntax:
            var array [size]type = [size]type {}
        Example 1:
            package main

            import "fmt"

            func main() {
                fmt.Println("Demo: Arrays....")

                // integer array that can hold 3 elements

                var intArray [3]int = [3]int{}

                intArray[0] = 1
                intArray[1] = 2
                intArray[2] = 3

                fmt.Println(intArray)

            }
        Syntax: 
            var array = [size]type {}
        Example 2:
            package main

            import "fmt"

            func main() {

                var intArray = [3]int{}

                intArray[0] = 10
                intArray[1] = 20
                intArray[2] = 30

                fmt.Println("Contents of int Array", intArray)
            }
        Syntax:
            array := [size]type {}
        Example 3:
            package main

            import "fmt"

            func main() {
                fmt.Println("Demo: Array declaration")
                array := [3]int{}

                array[0] = 100
                array[1] = 200
                array[2] = 300

                fmt.Println("Content of array are:", array)
            }

    2. Assigning and accessing elements
        Array elements can be assigned with array-name suffixed by index in []
        Syntax:
            var array [size]type = [size]type{}
            array[idx] = expression
            val = array[idx]
        Example:
            package main

            import "fmt"

            func main() {
                fmt.Println("Demo: Assigining and accessing arrays")
                var color [5]string = [5]string{}

                // Assigining array elements
                color[0] = "Red"
                color[1] = "Green"
                color[2] = "Blue"
                color[3] = "Orrange"
                color[4] = "Purple"

                fmt.Println("Contents of color array:", color)

                // Accessing array elements...

                fmt.Println("color[0]:", color[0])
                fmt.Println("color[1]:", color[1])
                fmt.Println("color[2]:", color[2])
                fmt.Println("color[3]:", color[3])
                fmt.Println("color[4]:", color[4])

            }

    3. Initializing array
        Array can be intialise at the time of declaration

        syntax:
            var array [size]type = [size]type{val1, val2...size times}
        Example.

            package main

            import "fmt"

            func main() {

                fmt.Println("Demo: Array Initialisation 1...")

                var intArray [3]int = [3]int{1, 2, 3}

                fmt.Println("Contents of intArray are:", intArray)

                var strArray = [3]string{"Hello ", "Hi ", "How are you? "}

                fmt.Println("Contents of str Array", strArray)

                uintArray := [3]uint8{10, 20, 30}

                fmt.Println("Contents of uintArray", uintArray)
            }

    4. Initialising specific elements in array
    We can intialize the specific elements of an array.

        Syntax:
            var array[size]type = [size]type{idx1:val1, idx2:val2, idx3:, val3,... size}
        
        Example:

            package main

            import "fmt"

            func main() {
                fmt.Println("Demo: Initializing specific elements of array")

                // we will declare an array of int with size 10 and intializing index 0, 5, 9

                var integerArray = [10]int{0: 100, 5: 200, 9: 300}

                fmt.Println("Contents of array are ", integerArray)
            }
        
    5. Initialising array with ellipses
        While declaring array if we are sure of the number of elements it will have we can use ellipses
        Syntax:

            var array = [...]type {val1, val2, val3...}
        
        Example:
            package main

            import "fmt"

            func main() {

                fmt.Println("Demo: Declaration and initialisation with elipsis")

                var arrayInt = [...]int{10, 20, 30, 40, 50}

                fmt.Println("Contents of arrayInt:", arrayInt)

                strArray := [...]string{"Red", "Green", "Blue"}

                fmt.Println("Contents of strArray:", strArray)

                strNewArray := [...]string{1: "Hello", 5: "Namaste", 10: "How are you?"}

                fmt.Println("Contents of strNewArray:", strNewArray)
            }



    6. len function
    6. loop through an array
2. Copying an Array into Another Array
    1. Copying Array by value
    2. Copuing Array by reference.

3. Filter Array Elements
4. Multidimensional Array

5. Slices in Golang
6. Slice Composite Literal
7. Copying one Slice into another Slice
8. Comparing two Slices in Golang
9. Checking the Equality of Slices in Golang
10. Sorting a Slice in Golang
11. Trimming a Slice in Golang
12. Splitting a Slice in Golang
13. Slice Sort, 
14. Reverse, 
15. Search Functions