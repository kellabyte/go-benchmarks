//go:linkname nanotime runtime.nanotime
package time

import (
	"testing"
	_ "unsafe"

	"github.com/codahale/hdrhistogram"
	"github.com/loov/hrtime"
	"github.com/tylertreat/hdrhistogram-writer"
)

func nanotime() int64

func BenchmarkNanotime(b *testing.B) {
	startNanos := make([]int64, b.N)
	endNanos := make([]int64, b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		startNanos[i] = nanotime()
		endNanos[i] = nanotime()
		b.SetBytes(1)
	}

	b.StopTimer()
	recordLatencyDistribution("nanotime", b.N, startNanos, endNanos)
}

func BenchmarkHrtime(b *testing.B) {
	bench := hrtime.NewBenchmarkTSC(b.N)

	b.ResetTimer()
	for bench.Next() {
		b.SetBytes(1)
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
