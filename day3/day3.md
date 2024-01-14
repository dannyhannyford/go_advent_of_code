## pointers and values
use pointer when the struct is really big, like http.Request

## rule of thumb
structs need pointers

structs are composite data types that groups variables under a single name

interfaces do not need pointers

interfaces are sets of method signatures that a type must implement



`&` address operator
    get the address of a variable
    returns a pointer to the memory location where the variable is stored
    in short: create a pointer from a value
    ex

    ```
    x := 5
    p := &x // p is a pointer to int, storing the address of x
    ```



`*` dereference operator
    access or modify the value at a memory address.
    It is called dereferencing a pointer

    ex, continued from above

    ```
    val := *p // val is now t, the value at the memory address stored in p
    ```

Go allows to avoid explicit dereference.


# pointers vs values for receivers

value methods can be invoked on pointers and values, but pointer methods can only be invoked on pointers

pointer methods can modify the reciever

when a value is addressable the language takes care ofthe common case of invoking a pointer method on a value by inserting the address operator automatically
----------------------------------------------
package main

import "fmt"

// Function that takes a pointer to an integer and modifies its value
func modifyValue(num *int) {
    *num = 10 // Dereference the pointer and assign a new value
}

func main() {
    value := 5
    fmt.Println("Before:", value) // Output: Before: 5

    modifyValue(&value) // Pass the address of 'value'

    fmt.Println("After:", value) // Output: After: 10
}

------------------------------------------------
When working with pointers, always consider the possibility of nil and handle it appropriately to avoid runtime panics.
In concurrent programming, be cautious with pointers and consider synchronization techniques or immutable data structures to avoid race conditions.

### you must dereference before indexing
(*matrix)[0][4]