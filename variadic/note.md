## Variadic params

Then why do we even need variadic functions when we can achieve the same functionality using slices as below:
```
func find(num int, nums []int) {
    fmt.Printf("type of nums is %T\n", nums)
    found := false
    for i, v := range nums {
        if v == num {
        fmt.Println(num, "found at index", i, "in", nums)
        found = true
        }
    }
    if !found {
        fmt.Println(num, "not found in ", nums)
    }
    fmt.Printf("\n")
}
```

The following of the advantages of using variadic arguments instead of slices.
- There is no need to create a slice during each function call.
- In line no.25 of the program above, we are creating an empty slice just to satisfy the signature of the find function. This is totally not needed in the case of variadic functions. This line can just be find(87) when variadic function is used.
- I personally feel that the program with variadic functions is more readable than the once with slices :)

``So is there a way to pass a slice to a variadic function? The answer is yes.``
There is a syntactic sugar which can be used to pass a slice to a variadic function. You have to suffix the slice with ellipsis ```...``` If that is done, the slice is directly passed to the function without a new slice being created.