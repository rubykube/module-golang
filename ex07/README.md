# Goroutines

Implement a function `Process` that returns channel of string(output channel) and takes a channel of string(input channel) as an argument.

*Run two goroutines in the function:*
- One is to perform operations on the strings received from the input channel and then write these strings to the output channel.
- Other one is to close the output channel after receiving all the objects from input channel.


 *Operations performed by the first goroutine:*
1. Put string into parentheses.
for example:
`hello` will be converted to `(hello)`
`hello world !!` will be converted to `(hello world !!)`
`(hello)` will be converted to `((hello))`
