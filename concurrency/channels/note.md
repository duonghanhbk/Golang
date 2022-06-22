## Channels

1. What are channels

    Channels can be thought of as pipes using which Goroutines communicate. Similar to how water flows from one end to another in a pipe, data can be sent from one end and received from the other end using channels.

    ``Declaring channels``

    - Each channel has a type associated with it. This type is the type of data that the channel is allowed to transport. No other type is allowed to be transported using the channel.
    - chan T is a channel of type T
    - The zero value of a channel is nil. nil channels are not of any use and hence the channel has to be defined using make similar to maps and slices.

2. Sending and receiving from a channel

    ```
    data := <- a // read from channel a
    a <- data // write to channel a
    ```

    **
    Sends and receives are blocking by default
    **

3. Deadlock

    - One important factor to consider while using channels is deadlock. If a Goroutine is sending data on a channel, then it is expected that some other Goroutine should be receiving the data. If this does not happen, then the program will panic at runtime with Deadlock.
    - Similarly, if a Goroutine is waiting to receive data from a channel, then some other Goroutine is expected to write data on that channel, else the program will panic.

4. Buffered channels

    It is possible to create a channel with a buffer. Sends to a buffered channel are blocked only when the buffer is full. Similarly receives from a buffered channel are blocked only when the buffer is empty.