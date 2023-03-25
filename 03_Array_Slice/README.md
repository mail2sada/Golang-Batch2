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
    len function is go built-in function, len function returns the number of items in the string.
    func len(v Type) int

    Takes any type as argument and returns the number of items

        Syntax:
            size := len(collection)

        Example:
            package main

            import "fmt"

            func main() {

                fmt.Println("Demo: len function on collection")

                var collectionInt [10]int = [10]int{1, 2, 3, 4, 5}

                fmt.Println("Length of collectionInt is:", len(collectionInt))
            }

    6. loop through an array
    There are multiple ways of iterate over a array, we can iterate using for loop and range operator

    Below are the examples to iterate over the loop

        Example 1:
            package main

            import "fmt"

            func main() {

                fmt.Println("Demo: Iterating over a collection")

                var collectionInt = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

                for i := 0; i < len(collectionInt); i++ {
                    fmt.Printf("collectionInt[%d]=%d\n", i, collectionInt[i])
                }
            }
        Example 2:
            package main

            import "fmt"

            func main() {
                fmt.Println("Demo: iterating loop using range")

                var strArray = [10]string{"Good", "Morning", "How", "Are", "You", "Wonderful", 9: "Amazing"}

                for idx, val := range strArray {
                    fmt.Println("strArray @Index:", idx, " is:", val)
                }
            }

2. Copying an Array into Another Array

    Array in go are of value type, when you assign an array to another variable, it creates a copy of array. When you pass array to a function it creates copy to a array, any changes made to the copied array will not impact original array

        Syntax:
            var array [size]type = [size]type {val1, val2, val3, } 

            var arrayCopy = array

        Example:
            package main

            import "fmt"

            func main() {

                fmt.Println("Demo copying array...")

                var array [5]string = [5]string{"Golang", "is", "a", "wonderful", "language.."}

                var copyArray = array

                fmt.Println("Contents of array are:", array)

                fmt.Println("Contents of copyArray are:", copyArray)

                copyArray[3] = "I have changed this..."

                fmt.Println("Printing array and copyArray after changing copyArray[3]=\"I have changed this...\"")

                fmt.Println("Contents of array are:", array)

                fmt.Println("Contents of copyArray:", copyArray)
            }

3. Multidimensional Array

    Golang supports multi-dimensional arrays, 

        Syntax:
            var multiDimensionArray [size1][size2]type = [size1][size2]type {}
        
        Example:
            package main

            import "fmt"

            func main() {
                fmt.Println("Demo: Multi-dimentsional array")

                var multiDimensionalArray [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

                for i := 0; i < 3; i++ {
                    for j := 0; j < 3; j++ {
                        fmt.Printf(" %d ", multiDimensionalArray[i][j])
                    }
                    fmt.Println()
                }

            }


4. Filter Array Elements
    We can filter out specific number of elements from array, using (:). We will explore different ways of using (:)

        Examples:
            package main

            import "fmt"

            func main() {

                fmt.Println("Demo: Filtering elements of array..")

                var array [10]int = [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

                fmt.Println("Filering out first 5 elements", array[:5])

                fmt.Println("Fileterning out last 5 elements", array[5:])
                fmt.Println("Filterning out from index 3 to 7", array[3:8])
            }

6. Comparison of Arrays.

        Exmaple:
            package main

            import "fmt"

            func main() {
                fmt.Println("Demo: comparision of arrays")

                var stringArray = [5]string{"Hello", "how ", "are", "you", "?"}

                fmt.Println("Contents of stringArray:", stringArray)

                fmt.Println("Copying array...")
                copyArray := stringArray

                if copyArray == stringArray {
                    fmt.Println("Both arrays are equal..")
                }

                if copyArray != stringArray {
                    fmt.Println("Both arrays are unequal..")
                }

                fmt.Println("Changing contents of array..")

                stringArray[4] = "Wow"

                if copyArray == stringArray {
                    fmt.Println("Both arrays are equal..")
                }

                if copyArray != stringArray {
                    fmt.Println("Both arrays are unequal..")
                }

            }

5. Slices in Golang
    A slice is a flexible and extensible data structure to implement and manage collections of data. Slices are made up of multiple elements of same type. A slice is a segment of dynamic arrays that can grow and shrink. Like arrays, slices are index-able and have a length and capacity property. 

    Slice contains 3 components,
    1. Pointer: The pointer is used to points to the first element of the array that is accessible through the slice. 
    2. Length: The length is the total number of elements present in the array.
    3. Capacity: The capacity represents the maximum size upto which it can expand.

    Slices are of reference type, data is accessible with the name of slice followec by index in []. Length represents the length of the slice and it is retrieved using len function. Capacity indicates the current capacity of the slice, when new elements are appended in array up to full capacity then slice data will be reallocated to a new continues memory and capacity will be recalculated. Calculating this capacity is internal to go.

    ### Declaring Slice

        Syntax:
            // Only difference between array and slice declaration is size, while declaring slice we won't mention slice.
            var intSlice []int = []int {}

            var sliceString []string = make([]string, 10)
        Exmaple:

            package main

            import "fmt"

            func main() {

                fmt.Println("Demo: Slice Declaration...")

                var intSlice []int = []int{1, 2, 3}

                intSlice = append(intSlice, 4, 5, 6, 7, 8, 9, 10)

                fmt.Println("Length of the slice:", len(intSlice))
                fmt.Println("Capacity of the slice:", cap(intSlice))

                fmt.Println("Accessing element at index [3]:", intSlice[3])

                fmt.Println("Iterating slice using len")
                for i := 0; i < len(intSlice); i++ {
                    fmt.Println("Content at index [", i, "]  = ", intSlice[i])
                }

                fmt.Println("Iterating slice using range operator")
                for idx, val := range intSlice {
                    fmt.Println("Content at index [", idx, "]  = ", val)

                }

            }
        Example2:

            package main

            import "fmt"

            func main() {

                fmt.Println("Demo: Declaration of slice using make...")

                var myStringSlice []string = make([]string, 5)

                fmt.Println("Length of myStringSlice:", len(myStringSlice))
                fmt.Println("Capacity of myStringSlice:", cap(myStringSlice))

                //populating list with few elements.

                myStringSlice = append(myStringSlice, "Hello", "how", "are", "you")

                fmt.Println("Length of myStringSlice:", len(myStringSlice))
                fmt.Println("Capacity of myStringSlice:", cap(myStringSlice))

                fmt.Println("Iterating through the slice using len")

                for i := 0; i < len(myStringSlice); i++ {

                    fmt.Println("Contents of myStringSlice[", i, " ]", myStringSlice[i])
                }

                fmt.Println("Iterating through the slice using range")

                for idx, val := range myStringSlice {
                    fmt.Println("Contents of myStringSlice[", idx, " ]", val)

                }

            }

    ### make 

    make is a Go function, that enables you to create slice, map, channel, etc. Data types will be allocated in heap.
    func make(t Type, size ...IntegerType) Type

    #### Parameters:
    1. t Type: The type that is allocated and for which the reference will be returned. Example: map, slice, etc.
    2. size: The size of the container.
    3. capacity: The total capacity that will be allocated. capacity must be greater than or equal to size.

    #### Return value
    make() returns a reference to the map, slice, or channel that is allocated on the memory.

        syntax:
            var data = make([]type, length, capacity)
            //capacity is a optional parameter.
            package main

            import "fmt"

            func main() {
                fmt.Println("Demo: Usage of make in slice...")

                // statement below will allocate a string slice of lenght 5
                var stringSlice = make([]string, 5)

                // here both capacity and length should be same

                fmt.Println("Length of stringSlice:", len(stringSlice))
                fmt.Println("Capacity of stringSlice:", cap(stringSlice))

                fmt.Println("\nWe are re allocating stringSlice using length and capacity")

                stringSlice = make([]string, 5, 10)
                fmt.Println("Length of stringSlice:", len(stringSlice))
                fmt.Println("Capacity of stringSlice:", cap(stringSlice))
            }

6. Composite Literal

    Composite literal is a literal which is usied to initialize composite type. Go supports a range of composite types such as arrays, slice, structs, maps, channel, etc. A literal defining composite type are known as composite literals.

    Example
            var testArray = [10]int{1,2,3,4,5,6,7,8,9,10}

            // {1,2,3,4,5,6,7,8,9,10} is a composite literal.

            val multiDimensionArray [3][3] int = {{1,2,3},
                                                  {4,5,6},
                                                  {7,8,9}}
            // {{1,2,3},
            // {4,5,6}
            // {7,8,9}}

7. Functions to handle slice.
    Go built-in function suit provide below methods for slice handling.
    len() // returns the lengh of composite type, works on composite type slice, array, chan, map
    cap() // return the the capacity of slice
    make() // function allocates the composite type slice, chan, map
    append() //function enables you to add new elements in slice.

7. Appending, Updating and Deleting contents of slice


8. Copying one Slice into another Slice
9. Comparing two Slices in Golang
10. Checking the Equality of Slices in Golang
12. Sorting a Slice in Golang
12. Trimming a Slice in Golang
13. Splitting a Slice in Golang
14. Slice Sort, 
15. Reverse, 
16. Search Functions