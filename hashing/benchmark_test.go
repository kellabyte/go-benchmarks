package hashing

import (
	"encoding/binary"
	"encoding/hex"
	"hash/crc32"
	"hash/fnv"
	"strconv"
	"testing"

	xxhash "github.com/OneOfOne/xxhash"
	xxhashfast "github.com/cespare/xxhash"
	dchestsip "github.com/dchest/siphash"
	"github.com/dgryski/go-farm"
	highway "github.com/dgryski/go-highway"
	"github.com/dgryski/go-marvin32"
	"github.com/dgryski/go-metro"
	"github.com/dgryski/go-sip13"
	"github.com/dgryski/go-spooky"
	"github.com/dgryski/go-stadtx"
	"github.com/dgryski/go-t1ha"
	tsip "github.com/dgryski/trifles/tsip/go"
	blake2b "github.com/minio/blake2b-simd"
	miniohighway "github.com/minio/highwayhash"
	"github.com/opennota/fasthash"
	"github.com/rbastic/go-zaphod64"
	"github.com/surge/cityhash"
)

// Written in 2012 by Dmitry Chestnykh.
// Expanded to include multiple hash functions by Damian Gryski, 2015.
//
// To the extent possible under law, the author have dedicated all copyright
// and related and neighboring rights to this software to the public domain
// worldwide. This software is distributed without any warranty.
// http://creativecommons.org/publicdomain/zero/1.0/

var (
	key0, key1         uint64
	buf                = make([]byte, 8<<10)
	crc32Table         = crc32.MakeTable(crc32.Castagnoli)
	minioHighwayKey, _ = hex.DecodeString("000102030405060708090A0B0C0D0E0FF0E0D0C0B0A090807060504030201000")
)

var hblake2b = func(k []byte) uint64 {
	sum := blake2b.Sum256(k)
	return binary.LittleEndian.Uint64(sum[:8])
}

func BenchmarkBlake2B(b *testing.B) { benchmarkHash(b, "Blake2B", hblake2b) }

var hcrc32 = func(k []byte) uint64 { return uint64(crc32.Checksum(k, crc32Table)) }

func BenchmarkCRC32(b *testing.B) { benchmarkHash(b, "CRC32", hcrc32) }

var hspooky = func(k []byte) uint64 { return spooky.Hash64(k) }

func BenchmarkSpooky(b *testing.B) { benchmarkHash(b, "Spooky", hspooky) }

var hsiphash = func(k []byte) uint64 { return dchestsip.Hash(0, 0, k) }

func BenchmarkSipHash(b *testing.B) { benchmarkHash(b, "SipHash", hsiphash) }

var hfarm = func(k []byte) uint64 { return farm.Hash64(k) }

func BenchmarkFarm(b *testing.B) { benchmarkHash(b, "Farm", hfarm) }

var hcity = func(k []byte) uint64 { return cityhash.CityHash64(k, uint32(len(k))) }

func BenchmarkCity(b *testing.B) { benchmarkHash(b, "City", hcity) }

var hmetro = func(k []byte) uint64 { return metro.Hash64(k, 0) }

func BenchmarkMetro(b *testing.B) { benchmarkHash(b, "Metro", hmetro) }

var hxxhash = func(k []byte) uint64 { return xxhash.Checksum64(k) }

func BenchmarkXXHash(b *testing.B) { benchmarkHash(b, "XXHash", hxxhash) }

var hxxhashfast = func(k []byte) uint64 { return xxhashfast.Sum64(k) }

func BenchmarkXXFast(b *testing.B) { benchmarkHash(b, "XXFast", hxxhashfast) }

var fsthash = func(k []byte) uint64 { return fasthash.Hash64(0, k) }

func BenchmarkFasthash(b *testing.B) { benchmarkHash(b, "Fasthash", fsthash) }

var high = func(k []byte) uint64 { return highway.Hash(highway.Lanes{}, k) }

func BenchmarkHighway(b *testing.B) { benchmarkHash(b, "Highway", high) }

var hminiohighway = func(k []byte) uint64 { return miniohighway.Sum64(k, minioHighwayKey) }

func BenchmarkMinioHighway(b *testing.B) { benchmarkHash(b, "MinioHighway", hminiohighway) }

var marvin = func(k []byte) uint64 { return uint64(marvin32.Sum32(0, k)) }

func BenchmarkMarvin32(b *testing.B) { benchmarkHash(b, "Marvin32", marvin) }

var sip13hash = func(k []byte) uint64 { return sip13.Sum64(0, 0, k) }

func BenchmarkSip13Hash(b *testing.B) { benchmarkHash(b, "Sip13", sip13hash) }

var fnvh = fnv.New64a()

var fnv64 = func(k []byte) uint64 {
	fnvh.Reset()
	fnvh.Write(k)
	return fnvh.Sum64()
}

func BenchmarkFNV1(b *testing.B) { benchmarkHash(b, "fnv1a", fnv64) }

var ht1ha = func(k []byte) uint64 { return t1ha.Sum64(k, 0) }

func BenchmarkT1ha(b *testing.B) { benchmarkHash(b, "T1ha", ht1ha) }

var zaphodSeed zaphod64.State

var hzaphod64 = func(k []byte) uint64 { return uint64(zaphod64.HashWithState(&zaphodSeed, k, uint64(len(k)))) }

func BenchmarkZaphod64(b *testing.B) { benchmarkHash(b, "Zaphod64", hzaphod64) }

var stadtxState stadtx.State

var hstadtx = func(k []byte) uint64 { return stadtx.Hash(&stadtxState, k) }

func BenchmarkStadtx(b *testing.B) { benchmarkHash(b, "Stadtx", hstadtx) }

var htsip = func(k []byte) uint64 { return tsip.HashASM(0, 0, k) }

func BenchmarkTsip(b *testing.B) { benchmarkHash(b, "Tsip", htsip) }

func benchmarkHash(b *testing.B, str string, h func([]byte) uint64) {
	var sizes = []int{1, 2, 4, 8, 32, 64, 128, 256, 512, 1024, 4096, 8192}
	for _, n := range sizes {
		b.Run(strconv.Itoa(n), func(b *testing.B) { benchmarkHashn(b, int64(n), h) })
	}
}

var total uint64

func benchmarkHashn(b *testing.B, size int64, h func([]byte) uint64) {
	b.SetBytes(size)
	for i := 0; i < b.N; i++ {
		total += h(buf[:size])
	}
}
