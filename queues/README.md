# Introduction
This directory includes queue implementation benchmarks for various different configurations of producer and consumer counts. Each data structure may have different safety and ordering guarantees. This document should include results along with safety and ordering information to inform the viewers of the tradeoffs.

# Guarantees
Different queue-like data structures sometimes make tradeoffs for performance so we can't treat all of these libraries as equals. What they do is give you a basic FIFO interface but what they guarantee in safety and ordering can differ. This section documents the guarantees each provide so you can decide what is best for you.

```
Library             Data loss    Order preserving    Producers    Consumers
--------------------------------------------------------------------------------
Channels            No           Yes                 1+           1+
Diodes              Yes          Yes                 1+           1+
Fastlane            No           Yes                 1            1
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
Benchmark1Producer1ConsumerChannel-8    	20000000	        84.8 ns/op	       0 B/op	       0 allocs/op
Benchmark1Producer1ConsumerDiode-8      	20000000	        87.2 ns/op	      15 B/op	       0 allocs/op
Benchmark1Producer1ConsumerFastlane-8   	20000000	        66.7 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/kellabyte/go-benchmarks/queues	5.235s
```