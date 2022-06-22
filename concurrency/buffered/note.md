## Buffered Channels

1. What are buffered channels?

    It is possible to create a channel with a buffer. Sends to a buffered channel are blocked only when the buffer is full. Similarly receives from a buffered channel are blocked only when the buffer is empty.

    Buffered channels can be created by passing an additional capacity parameter to the make function which specifies the size of the buffer.

    ```
    ch := make(chan type, capacity)
    ```