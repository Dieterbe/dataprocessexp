package exp

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestSumMT(t *testing.T) {
	in := []seriesMT{
		[]pointMT{
			{0, 1.5},
			{10, 0.6},
		},
		[]pointMT{
			{0, 8.5},
			{10, 1.4},
		},
	}
	exp := seriesMT{
		pointMT{0, 10},
		pointMT{10, 2},
	}
	out := sumMT(in)
	if !reflect.DeepEqual(out, exp) {
		t.Fatalf("exp %v - got %v", exp, out)
	}

	out = sumMTSeriesBySeries(in)
	if !reflect.DeepEqual(out, exp) {
		t.Fatalf("exp %v - got %v", exp, out)
	}

}

func TestSumLM(t *testing.T) {
	in := []seriesLM{
		{
			0,
			10,
			[]float64{1.5, 0.6},
		},
		{
			0,
			10,
			[]float64{8.5, 1.4},
		},
	}
	exp := seriesLM{0, 10, []float64{10, 2}}
	out := sumLM(in)
	if !reflect.DeepEqual(out, exp) {
		t.Fatalf("exp %v - got %v", exp, out)
	}
	out = sumLMSeriesBySeries(in)
	if !reflect.DeepEqual(out, exp) {
		t.Fatalf("exp %v - got %v", exp, out)
	}

}

func TestSumIface(t *testing.T) {
	in := []iface{
		NewIterMT(seriesMT([]pointMT{
			{0, 1.5},
			{10, 0.6},
		})),
		NewIterLMState(seriesLM{
			0,
			10,
			[]float64{8.5, 1.4},
		}),
		NewIterLMMultiply(seriesLM{
			0,
			10,
			[]float64{3.5, 2.1},
		}),
	}
	exp := seriesMT{
		pointMT{0, 13.5},
		pointMT{10, 4.1},
	}
	out := sumIface(in)
	if !reflect.DeepEqual(out, exp) {
		t.Fatalf("exp %v - got %v", exp, out)
	}

}

var outMT seriesMT
var outLM seriesLM

func BenchmarkSumMT(b *testing.B) {
	numSeries := 5
	in := make([]seriesMT, numSeries)
	for i := 0; i < numSeries; i++ {
		in[i] = make([]pointMT, b.N)
		for j := 0; j < b.N; j++ {
			in[i][j] = pointMT{
				uint32(j),
				rand.Float64(),
			}
		}
	}
	b.ResetTimer()
	outMT = sumMT(in)
}

func BenchmarkSumMTSeriesBySeries(b *testing.B) {
	numSeries := 5
	in := make([]seriesMT, numSeries)
	for i := 0; i < numSeries; i++ {
		in[i] = make([]pointMT, b.N)
		for j := 0; j < b.N; j++ {
			in[i][j] = pointMT{
				uint32(j),
				rand.Float64(),
			}
		}
	}
	b.ResetTimer()
	outMT = sumMTSeriesBySeries(in)
}

func BenchmarkSumLM(b *testing.B) {
	numSeries := 5
	in := make([]seriesLM, numSeries)
	for i := 0; i < numSeries; i++ {
		in[i] = seriesLM{
			0,
			1,
			make([]float64, b.N),
		}
		for j := 0; j < b.N; j++ {
			in[i].data[j] = rand.Float64()
		}
	}
	b.ResetTimer()
	outLM = sumLM(in)
}

func BenchmarkSumLMSeriesBySeries(b *testing.B) {
	numSeries := 5
	in := make([]seriesLM, numSeries)
	for i := 0; i < numSeries; i++ {
		in[i] = seriesLM{
			0,
			1,
			make([]float64, b.N),
		}
		for j := 0; j < b.N; j++ {
			in[i].data[j] = rand.Float64()
		}
	}
	b.ResetTimer()
	outLM = sumLMSeriesBySeries(in)
}

func BenchmarkSumIfaceMT(b *testing.B) {
	numSeries := 5
	in := make([]iface, numSeries)
	for i := 0; i < numSeries; i++ {
		s := make([]pointMT, b.N)
		for j := 0; j < b.N; j++ {
			s[j] = pointMT{
				uint32(j),
				rand.Float64(),
			}
		}
		in[i] = NewIterMT(s)
	}
	b.ResetTimer()
	outMT = sumIface(in)
}

func BenchmarkSumIfaceLMState(b *testing.B) {
	numSeries := 5
	in := make([]iface, numSeries)
	for i := 0; i < numSeries; i++ {
		s := seriesLM{
			0,
			1,
			make([]float64, b.N),
		}
		for j := 0; j < b.N; j++ {
			s.data[j] = rand.Float64()
		}
		in[i] = NewIterLMState(s)
	}
	b.ResetTimer()
	outMT = sumIface(in)
}

func BenchmarkSumIfaceLMMultiply(b *testing.B) {
	numSeries := 5
	in := make([]iface, numSeries)
	for i := 0; i < numSeries; i++ {
		s := seriesLM{
			0,
			1,
			make([]float64, b.N),
		}
		for j := 0; j < b.N; j++ {
			s.data[j] = rand.Float64()
		}
		in[i] = NewIterLMMultiply(s)
	}
	b.ResetTimer()
	outMT = sumIface(in)
}
