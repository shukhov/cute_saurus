package table

import (
	"fmt"
	"testing"
)

func TestExportFieldName(t *testing.T) {
	for i, testCase := range []struct {
		input, awaited string
	}{
		{"field", "Field"},
		{"separated_field", "SeparatedField"},
		{"", "Field"},
	} {
		t.Run(fmt.Sprintf("Test case num.: %d", i), func(t *testing.T) {
			awaited := testCase.awaited
			realField := exportFieldName(testCase.input)
			t.Logf("Awaited value: %s", awaited)
			t.Logf("Real value: %s", realField)

			if awaited != realField {
				t.Fatalf("Values is not equal: %s!=%s", awaited, realField)
			}
		})
	}
}
