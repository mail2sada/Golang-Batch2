# Overview
In this tutorial, we will study how to do error handling in golang. As compared to other conventional languages go doesn’t have exceptions and try-catch for handling the error.
The error handling in golang can be done in two ways

Using type which implements error interface –   it is a conventional way to represent an error and  also idiomatic
Using panic and recover
We will only be discussing the first part in this article.

Using type which implements error interface
Go’s way of dealing with an error is to explicitly return the error as a separate value.

The crux to understanding error in golang is to understand the error interface. This is how the error interface looks as defined in the builtin package

https://golang.org/pkg/builtin/#error


            type error interface {
                Error() string
            }

The error interface is the conventional way to represent an error condition in golang. If this is nil then it means no error. So in go error is treated as a value. This value implements the error interface and which can be passed around to a function, returned from a function, and which can be stored in variables.

So any type which defines the Error method is set to be implementing the error interface.  As we mentioned earlier that go doesn’t have exceptions and try-catch so an idiomatic way of handling error condition is to return the error as the last return value. The value can then be checked for nil. If it is nil then the  called function did not return the  error otherwise the called function returned the error. So let’s see a program which demonstrates what we just talked.

            package main

            import (
                "fmt"
                "os"
            )

            func main() {
                file, err := os.Open("non-existing.txt")
                if err != nil {
                    fmt.Println(err)
                } else {
                    fmt.Println(file.Name() + "opened succesfully")
                }
            }



# Different ways of creating an error
Let’s see now different methods of creating an error

## Using errors.New(“some_error_message”)
            package main

            import (
                "errors"
                "fmt"
            )

            func main() {
                sampleErr := errors.New("error occured")
                fmt.Println(sampleErr)
            }

## Using fmt.Errorf(“error is %s”, “some_error_message”).
This way creates error with formatting of error messages

            package main

            import (
                "fmt"
            )

            func main() {
                sampleErr := fmt.Errorf("Err is: %s", "database connection issue")
                fmt.Println(sampleErr)
            }


#  Creating Custom error
The below example illustrates the use of custom error. In below example

inputError is a struct that has the Error() method hence it is of type error interface
You can also add additional information to the custom error by extending its fields or by adding new methods. inputError has an additional field named missingFields and a function getMissingFields function.
We can use type assertion to convert from error to inputError
Example:

package main

import "fmt"

type inputError struct {
    message      string
    missingField string
}

func (i *inputError) Error() string {
    return i.message
}

func (i *inputError) getMissingField() string {
    return i.missingField
}

func main() {
    err := validate("", "")
    if err != nil {
        if err, ok := err.(*inputError); ok {
            fmt.Println(err)
            fmt.Printf("Missing Field is %s\n", err.getMissingField())
        }
    }
}

func validate(name, gender string) error {
    if name == "" {
        return &inputError{message: "Name is mandatory", missingField: "name"}
    }
    if gender == "" {
        return &inputError{message: "Gender is mandatory", missingField: "gender"}
    }
    return nil
}

Ignoring errors
Underscore (‘_’) operator can be used to ignore the error returned from a function call.  Before we see a program it’s important to note that error should never be ignored. It is not a recommended way. Let’s see a program

package main
import (
    "fmt"
    "os"
)
func main() {
    file, _ := os.Open("non-existing.txt")
    fmt.Println(file)
}
Output

{nil}
In the above program, we used the underscore operator to ignore the error while opening a non-existing file. That is why the file instance returned by the function is nil. Therefore it is better to check for an error before using any other argument returned by the function as that can be nil and would result in unwanted issues and also sometimes it might also result in a panic.


# Overview
In this section we will be covering the advanced topics related to error in go.

- Wrapping and un-wrapping errors
- Error comparison 
- Extract underlying type from error
- As and Is function of errors package
- Please refer to link below first which starts with basics of error in go.

# Wrapping of error 
In go, error can wrap another error as well. 

Wrapping of errors means to create a hierarchy of errors in which a  particular instance of error wraps another error and that particular instance itself can be wrapped inside another error.  Below is the syntax for wrapping an error

e := fmt.Errorf("... %w ...", ..., err, ...)
%w directive Is used for wrapping the error.  The fmt.Errorf should be called with only one %w directive. Let’s see an example.

            package main

            import (
                "fmt"
            )

            type errorOne struct{}

            func (e errorOne) Error() string {
                return "Error One happended"
            }

            func main() {

                e1 := errorOne{}

                e2 := fmt.Errorf("E2: %w", e1)

                e3 := fmt.Errorf("E3: %w", e2)

                fmt.Println(e2)

                fmt.Println(e3)

            }

In the above program, we created a struct errorOne that has an Error method hence it implements the error interface. Then we created an instance of the errorOne struct named e1. Then we wrapped that instance e1 into another error e2 like this

            e2 := fmt.Errorf("E2: %w", e1)
Then we wrapped e2 into e3 like below. 

            e3 := fmt.Errorf("E3: %w", e2)
So so we created a hierarchy of errors in which e3 wraps e2 and e2 wraps e1.  Thus e3 also wraps e1 transitively. When we print e2  it also prints the error from e1 and gives the output.

            E2: Error One happended
When we print e3 it prints the error from e2 as well as e1 and gives the output.

            E3: E2: Error One happended
Now the question which comes to the mind that whats the use case of wrapping the errors. To understand it let’s see an example

            package main

            import (
                "fmt"
            )

            type notPositive struct {
                num int
            }

            func (e notPositive) Error() string {
                return fmt.Sprintf("checkPositive: Given number %d is not a positive number", e.num)
            }

            type notEven struct {
                num int
            }

            func (e notEven) Error() string {
                return fmt.Sprintf("checkEven: Given number %d is not an even number", e.num)
            }

            func checkPositive(num int) error {
                if num < 0 {
                    return notPositive{num: num}
                }
                return nil
            }

            func checkEven(num int) error {
                if num%2 == 1 {
                    return notEven{num: num}
                }
                return nil
            }

            func checkPostiveAndEven(num int) error {
                if num > 100 {
                    return fmt.Errorf("checkPostiveAndEven: Number %d is greater than 100", num)
                }

                err := checkPositive(num)
                if err != nil {
                    return err
                }

                err = checkEven(num)
                if err != nil {
                    return err
                }

                return nil
            }

            func main() {
                num := 3
                err := checkPostiveAndEven(num)
                if err != nil {
                    fmt.Println(err)
                } else {
                    fmt.Println("Givennnumber is positive and even")
                }

            }
Output

checkEven: Given number 3 is not an even number
In the above program, we have a function checkPostiveAndEven that checks whether a number is even and positive. In turns, it calls the checkEven function to check if the number is even. And then it calls checkPositive function to check if the number is positive. If a number is not even and positive it an error is raised.

In the above program it is impossible to tell stack trace of the error. We know that this error came from checkEven function for the above output. But which function called the checkEven function is not clear from the error. This is where wrapping the error comes in the picture.  This becomes more useful when the project is big and there are a lot of functions calling each other.  Let’s rewrite the program by wrapping the error.

package main

import (
	"fmt"
)

type notPositive struct {
	num int
}

func (e notPositive) Error() string {
	return fmt.Sprintf("checkPositive: Given number %d is not a positive number", e.num)
}

type notEven struct {
	num int
}

func (e notEven) Error() string {
	return fmt.Sprintf("checkEven: Given number %d is not an even number", e.num)
}

func checkPositive(num int) error {
	if num < 0 {
		return notPositive{num: num}
	}
	return nil
}

func checkEven(num int) error {
	if num%2 == 1 {
		return notEven{num: num}
	}
	return nil
}

func checkPostiveAndEven(num int) error {
	if num > 100 {
		return fmt.Errorf("checkPostiveAndEven: Number %d is greater than 100", num)
	}

	err := checkPositive(num)
	if err != nil {
		return fmt.Errorf("checkPostiveAndEven: %w", err)
	}

	err = checkEven(num)
	if err != nil {
		return fmt.Errorf("checkPostiveAndEven: %w", err)
	}

	return nil
}

func main() {
	num := 3
	err := checkPostiveAndEven(num)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Given number is positive and even")
	}

}
Output

checkPostiveAndEven: checkEven: Given number 3 is not an even number
 The above program is same as the previous program just that in the checkPostiveAndEven function , we wrap the errors like below.

fmt.Errorf("checkPostiveAndEven: %w", err)
So the output is more clear and the error is more informative. The output clearly mentions the sequence of calling as well

checkPostiveAndEven: checkEven: Given number 3 is not an even number

