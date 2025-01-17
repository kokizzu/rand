// Copyright 2022 Gregory Petrosyan <gregory.petrosyan@gmail.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

//go:build benchstd

package rand_test

import (
	"math"
	"math/rand"
	"testing"
)

var (
	sinkRand *rand.Rand
)

func BenchmarkUint64(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s uint64
		for pb.Next() {
			s = rand.Uint64()
		}
		sinkUint64 = s
	})
}

func BenchmarkFloat64(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s float64
		for pb.Next() {
			s = rand.Float64()
		}
		sinkFloat64 = s
	})
}

func BenchmarkIntn(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s int
		for pb.Next() {
			s = rand.Intn(small)
		}
		sinkInt = s
	})
}

func BenchmarkIntn_Big(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var s int
		for pb.Next() {
			s = rand.Intn(math.MaxInt - small)
		}
		sinkInt = s
	})
}

func BenchmarkRand_New(b *testing.B) {
	var s *rand.Rand
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s = rand.New(rand.NewSource(int64(i)))
	}
	sinkRand = s
}

func BenchmarkRand_ExpFloat64(b *testing.B) {
	var s float64
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.ExpFloat64()
	}
	sinkFloat64 = s
}

func BenchmarkRand_Float32(b *testing.B) {
	var s float32
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.Float32()
	}
	sinkFloat32 = s
}

func BenchmarkRand_Float64(b *testing.B) {
	var s float64
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.Float64()
	}
	sinkFloat64 = s
}

func BenchmarkRand_Int(b *testing.B) {
	var s int
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.Int()
	}
	sinkInt = s
}

func BenchmarkRand_Int31(b *testing.B) {
	var s int32
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.Int31()
	}
	sinkInt32 = s
}

func BenchmarkRand_Int31n(b *testing.B) {
	var s int32
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.Int31n(small)
	}
	sinkInt32 = s
}

func BenchmarkRand_Int31n_Big(b *testing.B) {
	var s int32
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.Int31n(math.MaxInt32 - small)
	}
	sinkInt32 = s
}

func BenchmarkRand_Int63(b *testing.B) {
	var s int64
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.Int63()
	}
	sinkInt64 = s
}

func BenchmarkRand_Int63n(b *testing.B) {
	var s int64
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.Int63n(small)
	}
	sinkInt64 = s
}

func BenchmarkRand_Int63n_Big(b *testing.B) {
	var s int64
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.Int63n(math.MaxInt64 - small)
	}
	sinkInt64 = s
}

func BenchmarkRand_Intn(b *testing.B) {
	var s int
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.Intn(small)
	}
	sinkInt = s
}

func BenchmarkRand_Intn_Big(b *testing.B) {
	var s int
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.Intn(math.MaxInt - small)
	}
	sinkInt = s
}

func BenchmarkRand_NormFloat64(b *testing.B) {
	var s float64
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		s = r.NormFloat64()
	}
	sinkFloat64 = s
}

func BenchmarkRand_Perm(b *testing.B) {
	b.ReportAllocs()
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		r.Perm(tiny)
	}
}

func BenchmarkRand_Read(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	p := make([]byte, 256)
	b.SetBytes(int64(len(p)))
	for i := 0; i < b.N; i++ {
		_, _ = r.Read(p[:])
	}
}

func BenchmarkRand_Seed(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	for i := 0; i < b.N; i++ {
		r.Seed(int64(i))
	}
}

func BenchmarkRand_Shuffle(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	a := make([]int, tiny)
	for i := 0; i < b.N; i++ {
		r.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	}
}

func BenchmarkRand_ShuffleOverhead(b *testing.B) {
	r := rand.New(rand.NewSource(1))
	a := make([]int, tiny)
	for i := 0; i < b.N; i++ {
		r.Shuffle(len(a), func(i, j int) {})
	}
}

func BenchmarkRand_Uint32(b *testing.B) {
	var s uint32
	r := rand.New(rand.NewSource(1))
	b.SetBytes(4)
	for i := 0; i < b.N; i++ {
		s = r.Uint32()
	}
	sinkUint32 = s
}

func BenchmarkRand_Uint64(b *testing.B) {
	var s uint64
	r := rand.New(rand.NewSource(1))
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		s = r.Uint64()
	}
	sinkUint64 = s
}
