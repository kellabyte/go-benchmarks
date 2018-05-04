//go:linkname nanotime runtime.nanotime
package queues

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	_ "unsafe"

	diodes "code.cloudfoundry.org/go-diodes"
	"github.com/codahale/hdrhistogram"
	"github.com/loov/hrtime"
	"github.com/pltr/onering"
	"github.com/tidwall/fastlane"
	"github.com/tylertreat/hdrhistogram-writer"
)

func nanotime() int64

func Benchmark1Producer1ConsumerChannel(b *testing.B) {
	startNanos := make([]int64, b.N)
	endNanos := make([]int64, b.N)

	q := make(chan int64, 8192)

	var wg sync.WaitGroup
	wg.Add(2)

	go func(n int) {
		runtime.LockOSThread()
		for i := 0; i < n; i++ {
			q <- int64(i)
		}
		wg.Done()
	}(b.N)

	b.ResetTimer()
	go func(n int) {
		runtime.LockOSThread()
		for i := 0; i < n; i++ {
			startNanos[i] = nanotime()
			<-q
			endNanos[i] = nanotime()
			b.SetBytes(1)
		}
		wg.Done()
	}(b.N)

	wg.Wait()

	b.StopTimer()
	recordLatencyDistribution("channel", b.N, startNanos, endNanos)
}

func Benchmark1Producer1ConsumerDiode(b *testing.B) {
	startNanos := make([]int64, b.N)
	endNanos := make([]int64, b.N)

	d := diodes.NewPoller(diodes.NewOneToOne(b.N, diodes.AlertFunc(func(missed int) {
		panic("Oops...")
	})))

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < b.N; i++ {
			d.Set(diodes.GenericDataType(&i))
		}
		wg.Done()
	}()

	b.ResetTimer()
	go func(n int) {
		for i := 0; i < b.N; i++ {
			startNanos[i] = nanotime()
			d.Next()
			endNanos[i] = nanotime()
			b.SetBytes(1)
		}
		wg.Done()
	}(b.N)

	wg.Wait()

	b.StopTimer()
	recordLatencyDistribution("diode", b.N, startNanos, endNanos)
}

func Benchmark1Producer1ConsumerFastlane(b *testing.B) {
	startNanos := make([]int64, b.N)
	endNanos := make([]int64, b.N)

	var ch fastlane.ChanUint64
	var wg sync.WaitGroup
	wg.Add(2)

	go func(n int) {
		for i := 0; i < n; i++ {
			startNanos[i] = nanotime()
			ch.Recv()
			endNanos[i] = nanotime()
			b.SetBytes(1)
		}
		wg.Done()
	}(b.N)

	b.ResetTimer()
	go func(n int) {
		for i := 0; i < n; i++ {
			ch.Send(uint64(i))
		}
		wg.Done()
	}(b.N)

	wg.Wait()

	b.StopTimer()
	recordLatencyDistribution("fastlane", b.N, startNanos, endNanos)
}

func Benchmark1Producer1ConsumerOneRing(b *testing.B) {
	startNanos := make([]int64, b.N)
	endNanos := make([]int64, b.N)

	var ring onering.SPSC
	ring.Init(8192)
	var wg sync.WaitGroup
	wg.Add(2)

	go func(n int) {
		runtime.LockOSThread()
		for i := 0; i < n; i++ {
			ring.Put(int64(i))
		}
		ring.Close()
		wg.Done()
	}(b.N)

	b.ResetTimer()
	go func(n int64) {
		runtime.LockOSThread()
		var i, v int64
		//startNanos[i] = time.Now().UnixNano()
		startNanos[i] = nanotime()
		for ring.Get(&v) {
			if v != i {
				fmt.Printf("Expected %d got %d", i, v)
				panic(v)
			}
			//endNanos[i] = time.Now().UnixNano()
			endNanos[i] = nanotime()
			i++
			b.SetBytes(1)

			if i < n {
				//startNanos[i] = time.Now().UnixNano()
				startNanos[i] = nanotime()
			}
		}
		wg.Done()
	}(int64(b.N))

	wg.Wait()

	b.StopTimer()
	recordLatencyDistribution("onering", b.N, startNanos, endNanos)
}

func Benchmark1Producer1ConsumerOneRing2(b *testing.B) {
	bench := hrtime.NewBenchmarkTSC(b.N)

	var ring onering.SPSC
	ring.Init(8192)
	var wg sync.WaitGroup
	wg.Add(2)

	go func(n int) {
		runtime.LockOSThread()
		for i := 0; i < n; i++ {
			ring.Put(int64(i))
		}
		ring.Close()
		wg.Done()
	}(b.N)

	b.ResetTimer()
	go func(n int64) {
		runtime.LockOSThread()
		var v int64
		for bench.Next() {
			ring.Get(&v)
			b.SetBytes(1)
		}
		wg.Done()
	}(int64(b.N))

	wg.Wait()

	b.StopTimer()
	recordLatencyDistributionBenchmark("onering-hrtime", bench)
}

func BenchmarkNanotimeOverhead(b *testing.B) {
	startNanos := make([]int64, b.N)
	endNanos := make([]int64, b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		startNanos[i] = nanotime()
		endNanos[i] = nanotime()
	}

	b.StopTimer()
	recordLatencyDistribution("nanotime", b.N, startNanos, endNanos)
}

func BenchmarkHrtimeOverhead(b *testing.B) {
	bench := hrtime.NewBenchmarkTSC(b.N)

	b.ResetTimer()
	for bench.Next() {
	}

	b.StopTimer()
	recordLatencyDistributionBenchmark("hrtime", bench)
}

func recordLatencyDistribution(name string, count int, startNanos []int64, endNanos []int64) {
	histogram := hdrhistogram.New(1, 1000000, 5)
	for i := 0; i < count; i++ {
		diff := endNanos[i] - startNanos[i]
		histogram.RecordValue(diff)
	}
	histwriter.WriteDistributionFile(histogram, nil, 1.0, "../results/"+name+".histogram")
}

func recordLatencyDistributionBenchmark(name string, bench *hrtime.BenchmarkTSC) {
	histogram := hdrhistogram.New(1, 1000000, 5)
	for _, lap := range bench.Laps() {
		histogram.RecordValue(int64(lap))
	}
	histwriter.WriteDistributionFile(histogram, nil, 1.0, "../results/"+name+".histogram")
}
