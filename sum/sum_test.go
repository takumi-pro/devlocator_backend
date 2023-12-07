package sum_test

import (
	"testing"

	"devlocator/sum"

	"github.com/magiconair/properties/assert"
)

func TestSum_2つの数値の足し算(t *testing.T) {
	expected := 10
	actual := sum.Sum(3, 7)

	assert.Equal(t, expected, actual)
}
