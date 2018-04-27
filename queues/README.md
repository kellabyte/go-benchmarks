# Introduction
This directory includes queue implementation benchmarks for various different configurations of producer and consumer counts. Each data structure may have different safety and ordering guarantees. This document should include results along with safety and ordering information to inform the viewers of the tradeoffs.

# Guarantees
Different queue-like data structures sometimes make tradeoffs for performance so we can't treat all of these libraries as equals. What they do is give you a basic FIFO interface but what they guarantee in safety and ordering can differ. This section documents the guarantees each provide so you can decide what is best for you.

#### Data loss
Below is a description of data loss behaviours queues can experience that may be important for your use case to understand.

##### D1. Producers outpace consumers and overwrite unread records.
Often seen in ring buffer implementations if a producer hits the end of the buffer it loops back to the front and keeps writing. If the producer is outpacing the consumer it means the producer could eventually lap the consumer and overwrite records causing data loss.

##### D2. Producers outpace consumers and consumers skip records to keep up.
If the consumer isn't keeping up with the producer some implementations may skip records to keep up.

```
Library             Order preserving    Data loss   Producers    Consumers
--------------------------------------------------------------------------------
Channels            Yes                  No         1+           1+
Diodes              Yes                  D1         1+           1+
Fastlane            Yes                  No         1            1
```

# Hardware
```
MacBookPro11,2
Processor: Intel Core i7 2.2 GHz 4 Cores
L2 Cache per core: 256 KB
L3 Cache: 6 MB
Memory: 16GB
```

# Results
```
make queues

goos: darwin
goarch: amd64
pkg: github.com/kellabyte/go-benchmarks/queues
Benchmark1Producer1ConsumerChannel-8    20000000        84.8 ns/op       0 B/op	       0 allocs/op
Benchmark1Producer1ConsumerDiode-8      20000000        87.2 ns/op      15 B/op	       0 allocs/op
Benchmark1Producer1ConsumerFastlane-8   20000000        66.7 ns/op       0 B/op	       0 allocs/op
PASS
ok  	github.com/kellabyte/go-benchmarks/queues	5.235s
```