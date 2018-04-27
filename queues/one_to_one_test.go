package queues

import (
	"runtime"
	"sync"
	"testing"

	diodes "code.cloudfoundry.org/go-diodes"
	"github.com/tidwall/fastlane"
)

func BenchmarkSingleProducerSingleConsumerChannel(b *testing.B) {
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
			<-q
		}
		wg.Done()
	}(b.N)

	wg.Wait()
}

func BenchmarkSingleProducerSingleConsumerDiode(b *testing.B) {
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
			d.Next()
		}
		wg.Done()
	}(b.N)

	wg.Wait()
}

func BenchmarkSingleProducerSingleConsumerFastlane(b *testing.B) {
	var ch fastlane.ChanUint64
	var wg sync.WaitGroup
	wg.Add(2)

	go func(n int) {
		for i := 0; i < n; i++ {
			ch.Recv()
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
}
