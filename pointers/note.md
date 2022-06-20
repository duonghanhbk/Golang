## Pointers

1. What is a pointer?
    A pointer is a variable that stores the memory address of another variable.

2. Declaring pointers
    ```*T``` is the type of the pointer variable which points to a value of type T.

3. Zero value of a pointer
    The zero value of a pointer is ```nil```.

4. Creating pointers using the new function
    Go also provides a handy function ```new``` to create pointers. The new function takes a type as an argument and returns a pointer to a newly allocated zero value of the type passed as argument.

5. Dereferencing a pointer

    Dereferencing a pointer means accessing the value of the variable to which the pointer points. *a is the syntax to deference a.

6. Do not pass a pointer to an array as an argument to a function. Use slice instead.

    Although this way of passing a pointer to an array as an argument to a function and making modification to it works, it is not the idiomatic way of achieving this in Go. We have slices for this.

7. Go does not support pointer arithmetic

    Go does not support pointer arithmetic which is present in other languages like C and C++.

    ```
    package main

    func main() {
        b := [...]int{109, 110, 111}
        p := &b
        p++
    }

    The above program will throw compilation error main.go:6: invalid operation: p++ (non-numeric type *[3]int)


```

