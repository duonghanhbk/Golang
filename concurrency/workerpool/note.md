## Worker Pool

1. Worker Pool Implementation

    - Creation of a pool of Goroutines which listen on an input buffered channel waiting for jobs to be assigned
    - Addition of jobs to the input buffered channel
    - Writing results to an output buffered channel after job completion
    - Read and print results from the output buffered channel