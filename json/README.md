This directory includes json benchmarks.

# Hardware
```
MacBookPro11,3
Processor: Intel Core i7 2.8 GHz 4 Cores
L2 Cache per core: 256 KB
L3 Cache: 6 MB
Memory: 16GB
```

# Results
```
make json

goos: darwin
goarch: amd64
pkg: github.com/euskadi31/go-benchmarks/json
BenchmarkEasyjsonUnmarshalLargeData-8   	   30000	     50991 ns/op	 316.74 MB/s	    9792 B/op	     128 allocs/op
BenchmarkEasyjsonUnmarshalSmallData-8   	 3000000	       629 ns/op	 130.18 MB/s	     128 B/op	       3 allocs/op
BenchmarkEasyjsonMarshalLargeData-8     	  100000	     17635 ns/op	 510.17 MB/s	   10369 B/op	      10 allocs/op
BenchmarkEasyjsonMarshalSmallData-8     	 5000000	       307 ns/op	 263.41 MB/s	     240 B/op	       2 allocs/op
BenchmarkStdUnmarshalLargeData-8        	   10000	    169305 ns/op	  95.40 MB/s	   10720 B/op	     140 allocs/op
BenchmarkStdUnmarshalSmallData-8        	 1000000	      1805 ns/op	  45.42 MB/s	     608 B/op	      11 allocs/op
BenchmarkStdMarshalLargeData-8          	   20000	     82255 ns/op	 109.38 MB/s	   20291 B/op	      17 allocs/op
BenchmarkStdMarshalSmallData-8          	 1000000	      1372 ns/op	  59.00 MB/s	     552 B/op	       7 allocs/op
PASS
ok  	github.com/euskadi31/go-benchmarks/json	15.687s
```
