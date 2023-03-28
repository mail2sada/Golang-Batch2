# Strings

## String Concatenation
### Using the + operator
We can simply concatenate strings using the + operator, though keep in mind you should only concatenate the string with a string and not any other data type like integer, or float, it will throw out errors for mismatched string types.

            package main

            import (
                "fmt"
            )
            func main() {
                s1 := "Go"
                s2 := "Programming"
                s3 := "Language"

                s := s1 + s2 + s3
                fmt.Println(s)
            }

## Using the += operator

The other way to continuously append a string to an existing string, we can use the += operator. This operator will append the provided string to the end of the original string.

            package main

            import (
                "fmt"
            )
            func main() {
                p := "Go"
                p += "Programming"
                fmt.Println(p)
            }
## Using the Join method

The join method is a function available in the string package in Golang. We can join strings elements in a slice or an array using the Join method in the strings package in golang. The Join method will combine all the elements in between the elements with a particular string. So, the function takes two parameters Join(array, string), the array or a slice is parsed into the function which will be used to insert the provided string in between the elements of the slice.

Join
Parameters : array/slice, string
Return Value : string

            package main        

            import (
                "fmt"
                "strings"
            )
            func main() {
                q := []string{"youtube.com", "channel", "golang", "string"}
                r := strings.Join(q, "/")
                fmt.Println(r)
            }
## Using Sprintf method
We can use the Sprintf function from the fmt package to format the string by storing the string rather than printing it to the console. The sprintf function is quite similar to the Printf but it only parses strings rather than printing them directly to the console.

            package main
            import (
                "fmt"
            )
            func main() {
                name := "Ajay Khanna"
                domain := "Mavenir"
                service := "Engineer"

                email := fmt.Sprintf("%s.%s@%s.com", service, name, domain)
                fmt.Println(email)
            }

## Using Go string Builder method

The Builder type is provided by the strings package in Golang. The Builder type helps in building strings in an efficient way. By creating a string builder object, we can perform operations on a String.

import (
    "fmt"
    "strings"
)

            func main() {
            // Using Builder function

            c := []string{"j", "a", "v", "a"}
            var builder strings.Builder
            for _, item := range c {
                builder.WriteString(item)
            }
            fmt.Println("builder = ", builder.String())
            b := []byte{'s', 'c', 'r', 'i', 'p', 't'}
            for _, item := range b {
                builder.WriteByte(item)
            }
            fmt.Println("builder = ", builder.String())
            builder.WriteRune('s')
            fmt.Println("builder = ", builder.String())
            fmt.Println("builder = ", builder)
            }

## Using the Bytes buffer method

The bytes package also has something similar to Builder type in string as Buffer. It has almost the same set of methods and properties. The main difference is the efficiency, strings.Builder is comparatively faster than the bytes.Buffer type due to several low-level implementations. We can discuss those fine details in a separate article but right now we'll focus on the ways we can utilize this type for string concatenation.

            package main

            import (
                "fmt"
                "bytes"
            )

            func main() {
                // Using bytes buffer method

                var buf bytes.Buffer

                for i := 0; i < 2; i++ {
                    buf.WriteString("ja")
                }
                fmt.Println(buf.String())

                buf.WriteByte('r')

                fmt.Println(buf.String())

                k := []rune{'s', 'c', 'r', 'i', 'p', 't'}
                for _, item := range k {
                    buf.WriteRune(item)
                }
                fmt.Println(buf.String())
            }

## String Comparison

### Using Comparison operators

The basic comparison can be done with the comparison operators provided by Golang. Just like we compare numeric data we can compare strings. Though the factor with which we compare them is different. We compare them by the lexical order of the string characters.

            package main

            import "fmt"

            func main() {
                s1 := "gopher"
                s2 := "Gopher"
                s3 := "gopher"

                isEqual := s1 == s2

            //"gopher" is not same as "Gopher" and hence `false`
                fmt.Printf("S1 and S2 are Equal? %t \n", isEqual)
                fmt.Println(s1 == s2)

            // "gopher" is not equal to "Gopher" and hence `true`
                fmt.Println(s1 != s2)

            // "Gopher" comes first lexicographically than "gopher" so return true 
            // G -> 71 in ASCII and g -> 103
                fmt.Println(s2 < s3)
                fmt.Println(s2 <= s3)

            // "Gopher" is not greater than "gopher" as `G` comes first in ASCII table
            // So G value is less than g i.e. 71 > 103 which is false
                fmt.Println(s2 > s3)
                fmt.Println(s2 >= s3)

            }

### Using Compare method
We also have the Compare method in the strings package for comparing two strings. The comparison method returns an integer value of either -1, 0, or 1. If the two strings are equal, it will return 0. Else if the first string is lexicographically smaller than the second string, it will return -1, else it will return +1.

strings.Compare
Return Type: Int (-1, 0, 1)
Parameters: string, string

            package main

            import(
            "fmt"
            "strings"
            )

            func main() {
                s1 := "gopher"
                s2 := "Gopher"
                s3 := "gopher"

                fmt.Println(strings.Compare(s1, s2))
                fmt.Println(strings.Compare(s1, s3))
                fmt.Println(strings.Compare(s2, s3))
            }


### Using strings EqualFold

We also have another method in the strings library called EqualFold which compares two strings lexicographically but without considering the case precedence. That is the upper case or lower case is ignored and considered equal. So we are truly case-insensitively comparing the strings.

strings.EqualFold
Return Type: bool (true or false)
Parameters: string, string

            package main

            import(
            "fmt"
            "strings"
            )

            func main() {

                s1 := "gopher"
                s2 := "Gopher"
                s3 := "gophy"

                fmt.Println(strings.EqualFold(s1, s2))
                fmt.Println(strings.EqualFold(s1, s3))
                fmt.Println(strings.EqualFold(s2, s3))
            }
## Using DeepEqual

package main

            import(
            "fmt"
            "strings"
            )

            func main() {

                s1 := "gopher"
                s2 := "Gopher"

                fmt.Println(reflect.DeepEqual(s1, s2))
            }

## String Manipulation and utility methods

Important packages to explore

    "golang.org/x/text/cases"
    "golang.org/x/text/language"


            strings.ToLower
            Return Type: string
            Parameters: string
<br/>
<br/>
<br/>

            strings.ToUpper
            Return Type: string
            Parameters: string

### Contains and ContainsAny functions

Contains and ContainsAny functions
In the strings package, we have the Contains and ContainsAny method which checks for the presence of substrings within a string. This will help in looking up hidden patterns and find for substrings in a string.

strings.Contains
Parameters: string, string
Return Type: bool (true, false)
The Contains method helps in getting the exact match of the substring in the given string. Firstly, the method takes two parameters one being the actual string and the other being the substring that you want to find inside the string. Let's say we have the string="bootcamp" and substr="camp", then the Contains method will return true because the string contains the substring camp.

strings.ContainsAny
Parameters: string, string
Return Type: bool (true, false)
The ContainsAny method just like the Contains method takes two parameters string and the other as the substring, but it would return true if it finds any character or a single byte(Unicode chars) inside the string. Let's say the string="google photos" and substring="abcde", then the ContainsAny method will return true because the string contains at least one character in the substring which is e.




