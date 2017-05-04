package exp

type iterLMState struct {
	series seriesLM
	pos    int
	ts     uint32
}

func NewIterLMState(series seriesLM) *iterLMState {
	return &iterLMState{
		series: series,
		pos:    0,
		ts:     series.t0,
	}
}

func (i *iterLMState) Len() int {
	return len(i.series.data)
}

func (i *iterLMState) Step() uint32 {
	return i.series.step
}

func (i *iterLMState) Next() (uint32, float64, bool) {
	if i.pos == len(i.series.data) {
		return 0, 0, false
	}
	val := i.series.data[i.pos]
	ts := i.ts
	i.pos += 1
	i.ts += i.series.step
	return ts, val, true
}

func (i *iterLMState) Range() (uint32, uint32) {
	return i.series.t0, i.series.t0 + i.series.step*uint32(len(i.series.data)-1)
}
