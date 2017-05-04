comparing 3 main approaches of datastructures:

* MT style (array of points, each point having val and ts). for evenly spaced data, Low-memory style (see below) may be better. this shines more for sparse data.
* Low memory style (tracks timestamp of first point and step, and an array of just values), we don't need to specify timestamp of each point if points are evenly spaced
* interface: an abstraction above the two above approaches.  The idea is to allow flexibility in the processing library for callers (metrictank, grafana, ...) to use whichever datatype they prefer, as long as it matches a simple interface. For LM-style we actually compare 2 slightly different approaches of implementations for the interface


Here, the benchmarks just use a single "sum" call with 5 series as input. This just represents a procesing step. since summing itself is pretty cheap (compared to forecasts, derivatives, etc), this highlights the cost of the different datastructure more. In practice you'd have possibly many such steps in sequence, the cost would then just multiply, so this is a good approximation of a single processing step.


NOT included:
* handling of NaN values, since that concern should be same across different implementations


## dieters conclusions:


```
go test -bench=.
BenchmarkSumMT-8                 	200000000	         7.75 ns/op
BenchmarkSumMTSeriesBySeries-8   	100000000	        12.9 ns/op
BenchmarkSumLM-8                 	300000000	         5.32 ns/op
BenchmarkSumLMSeriesBySeries-8   	200000000	         7.07 ns/op
BenchmarkSumIfaceMT-8            	50000000	        25.3 ns/op
BenchmarkSumIfaceLMState-8       	50000000	        26.5 ns/op
BenchmarkSumIfaceLMMultiply-8    	50000000	        25.3 ns/op
PASS
ok  	github.com/raintank/dataprocessexp	170.445s
```

* using an interface results in too much overhead.  Remember the benchmarks are per point. Let's say you have requests processing 400 series of each 10k points, and you have 5 functions in a row. do many requsets at the same time, and the spent cpu time quickly adds up!
* LM style is most efficient, in memory and cpu.  We should base our processing api's on this datatype. Data fed into the library by callers such as grafana or metrictank should be in this format.  applying this style also internally in those applications is recommended, but optional as far as the processing library is concerned.
* While MT stores chunks as a sequence of (ts, val) pairs (allowing for sparse data),
  in practice as soon as we retrieve any chunks, we quantize this data anyway to a consistent step, and with nulls filled in, so all internal data assumes quantized (non-sparse) data.
  - in all MT code that comes after decoding a chunk (quantizing and so on) we can also adopt the LM style.
  - in the future we could adopt a new chunk format that does not store timestamps for each point, but rather stores nulls in between points. assuming gaps are infrequent, this will be more efficient (chunk needs less space), and also means we don't need to quantize data after decoding the chunk.
