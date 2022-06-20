## Methods

1. Introduction

    A method is just a function with a special receiver type between the func keyword and the method name. The receiver can either be a struct type or non-struct type.

    ```
    func (t Type) methodName(parameter list) {}
    ```

2. Methods vs Functions

    Methods

    ```
    func (e Employee) displaySalary() {
        fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
    }
    ```

    Function

    ```
    func displaySalary(e Employee) {
        fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
    }
    ```

    So why do we have methods when we can write the same program using functions. There are a couple of reasons for this. Let's look at them one by one.

    - Go is not a pure object-oriented programming language and it does not support classes. Hence methods on types are a way to achieve behaviour similar to classes. Methods allow a logical grouping of behaviour related to a type similar to classes. In the above sample program, all behaviours related to the Employee type can be grouped by creating methods using Employee receiver type. For example, we can add methods like calculatePension, calculateLeaves and so on.

    - Methods with the same name can be defined on different types whereas functions with the same names are not allowed. Let's assume that we have a Square and Circle structure. It's possible to define a method named Area on both Square and Circle. This is done in the program below.

    - When a function has a value argument, it will accept only a value argument.

    - When a function has a pointer argument, it will accept only a pointer argument.

    - When a method has a value receiver or pointer receiver, it will accept both pointer and value receivers.

3. When to use pointer receiver and when to use value receiver

    Generally, pointer receivers can be used when changes made to the receiver inside the method should be visible to the caller.

    Pointers receivers can also be used in places where it's expensive to copy a data structure. Consider a struct that has many fields. Using this struct as a value receiver in a method will need the entire struct to be copied which will be expensive. In this case, if a pointer receiver is used, the struct will not be copied and only a pointer to it will be used in the method.

4. Methods with non-struct receivers

    To define a method on a type, the definition of the receiver type and the definition of the method should be present in the same package.