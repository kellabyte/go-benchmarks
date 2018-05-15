package main

import (
	"bytes"
	"testing"
)

const requests = "PUT / HTTP/1.1\r\nk: 0000000002\r\nv: 0000000002\r\nHost: 127.0.0.1:8000\r\n\r\n"

func BenchmarkScanParser(b *testing.B) {
	buffer := []byte(requests)
	b.ResetTimer()

	for i := 0; i < len(buffer); i++ {
		if buffer[i] == 'k' {

		}
		if buffer[i] == 'v' {

		}
	}

}

func BenchmarkIndexByteParser(b *testing.B) {
	buffer := []byte(requests)
	b.ResetTimer()
	idx := bytes.IndexByte(buffer, 'k')
	idx = bytes.IndexByte(buffer[idx:], 'v')
}
