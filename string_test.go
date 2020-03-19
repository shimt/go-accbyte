// Copyright 2020 Shinichi MOTOKI. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package accbyte provides definition for ASCII control characters.
package accbyte

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		name string
		c    byte
		want string
	}{
		{"NUL", 0x0, "NUL"},
		{"SOH", 0x1, "SOH"},
		{"STX", 0x2, "STX"},
		{"ETX", 0x3, "ETX"},
		{"EOT", 0x4, "EOT"},
		{"ENQ", 0x5, "ENQ"},
		{"ACK", 0x6, "ACK"},
		{"BEL", 0x7, "BEL"},
		{"BS", 0x8, "BS"},
		{"HT", 0x9, "HT"},
		{"LF", 0x0A, "LF"},
		{"VT", 0x0B, "VT"},
		{"FF", 0x0C, "FF"},
		{"CR", 0x0D, "CR"},
		{"SO", 0x0E, "SO"},
		{"SI", 0x0F, "SI"},
		{"DLE", 0x10, "DLE"},
		{"DC1", 0x11, "DC1"},
		{"DC2", 0x12, "DC2"},
		{"DC3", 0x13, "DC3"},
		{"DC4", 0x14, "DC4"},
		{"NAK", 0x15, "NAK"},
		{"SYN", 0x16, "SYN"},
		{"ETB", 0x17, "ETB"},
		{"CAN", 0x18, "CAN"},
		{"EM", 0x19, "EM"},
		{"SUB", 0x1A, "SUB"},
		{"ESC", 0x1B, "ESC"},
		{"FS", 0x1C, "FS"},
		{"GS", 0x1D, "GS"},
		{"RS", 0x1E, "RS"},
		{"US", 0x1F, "US"},
		{"DEL", 0x7F, "DEL"},
		{"<undefine>", 0xFF, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.c); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleString() {
	s1 := String(ESC)
	s2 := String(HT)
	fmt.Printf("ESC = \"%s\", HT = \"%s\"\n", s1, s2)

	fmt.Printf("String(0xFF) == \"\") = %v\n", String(0xFF) == "")

	// Output:
	// ESC = "ESC", HT = "HT"
	// String(0xFF) == "") = true
}
