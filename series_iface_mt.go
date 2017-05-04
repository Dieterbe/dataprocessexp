package exp

type iterMT struct {
	series seriesMT
	pos    int
}

func NewIterMT(series seriesMT) *iterMT {
	return &iterMT{
		series: series,
		pos:    0,
	}
}

func (i *iterMT) Len() int {
	return len(i.series)
}

func (i *iterMT) Step() uint32 {
	// series must have at least two points. also, assumed to have consistent step between points
	return i.series[1].ts - i.series[0].ts
}

func (i *iterMT) Next() (uint32, float64, bool) {
	if i.pos == len(i.series) {
		return 0, 0, false
	}
	p := i.series[i.pos]
	i.pos += 1
	return p.ts, p.val, true
}

func (i *iterMT) Range() (uint32, uint32) {
	return i.series[0].ts, i.series[len(i.series)-1].ts
}
