package json

import (
	"encoding/json"
	"strconv"
	"testing"
)

func BenchmarkStdjsonUnmarshal(b *testing.B) {
	b.Run(strconv.Itoa(len(smallStructText)), func(b *testing.B) { benchmarkStdUnmarshalSmallData(b) })
	b.Run(strconv.Itoa(len(largeStructText)), func(b *testing.B) { benchmarkStdUnmarshalLargeData(b) })
}

func benchmarkStdUnmarshalLargeData(b *testing.B) {
	b.SetBytes(int64(len(largeStructText)))

	for i := 0; i < b.N; i++ {
		var s LargeStruct

		if err := json.Unmarshal(largeStructText, &s); err != nil {
			b.Error(err)
		}
	}
}

func benchmarkStdUnmarshalSmallData(b *testing.B) {
	b.SetBytes(int64(len(smallStructText)))

	for i := 0; i < b.N; i++ {
		var s Entities

		if err := json.Unmarshal(smallStructText, &s); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkStdjsonMarshal(b *testing.B) {
	b.Run(strconv.Itoa(len(smallStructText)), func(b *testing.B) { benchmarkStdMarshalSmallData(b) })
	b.Run(strconv.Itoa(len(largeStructText)), func(b *testing.B) { benchmarkStdMarshalLargeData(b) })
}

func benchmarkStdMarshalLargeData(b *testing.B) {
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

func benchmarkStdMarshalSmallData(b *testing.B) {
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
