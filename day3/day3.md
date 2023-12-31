## pointers and values
use pointer when the struct is really big, like http.Request

## rule of thumb
structs need pointers

interfaces do not need pointers



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
    access the value at a memory address.
    It is called dereferencing a pointer

    ex, continued from above

    ```
    val := *p // val is now t, the value at the memory address sotred in p
    ```