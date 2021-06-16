package internal_test

import (
	"github.com/crisaltmann/fundament-stock-api/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RoundNumbers(t *testing.T) {
	assert.EqualValues(t, 12.52, internal.RoundFloat(float32(12.52112)))
	assert.EqualValues(t, 16.53, internal.RoundFloat(float32(16.52912)))
	assert.EqualValues(t, 1.07, internal.RoundFloat(float32(1.066666666666667)))
}