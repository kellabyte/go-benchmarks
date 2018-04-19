package json

import (
	"encoding/json"
	"testing"
)

func BenchmarkStdUnmarshalLargeData(b *testing.B) {
	b.SetBytes(int64(len(largeStructText)))

	for i := 0; i < b.N; i++ {
		var s LargeStruct

		if err := json.Unmarshal(largeStructText, &s); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkStdUnmarshalSmallData(b *testing.B) {
	b.SetBytes(int64(len(smallStructText)))

	for i := 0; i < b.N; i++ {
		var s Entities

		if err := json.Unmarshal(smallStructText, &s); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkStdMarshalLargeData(b *testing.B) {
	var l int64

	for i := 0; i < b.N; i++ {
		data, err := json.Marshal(&largeStructData)
		if err != nil {
			b.Error(err)
		}
		l = int64(len(data))
	}

	b.SetBytes(l)
}

func BenchmarkStdMarshalSmallData(b *testing.B) {
	var l int64

	for i := 0; i < b.N; i++ {
		data, err := json.Marshal(&smallStructData)
		if err != nil {
			b.Error(err)
		}
		l = int64(len(data))
	}

	b.SetBytes(l)
}
