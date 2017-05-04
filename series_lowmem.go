package exp

type seriesLM struct {
	t0   uint32
	step uint32
	data []float64
}

func sumLM(in []seriesLM) seriesLM {
	out := seriesLM{
		t0:   in[0].t0,
		step: in[0].step,
		data: make([]float64, len(in[0].data)),
	}
	for p := 0; p < len(in[0].data); p++ {
		val := in[0].data[p]

		for s := 1; s < len(in); s++ {
			val += in[s].data[p]
		}

		out.data[p] = val
	}
	return out
}

func sumLMSeriesBySeries(in []seriesLM) seriesLM {
	out := seriesLM{
		t0:   in[0].t0,
		step: in[0].step,
		data: make([]float64, len(in[0].data)),
	}
	for p := 0; p < len(in[0].data); p++ {
		out.data[p] = in[0].data[p]
	}
	for s := 1; s < len(in); s++ {
		for p := 0; p < len(in[s].data); p++ {
			out.data[p] += in[s].data[p]
		}
	}
	return out
}
