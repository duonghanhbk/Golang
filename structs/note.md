## Structs

1. What is a struct?

    A struct is a user-defined type that represents a collection of fields. It can be used in places where it makes sense to group the data into a single unit rather than having each of them as separate values.

2. Zero value of a struct

    When a struct is defined and it is not explicitly initialized with any value, the fields of the struct are assigned their zero values by default.

    It is also possible to specify values for some fields and ignore the rest.

3. Pointers to a struct

    It is also possible to create pointers to a struct.

4. Promoted fields

    Fields that belong to an anonymous struct field in a struct are called promoted fields since they can be accessed as if they belong to the struct which holds the anonymous struct field.

5. Exported structs and fields

    If a struct type starts with a capital letter, then it is an exported type and it can be accessed from other packages. Similarly, if the fields of a struct start with caps, they can be accessed from other packages.

6. Structs Equality

    Structs are value types and are comparable if each of their fields are comparable. Two struct variables are considered equal if their corresponding fields are equal.

