This directory includes queue implementation benchmarks for various different configurations of producer and consumer counts. Each data structure may have different safety and ordering guarantees. This document should include results along with safety and ordering information to inform the viewers of the tradeoffs.

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
BenchmarkSingleProducerSingleConsumerChannel-8 	20000000	84.8 ns/op	   0 B/op   0 allocs/op
BenchmarkSingleProducerSingleConsumerDiode-8   	20000000	87.6 ns/op	  15 B/op   0 allocs/op
PASS
ok  	github.com/kellabyte/go-benchmarks/queues	3.859s
```