## Interfaces

1. What is an interface?

    In Go, an interface is a set of method signatures. When a type provides definition for all the methods in the interface, it is said to implement the interface.

2. Interface internal representation

    An interface can be thought of as being represented internally by a tuple (type, value). type is the underlying concrete type of the interface and value holds the value of the concrete type.

3. Empty interface

    An interface that has zero methods is called an empty interface. It is represented as interface{}.

4. Type Switch

    A type switch is used to compare the concrete type of an interface against multiple types specified in various case statements. It is similar to switch case. The only difference being the cases specify types and not values as in normal switch.

    It is also possible to compare a type to an interface. If we have a type and if that type implements an interface, it is possible to compare this type with the interface it implements.

5. Embedding interfaces

    Although go does not offer inheritance, it is possible to create a new interfaces by embedding other interfaces.

6. Zero value of Interface

    The zero value of a interface is ```nil```. A ```nil``` interface has both its underlying value and as well as concrete type as ```nil```.
