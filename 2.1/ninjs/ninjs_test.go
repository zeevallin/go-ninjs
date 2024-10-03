package ninjs_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// AssertEqualType asserts that the expected and actual values are of the same type
func AssertEqualType(t testing.TB, expected, actual any) {
	t.Helper()

	expectedT := fmt.Sprintf("%T", expected)
	actualT := fmt.Sprintf("%T", actual)

	if expectedT != actualT {
		t.Errorf("expected type (-want +got):\n\t- %s\n\t+ %s", expectedT, actualT)
		t.Logf("error: %v", actual)
	}
}

func DiffError(t testing.TB, expected, actual error) {
	t.Helper()

	if expected == nil && actual == nil {
		return
	}

	if expected == nil || actual == nil {
		t.Errorf("expected error (-want +got):\n\t- %v\n\t+ %v", expected, actual)
		return
	}

	if diff := cmp.Diff(expected.Error(), actual.Error()); diff != "" {
		t.Errorf("unexpected error (-want +got):\n%s", diff)
	}
}
