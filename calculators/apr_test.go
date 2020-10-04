package calculators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculateAPR(t *testing.T) {
	expected := APRInfo{
		17.05,
		60.00 * 17.05,
		(60.00 * 17.05) - 1000.00 - 0.00,
	}

	actual := APR(1000.00, 60.00, 0.009, 0.00)
	assert.Equal(t, expected, actual, "Monthly payment is wromg")
}
