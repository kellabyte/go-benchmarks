package json

import (
	"strconv"
	"testing"

	"github.com/mailru/easyjson"
)

func BenchmarkEasyjsonUnmarshal(b *testing.B) {
	b.Run(strconv.Itoa(len(smallStructText)), func(b *testing.B) { benchmarkEasyjsonUnmarshalSmallData(b) })
	b.Run(strconv.Itoa(len(largeStructText)), func(b *testing.B) { benchmarkEasyjsonUnmarshalLargeData(b) })
}

func benchmarkEasyjsonUnmarshalLargeData(b *testing.B) {
	b.SetBytes(int64(len(largeStructText)))

	for i := 0; i < b.N; i++ {
		var s LargeStruct

		if err := s.UnmarshalJSON(largeStructText); err != nil {
			b.Error(err)
		}
	}
}

func benchmarkEasyjsonUnmarshalSmallData(b *testing.B) {
	b.SetBytes(int64(len(smallStructText)))

	for i := 0; i < b.N; i++ {
		var s Entities

		if err := s.UnmarshalJSON(smallStructText); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEasyjsonMarshal(b *testing.B) {
	b.Run(strconv.Itoa(len(smallStructText)), func(b *testing.B) { benchmarkEasyjsonMarshalSmallData(b) })
	b.Run(strconv.Itoa(len(largeStructText)), func(b *testing.B) { benchmarkEasyjsonMarshalLargeData(b) })
}

func benchmarkEasyjsonMarshalLargeData(b *testing.B) {
	var l int64

	for i := 0; i < b.N; i++ {
		data, err := easyjson.Marshal(&largeStructData)
		if err != nil {
			b.Error(err)
		}
		l = int64(len(data))
	}

	b.SetBytes(l)
}

func benchmarkEasyjsonMarshalSmallData(b *testing.B) {
	var l int64

	for i := 0; i < b.N; i++ {
		data, err := easyjson.Marshal(&smallStructData)
		if err != nil {
			b.Error(err)
		}
		l = int64(len(data))
	}

	b.SetBytes(l)
}
