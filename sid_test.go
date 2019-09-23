package sid

import (
	"encoding/hex"
	"strings"
	"testing"
)

func decode(t *testing.T, hexVal string) []byte {
	decoded, err := hex.DecodeString(hexVal)
	if err != nil {
		t.Fatal(err)
	}
	return decoded
}

func TestConvertToString(t *testing.T) {
	tCases := map[string]struct {
		sid []byte
		expected,
		errMsg string
	}{
		"Null Authority": {
			decode(t, "0100000000000000"),
			"S-1-0",
			"",
		},
		"Nobody": {
			decode(t, "010100000000000000000000"),
			"S-1-0-0",
			"",
		},
		"Local": {
			decode(t, "010100000000000200000000"),
			"S-1-2-0",
			"",
		},
		"Anonymous": {
			decode(t, "010100000000000507000000"),
			"S-1-5-7",
			"",
		},
		"Example": {
			// https://en.wikipedia.org/wiki/Security_Identifier#Machine_SIDs
			decode(t, "0104000000000005150000002e43ac40c085385d07e53b2b"),
			"S-1-5-21-1085031214-1563985344-725345543",
			"",
		},
		"Too short": {
			decode(t, "01"),
			"",
			"Byte string given must have at least 8 bytes, but got only 1 bytes",
		},
		"Bad length": {
			decode(t, "010200000000000000000000"),
			"",
			"According to byte 1 of the SID it total length should be 16 bytes, however its actual length is 12 bytes",
		},
	}
	for name, tCase := range tCases {
		t.Run(name, func(t *testing.T) {
			actual, err := ConvertToString(tCase.sid)
			if err != nil {
				if tCase.errMsg == "" {
					t.Fatalf("expected to succeed but failed: %s", err.Error())
				}
				if !strings.Contains(err.Error(), tCase.errMsg) {
					t.Fatalf("expected error to contain %q but was %q", tCase.errMsg, err.Error())
				}
				return
			}
			if err == nil && tCase.errMsg != "" {
				t.Fatal("expected to fail but succeeded")
			}
			if actual != tCase.expected {
				t.Fatalf("expected sid to be %q but was %q", tCase.expected, actual)
			}
		})
	}
}
