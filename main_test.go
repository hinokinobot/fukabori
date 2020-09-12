package main

import "testing"

func TestAdd(t *testing.T) {
	cases := map[string]struct {
		x, y, expected int
	}{
		"正数＋正数": {x: 2, y: 3, expected: 5},
		"正数＋負数": {x: 2, y: -3, expected: -1},
		"正数＋ゼロ": {x: 2, y: 0, expected: 2},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if actual := add(tc.x, tc.y); actual != tc.expected {
				t.Fatalf("expected: %v, actual: %v", tc.expected, actual)
			}
		})
	}
}
