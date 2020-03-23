// Copyright 2020 Shinichi MOTOKI. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package accbyte

var bytesTable = []byte{
	NUL, SOH, STX, ETX, EOT, ENQ, ACK, BEL, BS, HT, LF, VT, FF, CR, SO, SI, DLE,
	DC1, DC2, DC3, DC4, NAK, SYN, ETB, CAN, EM, SUB, ESC, FS, GS, RS, US,
	DEL,
}

// Bytes return byte slice of ASCII control character.
func Bytes(c byte) []byte {
	var s []byte

	switch {
	case c >= NUL && c <= US:
		s = bytesTable[c : c+1]
	case c == DEL:
		s = bytesTable[len(bytesTable)-1:]
	default:
	}

	return s
}
