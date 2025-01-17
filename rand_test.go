// Copyright 2022 Gregory Petrosyan <gregory.petrosyan@gmail.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

//go:build !benchx && !benchstd && !benchfast

package rand_test

import (
	"bytes"
	"math"
	"math/bits"
	"testing"

	"pgregory.net/rapid"

	"github.com/kokizzu/rand"
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
		s = rand.New(uint64(i))
	}
	sinkRand = s
}

func BenchmarkRand_New0(b *testing.B) {
	var s *rand.Rand
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s = rand.New()
	}
	sinkRand = s
}

func BenchmarkRand_New3(b *testing.B) {
	var s *rand.Rand
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s = rand.New(uint64(i), uint64(i), uint64(i))
	}
	sinkRand = s
}

func BenchmarkRand_NewInt(b *testing.B) {
	var s int
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s = rand.New().Int()
	}
	sinkInt = s
}

func BenchmarkRand_Get(b *testing.B) {
	var r rand.Rand
	var s *rand.Rand
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s = r.Get()
	}
	sinkRand = s
}

func BenchmarkRand_GetInt(b *testing.B) {
	var r rand.Rand
	var s int
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s = r.Get().Int()
	}
	sinkInt = s
}

func BenchmarkRand_ExpFloat64(b *testing.B) {
	var s float64
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.ExpFloat64()
	}
	sinkFloat64 = s
}

func BenchmarkRand_Float32(b *testing.B) {
	var s float32
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Float32()
	}
	sinkFloat32 = s
}

func BenchmarkRand_Float64(b *testing.B) {
	var s float64
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Float64()
	}
	sinkFloat64 = s
}

func BenchmarkRand_Int(b *testing.B) {
	var s int
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Int()
	}
	sinkInt = s
}

func BenchmarkRand_Int31(b *testing.B) {
	var s int32
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Int31()
	}
	sinkInt32 = s
}

func BenchmarkRand_Int31n(b *testing.B) {
	var s int32
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Int31n(small)
	}
	sinkInt32 = s
}

func BenchmarkRand_Int31n_Big(b *testing.B) {
	var s int32
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Int31n(math.MaxInt32 - small)
	}
	sinkInt32 = s
}

func BenchmarkRand_Int63(b *testing.B) {
	var s int64
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Int63()
	}
	sinkInt64 = s
}

func BenchmarkRand_Int63n(b *testing.B) {
	var s int64
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Int63n(small)
	}
	sinkInt64 = s
}

func BenchmarkRand_Int63n_Big(b *testing.B) {
	var s int64
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Int63n(math.MaxInt64 - small)
	}
	sinkInt64 = s
}

func BenchmarkRand_Intn(b *testing.B) {
	var s int
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Intn(small)
	}
	sinkInt = s
}

func BenchmarkRand_Intn_Big(b *testing.B) {
	var s int
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Intn(math.MaxInt - small)
	}
	sinkInt = s
}

func BenchmarkRand_NormFloat64(b *testing.B) {
	var s float64
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.NormFloat64()
	}
	sinkFloat64 = s
}

func BenchmarkRand_Perm(b *testing.B) {
	b.ReportAllocs()
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		r.Perm(tiny)
	}
}

func BenchmarkRand_Read(b *testing.B) {
	r := rand.New(1)
	p := make([]byte, 256)
	b.SetBytes(int64(len(p)))
	for i := 0; i < b.N; i++ {
		_, _ = r.Read(p[:])
	}
}

func BenchmarkRand_Seed(b *testing.B) {
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		r.Seed(uint64(i))
	}
}

func BenchmarkRand_Shuffle(b *testing.B) {
	r := rand.New(1)
	a := make([]int, tiny)
	for i := 0; i < b.N; i++ {
		r.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	}
}

func BenchmarkRand_ShuffleOverhead(b *testing.B) {
	r := rand.New(1)
	a := make([]int, tiny)
	for i := 0; i < b.N; i++ {
		r.Shuffle(len(a), func(i, j int) {})
	}
}

func BenchmarkRand_Uint32(b *testing.B) {
	var s uint32
	r := rand.New(1)
	b.SetBytes(4)
	for i := 0; i < b.N; i++ {
		s = r.Uint32()
	}
	sinkUint32 = s
}

func BenchmarkRand_Uint32n(b *testing.B) {
	var s uint32
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Uint32n(small)
	}
	sinkUint32 = s
}

func BenchmarkRand_Uint32n_Big(b *testing.B) {
	var s uint32
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Uint32n(math.MaxUint32 - small)
	}
	sinkUint32 = s
}

func BenchmarkRand_Uint64(b *testing.B) {
	var s uint64
	r := rand.New(1)
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		s = r.Uint64()
	}
	sinkUint64 = s
}

func BenchmarkRand_Uint64n(b *testing.B) {
	var s uint64
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Uint64n(small)
	}
	sinkUint64 = s
}

func BenchmarkRand_Uint64n_Big(b *testing.B) {
	var s uint64
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		s = r.Uint64n(math.MaxUint64 - small)
	}
	sinkUint64 = s
}

func BenchmarkRand_MarshalBinary(b *testing.B) {
	b.ReportAllocs()
	r := rand.New(1)
	for i := 0; i < b.N; i++ {
		_, _ = r.MarshalBinary()
	}
}

func BenchmarkRand_UnmarshalBinary(b *testing.B) {
	b.ReportAllocs()
	r := rand.New(1)
	buf, _ := r.MarshalBinary()
	for i := 0; i < b.N; i++ {
		_ = r.UnmarshalBinary(buf)
	}
}

func TestRand_Read(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		const N = 32
		s := rapid.Uint64().Draw(t, "s").(uint64)
		r := rand.New(s)
		buf := make([]byte, N)
		_, _ = r.Read(buf)
		r.Seed(s)
		buf2 := make([]byte, N)
		for n := 0; n < N; {
			c := rapid.IntRange(0, N-n).Draw(t, "c").(int)
			_, _ = r.Read(buf2[n : n+c])
			n += c
		}
		if !bytes.Equal(buf, buf2) {
			t.Fatalf("got %q instead of %q when reading in chunks", buf2, buf)
		}
	})
}

func TestRand_Float32(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		s := rapid.Uint64().Draw(t, "s").(uint64)
		r := rand.New(s)
		f := r.Float32()
		if f < 0 || f >= 1 {
			t.Fatalf("got %v outside of [0, 1)", f)
		}
	})
}

func TestRand_Float64(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		s := rapid.Uint64().Draw(t, "s").(uint64)
		r := rand.New(s)
		f := r.Float64()
		if f < 0 || f >= 1 {
			t.Fatalf("got %v outside of [0, 1)", f)
		}
	})
}

func TestRand_Int31n(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		s := rapid.Uint64().Draw(t, "s").(uint64)
		r := rand.New(s)
		n := rapid.Int32Range(1, math.MaxInt32).Draw(t, "n").(int32)
		v := r.Int31n(n)
		if v < 0 || v >= n {
			t.Fatalf("got %v outside of [0, %v)", v, n)
		}
	})
}

func TestRand_Int63n(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		s := rapid.Uint64().Draw(t, "s").(uint64)
		r := rand.New(s)
		n := rapid.Int64Range(1, math.MaxInt64).Draw(t, "n").(int64)
		v := r.Int63n(n)
		if v < 0 || v >= n {
			t.Fatalf("got %v outside of [0, %v)", v, n)
		}
	})
}

func TestRand_Intn(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		s := rapid.Uint64().Draw(t, "s").(uint64)
		r := rand.New(s)
		n := rapid.IntRange(1, math.MaxInt).Draw(t, "n").(int)
		v := r.Intn(n)
		if v < 0 || v >= n {
			t.Fatalf("got %v outside of [0, %v)", v, n)
		}
	})
}

func TestRand_Uint32n(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		s := rapid.Uint64().Draw(t, "s").(uint64)
		r := rand.New(s)
		n := rapid.Uint32Range(1, math.MaxUint32).Draw(t, "n").(uint32)
		v := r.Uint32n(n)
		if v < 0 || v >= n {
			t.Fatalf("got %v outside of [0, %v)", v, n)
		}
	})
}

func TestRand_Uint64n(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		s := rapid.Uint64().Draw(t, "s").(uint64)
		r := rand.New(s)
		n := rapid.Uint64Range(1, math.MaxUint64).Draw(t, "n").(uint64)
		v := r.Uint64n(n)
		if v < 0 || v >= n {
			t.Fatalf("got %v outside of [0, %v)", v, n)
		}
	})
}

func TestRand_MarshalBinary_Roundtrip(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		s := rapid.Uint64().Draw(t, "s").(uint64)
		r1 := rand.New(s)
		data1, err := r1.MarshalBinary()
		if err != nil {
			t.Fatalf("got unexpected marshal error: %v", err)
		}
		var r2 rand.Rand
		err = r2.UnmarshalBinary(data1)
		if err != nil {
			t.Fatalf("got unexpected unmarshal error: %v", err)
		}
		data2, err := r2.MarshalBinary()
		if err != nil {
			t.Fatalf("got unexpected marshal error: %v", err)
		}
		if !bytes.Equal(data1, data2) {
			t.Fatalf("data %q / %q after marshal/unmarshal", data1, data2)
		}
	})
}

func TestRand_Uint32nOpt(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		n := rapid.Uint32().Draw(t, "n").(uint32)
		v := rapid.Uint64().Draw(t, "v").(uint64)

		res, frac := bits.Mul32(n, uint32(v>>32))
		hi, _ := bits.Mul32(n, uint32(v))
		_, carry := bits.Add32(frac, hi, 0)
		res += carry

		res2, _ := bits.Mul64(uint64(n), v)

		if uint32(res2) != res {
			t.Fatalf("got %v instead of %v", res2, res)
		}
	})
}
