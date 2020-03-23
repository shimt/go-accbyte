// Copyright 2020 Shinichi MOTOKI. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package accbyte

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBytes(t *testing.T) {
	tests := []struct {
		name string
		c    byte
		want []byte
	}{
		{"NUL", NUL, []byte{0x0}},
		{"SOH", SOH, []byte{0x1}},
		{"STX", STX, []byte{0x2}},
		{"ETX", ETX, []byte{0x3}},
		{"EOT", EOT, []byte{0x4}},
		{"ENQ", ENQ, []byte{0x5}},
		{"ACK", ACK, []byte{0x6}},
		{"BEL", BEL, []byte{0x7}},
		{"BS", BS, []byte{0x8}},
		{"HT", HT, []byte{0x9}},
		{"LF", LF, []byte{0xA}},
		{"VT", VT, []byte{0xB}},
		{"FF", FF, []byte{0xC}},
		{"CR", CR, []byte{0xD}},
		{"SO", SO, []byte{0xE}},
		{"SI", SI, []byte{0xF}},
		{"DLE", DLE, []byte{0x10}},
		{"DC1", DC1, []byte{0x11}},
		{"DC2", DC2, []byte{0x12}},
		{"DC3", DC3, []byte{0x13}},
		{"DC4", DC4, []byte{0x14}},
		{"NAK", NAK, []byte{0x15}},
		{"SYN", SYN, []byte{0x16}},
		{"ETB", ETB, []byte{0x17}},
		{"CAN", CAN, []byte{0x18}},
		{"EM", EM, []byte{0x19}},
		{"SUB", SUB, []byte{0x1A}},
		{"ESC", ESC, []byte{0x1B}},
		{"FS", FS, []byte{0x1C}},
		{"GS", GS, []byte{0x1D}},
		{"RS", RS, []byte{0x1E}},
		{"US", US, []byte{0x1F}},
		{"DEL", DEL, []byte{0x7F}},
		{"<undefine>", 0xFF, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bytes(tt.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleBytes() {
	s1 := Bytes(ESC)
	s2 := Bytes(HT)
	fmt.Printf("ESC = %v, HT = %v\n", s1, s2)

	s1 = Bytes(HT)
	s2 = Bytes(HT)
	fmt.Printf("(&s1 == &s2) = %v\n", &s1 == &s2)
	fmt.Printf("(&s1[0] == &s2[0]) = %v\n", &s1[0] == &s2[0])

	fmt.Printf("(Bytes(0xFF) == nil) = %v\n", Bytes(0xFF) == nil)

	// Output:
	// ESC = [27], HT = [9]
	// (&s1 == &s2) = false
	// (&s1[0] == &s2[0]) = true
	// (Bytes(0xFF) == nil) = true
}
