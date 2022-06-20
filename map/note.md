## Golang Maps Tutorial
1. What is a map?

    A map is a builtin type in Go that is used to store key-value pairs.

2. How to create a map?
    ```
    make(map[type of key]type of value)
    ```

3. Zero value of a map
    The zero value of a map is nil. If you try to add elements to a nil map, a run-time panic will occur. Hence the map has to be initialized before adding     elements.

    ```
    func main() {
        var employeeSalary map[string]int
        employeeSalary["steve"] = 12000
    }
    ```
    ``panic: assignment to entry in nil map``

4. Checking if a key exists
    ```
    value, ok := map[key]
    ```

5. Deleting items from a map
    ```delete(map, key)``` is the syntax to delete ```key``` from a ```map```. The delete function does not return any value.

6. Maps are reference types

    Similar to slices, maps are reference types.


7. Maps equality

    Maps can't be compared using the == operator. The == can be only used to check if a map is nil.

    One way to check whether two maps are equal is to compare each one's individual elements one by one. The other way is using reflection. I would encourage you to write a program for this and make it work :).