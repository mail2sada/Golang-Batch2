# Interfaces
1. Overview
The interface is a custom type that is used to specify a set of one or more method signatures and the interface is abstract, so you are not allowed to create an instance of the interface. But you are allowed to create a variable of an interface type and this variable can be assigned with a concrete type value that has the methods the interface requires. Or in other words, the interface is a collection of methods as well as it is a custom type.

            type interface_name interface{
                // Method signatures
            }

            // Creating an interface
            type myinterface interface{

                // Methods
                fun1() int
                fun2() float64
            }
Here the interface name is enclosed between the type and interface keywords and the method signatures enclosed in between the curly braces.

In the Go language, it is necessary to implement all the methods declared in the interface for implementing an interface. The go language interfaces are implemented implicitly. And it does not contain any specific keyword to implement an interface just like other languages.

The zero value of the interface is nil.
When an interface contains zero methods, such types of interface is known as the empty interface. So, all the types implement the empty interface.

            Syntax:

                interface{}

### Interface Types: 
The interface is of two types 
-   #### static 
    The static type is the interface itself. However interface does not have a static value so it always points to the dynamic values.
-   #### dynamic type
The static type is the interface itself, for example, tank in the below example. But interface does not have a static value so it always points to the dynamic values.
A variable of the interface type containing the value of the Type which implements the interface, so the value of that Type is known as dynamic value and the type is the dynamic type. It is also known as concrete value and concrete type.


2. Multiple Interfaces
4. Embedding Interfaces
5. Polymorphism Using Interfaces
6. Methods in connection with interfaces