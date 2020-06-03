package assert

import (
	"fmt"
	"strings"
	"testing"
)

// That ...
func That(t *testing.T, got, expected interface{}) {
	a := fmt.Sprintf("%v", got)
	b := fmt.Sprintf("%v", expected)
	if !strings.EqualFold(a, b) {
		t.Errorf("got [%v] but expected [%v]", got, expected)
	}
}
