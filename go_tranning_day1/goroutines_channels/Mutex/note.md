- Mutex is used to provide a locking mechanism to ensure that only one Goroutine is running the critical section
   of code at any point in time to prevent race conditions from happening
- If one Goroutine already holds the lock and if a new Goroutine is trying to acquire a lock, the new Goroutine will
   be blocked until the mutex is unlocked
   