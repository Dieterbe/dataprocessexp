package exp

type iface interface {
	Len() int
	Step() uint32
	Next() (uint32, float64, bool)
	Range() (uint32, uint32)
}

func sumIface(in []iface) seriesMT {
	var out seriesMT = make([]pointMT, in[0].Len())
	for p := 0; p < in[0].Len(); p++ {
		ts, val, _ := in[0].Next()

		for s := 1; s < len(in); s++ {
			_, v, _ := in[s].Next()
			val += v
		}
		out[p] = pointMT{ts, val}
	}
	return out
}
