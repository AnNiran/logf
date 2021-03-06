package logf

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRFC3339TimeEncoder(t *testing.T) {
	tm := time.Unix(1542266559, 305941000).UTC()
	enc := testTypeEncoder{}
	RFC3339TimeEncoder(tm, &enc)

	assert.EqualValues(t, "2018-11-15T07:22:39Z", enc.result)
}

func TestRFC3339NanoTimeEncoder(t *testing.T) {
	tm := time.Unix(1542266559, 305941000).UTC()
	enc := testTypeEncoder{}
	RFC3339NanoTimeEncoder(tm, &enc)

	assert.EqualValues(t, "2018-11-15T07:22:39.305941Z", enc.result)
}

func TestLayoutTimeEncoder(t *testing.T) {
	tm := time.Unix(1542266559, 305941000).UTC()
	enc := testTypeEncoder{}
	LayoutTimeEncoder(time.StampNano)(tm, &enc)

	assert.EqualValues(t, "Nov 15 07:22:39.305941000", enc.result)
}

func TestUnixNanoTimeEncoder(t *testing.T) {
	tm := time.Unix(1542266559, 305941000).UTC()
	enc := testTypeEncoder{}
	UnixNanoTimeEncoder(tm, &enc)

	assert.EqualValues(t, 1542266559305941000, enc.result)
}

func TestNanoDurationEncoder(t *testing.T) {
	d := time.Duration(66559305941000)
	enc := testTypeEncoder{}
	NanoDurationEncoder(d, &enc)

	assert.EqualValues(t, 66559305941000, enc.result)
}

func TestFloatSecondsDurationEncoder(t *testing.T) {
	d := time.Duration(66559305941000)
	enc := testTypeEncoder{}
	FloatSecondsDurationEncoder(d, &enc)

	assert.InDelta(t, 66559.305941, enc.result, 0.0000005)
}

func TestStringDurationEncoder(t *testing.T) {
	d := time.Duration(66559305941000)
	enc := testTypeEncoder{}
	StringDurationEncoder(d, &enc)

	assert.EqualValues(t, "18h29m19.305941s", enc.result)
}
