# Concurrency

To begin with concurrency is not Parallelism. Concurrency is givining fair chance to the all the tasks.
Idea of concurrency comes from the Operating systems.

Job of the operating system is to provide fair chance to for all the process to access resources. 

# What is a process?
A process is a instance of a running program. 
A Process provides an environment for program to execute

OS creates a process and allocates memory in virtual space, virtual address space will contain
Code segment
Data region -Global
Heap - Dynamic Memory location
Stack - Storing local variables

Threads are smalles unit of execution, each process has one thread i.e main thread. Process can have multiple threads and they share same address space.
Each thread has its own stack.

Threads run independent of each other, OS scheduler make scheduling decisions at thread level not at process level
Threads can run concurrently -each thread taking trun on individual core, or in parallel by running in different cores at the same time.

Thread states.
1. Runnable - When a process is created, the main thread is placed in ready queue, once cpu is available each thread moves in to
2. Executing state- Each thread is provided with finite time slice, if time slice exausted thread is preempted and place back into Runnable state,

While executing if the thread gets blocked on i/o operation or networks operation or event from other process thread is moved to waiting state untill the required operation is complete.

3. Waiting: once the i/o operation is complete thread is put back to the Runnable state.

Can we achieve by dividing our application in to process and threads, yes we can to a extent. However context switching for Process are expensive where as  threads context switch are not.

How efficient are threads?, by creating more number of threads chances of each threads will have to wait in runnable state and will lead to inefficiencies.

ulimit -a
Each thread is allocated its own stack and in this machine stack size is 8176KB(8MB). If Machine has 16 GB of memory we can create only 2000 threads.

Concurrency is expensive.

Threads communicate by sharing memory, sharing memory between threads is complex, Multiple threads sharing same memory will lead to data race and results in unpredictable results.

Advantages of Goroutine compared to threads,

Goroutines are lightweight compared to threads
Stack size is very small compared to that of threads(2kb)
Context switching is very cheap as it happens in user space
We can create many goroutine compared to that of threads


Go schedulers (M:N Schedular)

Go schedulr is part of the go runtime and is known as M:N schedular. And go runtime is part of the executable binary. 

Go runtime uses os thread to schedule goroutines 

Goroutines run in context of OS threads.

Go runtime creates the number of worker threads equal to GOMAXPROCS environmental value

default value of GOMAXPROCS is number of processors on machine

in case of quad core it will be 4 and octacore it will be 8 

Go schedular distributes runnable goroutines over os multiple OS threads that are created

At anytim N goroutines could be created over M OS threads that run on GOMAXPROCS number of cores.

Go implements asynchronus preemption of goroutines, ensuring cpu intensive routines not hogging the CPU and providing equal opprotunity for all goroutines to run.

Goroutines have similar states as that of threads


Work stealing.





Go routine








