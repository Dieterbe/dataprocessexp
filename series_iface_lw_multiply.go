package exp

type iterLMMultiply struct {
	series seriesLM
	pos    int
}

func NewIterLMMultiply(series seriesLM) *iterLMMultiply {
	return &iterLMMultiply{
		series: series,
		pos:    0,
	}
}

func (i *iterLMMultiply) Len() int {
	return len(i.series.data)
}

func (i *iterLMMultiply) Step() uint32 {
	return i.series.step
}

func (i *iterLMMultiply) Next() (uint32, float64, bool) {
	if i.pos == len(i.series.data) {
		return 0, 0, false
	}
	val := i.series.data[i.pos]
	ts := i.series.t0 + uint32(i.pos)*i.series.step
	i.pos += 1
	return ts, val, true
}

func (i *iterLMMultiply) Range() (uint32, uint32) {
	return i.series.t0, i.series.t0 + i.series.step*uint32(len(i.series.data)-1)
}
