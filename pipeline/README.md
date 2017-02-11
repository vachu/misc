# The Pipeline

## Inspiration
I have always been enthralled by the elegance of Unix Pipes & Filters and always wanted to replicate in programming.  I wanted to chain multiple threads together in an arbitrary pipeline by which I could pump input data from one end and get the processed output data from the other end - just like the Unix pipline of Filters.

Long back I tried something with "queues" in Visual C++  & Win32 Threads but soon abadoned it as the code became so complex and unwieldy.  Then I got engrossed in Golang and quickly realized that the Channels and goroutines are directly helpful in coding what I wanted for a long time.  The culmination of that prolonged desire to implement something akin to Unix Pipeline, is this project

## Overview of this Implementation
As of this writing, this module was tested with go 1.7.4 (v1.8 is just a few weeks from now).  The main function is `BuildPipeline()` which takes variable list of functions as arguments and returns 3 channels `in`, `out` and `diag`.  The calling function (of `BuildPipeline()`) is then supposed to feed data through `in` and get the final output from `out`.  The `diag` channel is expected to stream diagnostic / error messages that might emanate from anywhere within the built pipeline.

Each of the functions supplied to `BuildPipeline` is akin to a Unix-Filter, taking a single "generic" parameter as an argument and returning a pair of generic return value and an error object.  This generic return value is then fed as input to the next function (if any) in the built pipeline.  Each of the 'pipelined' functions is run in a separate thread / goroutine.

The pipeline processing could be terminated normally by simpling closing the `in` channel, signifying that there are no further input to be processed.

## The Gory Details
`diag` is one of the channels returned by `BuildPipeline()` and it is important to note that it is a *merged channel*.  To understand better, I need to go into the implementation details a bit.
* Each of the functions supplied to `BuildPipeline()` is run as a separate **node** which is nothing but another goroutine and it's infact this node that internally consumes the channel-triplet created by the calling  `BuildPipeline()`
* `BuildPipeline()` then connects these individual nodes by making the `out` channel of node-n as the `in` channel of node-(n+1)
* Since a `diag` channel is created for each of the nodes, they need to be merged into a single channel before, so that the main calling (or client) function doesn't lose out on any of the diganostic messages from the nodes.
* Hence, in the ultimate channel triplet returned by `BuildPipeline()`
 * `in` channel is the one that's created for the first node
 * `out` channel is the one created for the last node
 * `diag` channel is the merged channel for the all `diag`-s created for the nodes

-----