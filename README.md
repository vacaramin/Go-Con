# Go-Con

## Repo Structure

- ../
- ./
  - channels/
  - go-routines/
  - mutex/
  - waitgroups/

Each of the above have concepts covered within the repo

# Problems Statement

You have been given a Go program where you need to implement a producer-consumer scenario using channels. 

- **Producer Goroutine:** Generates numbers from 1 to 10 and sends them to a channel.
- **Consumer Goroutine:** Receives numbers from the channel and prints them.

Ensure that the main function waits for both the producer and consumer to complete their execution.

### Requirements

1. Implement the producer goroutine that sends numbers to a channel.
2. Implement the consumer goroutine that receives and prints numbers from the channel.
3. Use `sync.WaitGroup` to synchronize the main function with the completion of both goroutines.
4. Use a buffered channel with a capacity of 5.


# Goroutines Problem

## Problem Statement

Write a Go program to compute the sum of numbers from 1 to 100 using multiple goroutines.

- Divide the range [1, 100] into 10 chunks.
- Each goroutine should compute the sum of one chunk.
- The main function should aggregate the results from all goroutines and print the final sum.

### Requirements

1. Create 10 goroutines where each computes the sum of a chunk of numbers.
2. Use a channel to collect the sum from each goroutine.
3. Ensure the main function waits for all goroutines to finish and then prints the total sum.


# Mutex Problem

## Problem Statement

You need to write a Go program that increments a shared counter variable concurrently using multiple goroutines.

- Create a shared counter variable.
- Use 10 goroutines to increment the counter 100 times each.
- Protect the shared counter variable with a mutex to prevent race conditions.

### Requirements

1. Implement a mutex to protect access to the shared counter variable.
2. Each goroutine should increment the counter 100 times.
3. After all goroutines complete, print the final value of the counter.

# WaitGroups Problem

## Problem Statement

Write a Go program to demonstrate the use of `sync.WaitGroup` to wait for a set of tasks to complete.

- Create a Go program that spawns 5 goroutines, each of which performs a task that takes a random amount of time to complete.
- Use `sync.WaitGroup` to wait for all goroutines to finish their tasks before printing a message indicating completion.

### Requirements

1. Implement a `sync.WaitGroup` to manage the completion of 5 goroutines.
2. Each goroutine should simulate a task with a random delay.
3. Print "All tasks completed" after all goroutines have finished.
