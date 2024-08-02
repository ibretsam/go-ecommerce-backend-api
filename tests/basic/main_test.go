package basic

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestAddOne(t *testing.T) {
	// var (
	// 	input  = 1
	// 	output = 2
	// )

	// actual := AddOne(1)
	// if actual != output {
	// 	t.Errorf("AddOne(%d): Expected the output to be %d but got %d", input, output, actual)
	// }

	assert.Equal(t, AddOne(1), 2, "AddOne(1): Expected the output to be 2")
}
