package exp

type seriesMT []pointMT

type pointMT struct {
	ts  uint32
	val float64
}

func sumMT(in []seriesMT) seriesMT {
	var out seriesMT = make([]pointMT, len(in[0]))
	for p := 0; p < len(in[0]); p++ {
		point := in[0][p]

		for s := 1; s < len(in); s++ {
			point.val += in[s][p].val
		}
		out[p] = point
	}
	return out
}

func sumMTSeriesBySeries(in []seriesMT) seriesMT {
	var out seriesMT = make([]pointMT, len(in[0]))

	for p := 0; p < len(in[0]); p++ {
		out[p] = in[0][p]
	}
	for s := 1; s < len(in); s++ {
		for p := 0; p < len(in[s]); p++ {
			out[p].val += in[s][p].val
		}
	}
	return out
}
