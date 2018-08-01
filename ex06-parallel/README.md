# Parallel Letter Frequency

Write a program that counts the frequency of letters in texts using parallel computation.

Parallelism is about doing things in parallel that can also be done
sequentially. A common example is counting the frequency of letters.

Create two functions:
Function Frequency(text string), to sequentially count letter frequencies in a single text.
Function ConcurrentFrequency(texts []string), to perform this exercise on
parallelism using Go concurrency features.
Make concurrent calls to Frequency and combine results to obtain the answer.
