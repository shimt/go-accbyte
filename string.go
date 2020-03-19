// Copyright 2020 Shinichi MOTOKI. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package accbyte provides definition for ASCII control characters.
package accbyte

var stringTable = []string{
	"NUL", "SOH", "STX", "ETX", "EOT", "ENQ", "ACK", "BEL", "BS", "HT", "LF", "VT", "FF", "CR", "SO", "SI", "DLE",
	"DC1", "DC2", "DC3", "DC4", "NAK", "SYN", "ETB", "CAN", "EM", "SUB", "ESC", "FS", "GS", "RS", "US",
	"DEL",
}

// String return display string of ASCII control character.
func String(c byte) string {
	var s string

	switch {
	case c >= NUL && c <= US:
		s = stringTable[c]
	case c == DEL:
		s = stringTable[len(stringTable)-1]
	default:
	}

	return s
}
