``There is a subtlety to be considered when using fallthrough. 
Fallthrough will happen even when the case evaluates to false``

```func main() {
    switch num := 25; {
    case num < 50:
        fmt.Printf("%d is lesser than 50\n", num)
        fallthrough
    case num > 100:
        fmt.Printf("%d is greater than 100\n", num)     
    }
}```

The program above will print: 

``
- 25 is lesser than 50
- 25 is greater than 100
``