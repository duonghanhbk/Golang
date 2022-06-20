## Strings

1. What is a String?

    A string is a slice of bytes in Go. Strings can be created by enclosing a set of characters inside double quotes " ".

2. Strings are immutable

    To workaround this string immutability, strings are converted to a slice of runes.

    ```
    package main

    import (
        "fmt"
    )

    func mutate(s []rune) string {
        s[0] = 'a'
        return string(s)
    }
    func main() {
        h := "hello"
        fmt.Println(mutate([]rune(h)))
    }
    ```