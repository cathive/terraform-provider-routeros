package routeros

import (
	"reflect"
	"testing"
)

func TestStringToBoolMap(t *testing.T) {
	tt := []struct{
		in string
		out map[string]bool
	}{
		{
			"str1,str2,!str3,str4,!str5",
			map[string]bool{
				"str1": true,
				"str2": true,
				"str3": false,
				"str4": true,
				"str5": false,
			},
		},
	}
	for _, tc := range tt {
		out := stringToBoolMap(tc.in)
		if !reflect.DeepEqual(out, tc.out) {
			t.Fail()
		}
	}
}

func TestBoolMapToString(t *testing.T) {
	tt := []struct{
		in map[string]bool
		out string
	}{
		{
			map[string]bool{
				"str1": true,
				"str2": true,
				"str3": false,
				"str4": true,
				"str5": false,
			},
			"!str3,!str5,str1,str2,str4",
		},
	}
	for _, tc := range tt {
		out := boolMapToString(tc.in)
		if !reflect.DeepEqual(out, tc.out) {
			t.Errorf("boolMapToString(...) -> %s, expected: %s", out, tc.out)
			t.Fail()
		}
	}
}
