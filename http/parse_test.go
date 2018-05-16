package main

import (
	"crypto/rand"
	"encoding/binary"
	"testing"
)

const requests = "PUT / HTTP/1.0\r\n"

// createHTTPReqyestBuffer creates a HTTP request buffer payload with the following layout.
//
// PUT / HTTP/1.0\r\n[uint32][uint32][key-bytes][value-bytes]
func createHTTPRequestBuffer(key []byte, value []byte) []byte {
	lengthOfRequestHeader := len(requests)
	lengthOfKey := len(key)
	lengthOfValue := len(value)
	buffer := make([]byte, lengthOfRequestHeader+4+4+lengthOfKey+lengthOfValue)

	offset := 0
	copy(buffer[offset:lengthOfRequestHeader], requests)
	offset += lengthOfRequestHeader
	binary.LittleEndian.PutUint32(buffer[offset:offset+4], uint32(lengthOfKey))
	offset += 4
	binary.LittleEndian.PutUint32(buffer[offset:offset+4], uint32(lengthOfValue))
	offset += 4
	copy(buffer[offset:], key)
	offset += len(key)
	copy(buffer[offset:], value)

	return buffer
}

func createRandomKeyValue(keyLength int, valueLength int) ([]byte, []byte) {
	key := make([]byte, keyLength)
	value := make([]byte, valueLength)
	rand.Read(key)
	rand.Read(value)
	return key, value
}

func createRequests(numberOfRequests int, keyLength int, valueLength int) []byte {
	lengthOfRequestHeaders := len(requests) * numberOfRequests
	lengthOfOffsets := (4 + 4) * numberOfRequests
	lengthOfKeys := keyLength * numberOfRequests
	lengthOfValues := valueLength * numberOfRequests
	lengthOfBuffer := lengthOfRequestHeaders + lengthOfOffsets + lengthOfKeys + lengthOfValues
	buffer := make([]byte, lengthOfBuffer)

	offset := 0
	for i := 0; i < numberOfRequests; i++ {
		key, value := createRandomKeyValue(16, 64)
		request := createHTTPRequestBuffer(key, value)
		copy(buffer[offset:lengthOfBuffer], request)
		offset += len(request)
	}
	return buffer
}

// BenchmarkFixedParser1000 splits b.N desired iterations from gobench (example: 200000000)
// into 1000 batches of runs so that we can parse large chunks of requests and invalidate
// CPU caches correctly without running out of memory trying to allocate 200000000 requests.
func BenchmarkFixedParser1000(b *testing.B) {
	benchmarkFixedParser(b, b.N, 1000, 16, 64)
}

func benchmarkFixedParser(b *testing.B, numberOfRequests int, splitSize int, keyLength int, valueLength int) {
	buffer := createRequests(numberOfRequests/splitSize, keyLength, valueLength)
	b.ResetTimer()

	for c := 0; c < splitSize; c++ {
		offset := 0
		for i := 0; i < numberOfRequests/splitSize; i++ {
			offset += 16
			lengthOfKey := int(binary.LittleEndian.Uint32(buffer[offset : offset+4]))
			offset += +4
			lengthOfValue := int(binary.LittleEndian.Uint32(buffer[offset : offset+4]))
			offset += 4
			key := buffer[offset : offset+lengthOfKey]
			offset += lengthOfKey
			value := buffer[offset : offset+lengthOfValue]
			offset += lengthOfValue

			_ = key
			_ = value
			b.SetBytes(1)
		}
	}
}
