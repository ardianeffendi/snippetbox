package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	// Create a slice of anonymous structs containing the test case name,
	// input to humanDate() function (the tm field), and the expected output
	// (the want field).

	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2023, 2, 19, 6, 0, 0, 0, time.UTC),
			want: "19 Feb 2023 at 06:00",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2023, 2, 19, 6, 0, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "19 Feb 2023 at 05:00",
		},
	}

	for _, tt := range tests {
		// Use the t.Run() function to run a sub-test for each test case.
		// The first parameter to this is the name of the test (which is used to
		// identify the sub-test in any log output) and the second parameter is
		// and anonymous function containing the actual test for each case.
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			if hd != tt.want {
				t.Errorf("want %q; got %q", tt.want, hd)
			}
		})
	}
}
