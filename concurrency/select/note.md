## Select

1. What is select?

    - The select statement is used to choose from multiple send/receive channel operations.
    - The select statement blocks until one of the send/receive operations is ready. If multiple operations are ready, one of them is chosen at random.
    - The syntax is similar to switch except that each of the case statements will be a channel operation. Let's dive right into some code for better understanding.