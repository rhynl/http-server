package flags

import (
	"testing"
)

func TestGetFlags(t *testing.T) {
	for _, test := range flagsTestCases {
		argsv = test.input
		actual := GetFlags()

		if actual != test.expected {
			t.Errorf("GetFlags()\nexpected %v\nactual %v", test.expected, actual)
		}
	}
}
