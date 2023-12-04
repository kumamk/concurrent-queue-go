## Description

A basic implementation of concurrent thread-safe queue in GO. Queues are everywhere and basic ones are not thread safe. Doing concurrent operations on these naive queues can lead to data/result inconsistencies. As shown in the example how thread-safe queues can overcome it and provide consisten behaviour.

Firstly basic queue example is shown, adding and removing item in sequential manner.
When tried to add 100 concurrent item in basic queue, results are inconsistent as shown below in test sample data.
Lastly when tried the concurrent operations on thread-safe queue, results are consistent as per expectation.

## Pre Requirements

* GO installed in the system and any supported IDE

## Test

```bash
# run test
$ go run main.go

# test run data
------Normal queue sequential flow-----

Add 3 items to normal queue
Remove 3 items from normal queue
Normal queue size 0

------Concurrent write to Normal queue-----

Adding 100 items conncurrently to normal queue
Normal queue size 54, ideally should be 100

-----Concurrent queue flow-----

Adding 100 items to concurrent queue
Removing 100 items from concurrent queue
Concurrent queue size 0
```