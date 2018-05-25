package queues

import (
	"runtime"
	"sync"
	"testing"
	"unsafe"

	"code.cloudfoundry.org/go-diodes"
	"github.com/codahale/hdrhistogram"
	"github.com/loov/hrtime"
	"github.com/pltr/onering"
	"github.com/tidwall/fastlane"
	"github.com/tylertreat/hdrhistogram-writer"
)

func mknumslice(n int) []int64 {
	var s = make([]int64, n)
	for i := range s {
		s[i] = int64(i)
	}
	return s
}

func BenchmarkChannel(b *testing.B) {
	bench := hrtime.NewBenchmarkTSC(b.N)
	q := make(chan *int64, 8192)
	var numbers = mknumslice(b.N)
	var wg sync.WaitGroup
	wg.Add(2)

	go func(n int) {
		runtime.LockOSThread()
		for i := 0; i < n; i++ {
			q <- &numbers[i]
		}
		wg.Done()
	}(b.N)

	b.ResetTimer()
	go func() {
		runtime.LockOSThread()
		for bench.Next() {
			<-q
			b.SetBytes(1)
		}
		wg.Done()
	}()

	wg.Wait()

	b.StopTimer()
	recordLatencyDistributionBenchmark("channel", bench)
}

func BenchmarkDiode(b *testing.B) {
	bench := hrtime.NewBenchmarkTSC(b.N)

	d := diodes.NewPoller(diodes.NewOneToOne(b.N, diodes.AlertFunc(func(missed int) {
		panic("Oops...")
	})))
	var numbers = mknumslice(b.N)
	var wg sync.WaitGroup
	wg.Add(2)

	go func(n int) {
		for i := 0; i < n; i++ {
			d.Set(diodes.GenericDataType(&numbers[i]))
		}
		wg.Done()
	}(b.N)

	b.ResetTimer()
	go func() {
		for bench.Next() {
			d.Next()
			b.SetBytes(1)
		}
		wg.Done()
	}()

	wg.Wait()

	b.StopTimer()
	recordLatencyDistributionBenchmark("diode", bench)
}

func BenchmarkFastlane(b *testing.B) {
	bench := hrtime.NewBenchmarkTSC(b.N)
	var numbers = mknumslice(b.N)
	var ch fastlane.ChanPointer
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for bench.Next() {
			ch.Recv()
			b.SetBytes(1)
		}
		wg.Done()
	}()

	b.ResetTimer()
	go func(n int) {
		for i := 0; i < n; i++ {
			ch.Send(unsafe.Pointer(&numbers[i]))
		}
		wg.Done()
	}(b.N)

	wg.Wait()

	b.StopTimer()
	recordLatencyDistributionBenchmark("fastlane", bench)
}

func BenchmarkOneRing(b *testing.B) {
	bench := hrtime.NewBenchmarkTSC(b.N)
	var numbers = mknumslice(b.N)
	var ring = onering.New{Size: 8192}.SPSC()
	var wg sync.WaitGroup
	wg.Add(2)

	go func(n int) {
		runtime.LockOSThread()
		for i := 0; i < n; i++ {
			ring.Put(&numbers[i])
		}
		ring.Close()
		wg.Done()
	}(b.N)

	b.ResetTimer()
	go func(n int64) {
		runtime.LockOSThread()
		var v *int64
		for bench.Next() {
			ring.Get(&v)
			b.SetBytes(1)
		}
		wg.Done()
	}(int64(b.N))

	wg.Wait()

	b.StopTimer()
	recordLatencyDistributionBenchmark("onering", bench)
}

func recordLatencyDistributionBenchmark(name string, bench *hrtime.BenchmarkTSC) {
	histogram := hdrhistogram.New(1, 1000000, 5)
	for _, lap := range bench.Laps() {
		histogram.RecordValue(int64(lap))
	}
	histwriter.WriteDistributionFile(histogram, nil, 1.0, "../../results/"+name+".histogram")
}
