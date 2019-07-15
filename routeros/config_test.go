package routeros

import "testing"

func TestNewConfig(t *testing.T) {
	var tt = []struct{
		inAddress string
		inPort string
		inUsername string
		inPassword string
		inTLSEnabled bool
		outValid bool
	}{
		{
			"localhost",
			"8728",
			"",
			"",
			false,
			true,
		},
		{
			"",
			"",
			"",
			"",
			false,
			false,
		},
	}
	for _, td := range tt {
		cfg, err := NewConfig(td.inAddress, td.inPort, td.inUsername, td.inPassword, td.inTLSEnabled)
		if td.outValid {
			if cfg == nil {
				t.Errorf("NewConfig(...) -> nil, ..., expected: non-nil value.")
			}
			if err != nil {
				t.Errorf("NewConfig(...) -> ..., non-nil, expected: no error in valid configuration.")
			}
		} else {
			if err == nil {
				t.Errorf("NewConfig(...) -> ...,nil, expected: non-nil error in invalid configuration.")
			}
		}
	}
}
