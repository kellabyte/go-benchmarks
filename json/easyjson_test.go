package json

import (
	"testing"

	"github.com/mailru/easyjson"
)

func BenchmarkEasyjsonUnmarshalLargeData(b *testing.B) {
	b.SetBytes(int64(len(largeStructText)))

	for i := 0; i < b.N; i++ {
		var s LargeStruct

		if err := s.UnmarshalJSON(largeStructText); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEasyjsonUnmarshalSmallData(b *testing.B) {
	b.SetBytes(int64(len(smallStructText)))

	for i := 0; i < b.N; i++ {
		var s Entities

		if err := s.UnmarshalJSON(smallStructText); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEasyjsonMarshalLargeData(b *testing.B) {
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

func BenchmarkEasyjsonMarshalSmallData(b *testing.B) {
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
