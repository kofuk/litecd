package config

import (
	"testing"
)

func TestExpandableStringVal(t *testing.T) {
	testcases := []struct {
		name         string
		rawVal       string
		secrets      map[string]string
		result       string
		expectsError bool
	}{
		{
			name:         "not a template",
			rawVal:       "foo",
			secrets:      map[string]string{},
			result:       "foo",
			expectsError: false,
		},
		{
			name:   "normal",
			rawVal: "{{ secret \"foo\" }}",
			secrets: map[string]string{
				"foo": "bar",
			},
			result:       "bar",
			expectsError: false,
		},
		{
			name:   "template syntax",
			rawVal: "{{ secret \"foo",
			secrets: map[string]string{
				"foo": "bar",
			},
			result:       "{{ secret \"foo",
			expectsError: false,
		},
		{
			name:         "unknown key",
			rawVal:       "{{ secret \"foo\" }}",
			secrets:      map[string]string{},
			result:       "",
			expectsError: false,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			instance := ExpandableString{rawVal: testcase.rawVal}
			result, err := instance.Val(testcase.secrets)
			if testcase.expectsError {
				if err == nil {
					t.Error("Error expected, but no error occurred")
				}
				return
			}
			if result != testcase.result {
				t.Errorf("expects %v, got %v\n", testcase.result, result)
			}
		})
	}
}
