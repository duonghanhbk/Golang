## Mutex

1. Critical section

    When a program runs concurrently, the parts of code which modify shared resources should not be accessed by multiple Goroutines at the same time. This section of code that modifies shared resources is called critical section.

2. Mutex

    A Mutex is used to provide a locking mechanism to ensure that only one Goroutine is running the critical section of code at any point in time to prevent race conditions from happening.

    If one Goroutine already holds the lock and if a new Goroutine is trying to acquire a lock, the new Goroutine will be blocked until the mutex is unlocked.

3. Mutex vs Channels

    We have solved the race condition problem using both mutexes and channels. So how do we decide what to use when?

    - Use channels when Goroutines need to communicate with each other
    - Use mutexes when only one Goroutine should access the critical section of code.