// Copyright 2020 Shinichi MOTOKI. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package accbyte provides definition for ASCII control characters.
package accbyte

const (
	// NUL is ascii code of "Null".
	//
	// (Bin: 0b0000_0000 Oct: 0o0 Dec: 0 Hex: 0x00)
	NUL byte = 0x00
	// SOH is ascii code of "Start of Heading".
	//
	// (Bin: 0b0000_0001 Oct: 0o1 Dec: 1 Hex: 0x01)
	SOH = 0x01
	// STX is ascii code of "Start of Text".
	//
	// (Bin: 0b0000_0010 Oct: 0o2 Dec: 2 Hex: 0x02)
	STX = 0x02
	// ETX is ascii code of "End of Text".
	//
	// (Bin: 0b0000_0011 Oct: 0o3 Dec: 3 Hex: 0x03)
	ETX = 0x03
	// EOT is ascii code of "End of Transmission".
	//
	// (Bin: 0b0000_0100 Oct: 0o4 Dec: 4 Hex: 0x04)
	EOT = 0x04
	// ENQ is ascii code of "Enquiry".
	//
	// (Bin: 0b0000_0101 Oct: 0o5 Dec: 5 Hex: 0x05)
	ENQ = 0x05
	// ACK is ascii code of "Acknowledgement".
	//
	// (Bin: 0b0000_0110 Oct: 0o6 Dec: 6 Hex: 0x06)
	ACK = 0x06
	// BEL is ascii code of "Bell".
	//
	// (Bin: 0b0000_0111 Oct: 0o7 Dec: 7 Hex: 0x07)
	BEL = 0x07
	// BS is ascii code of "Backspace".
	//
	// (Bin: 0b0000_1000 Oct: 0o10 Dec: 8 Hex: 0x08)
	BS = 0x08
	// HT is ascii code of "Horizontal Tab".
	//
	// (Bin: 0b0000_1001 Oct: 0o11 Dec: 9 Hex: 0x09)
	HT = 0x09
	// LF is ascii code of "Line Feed".
	//
	// (Bin: 0b0000_1010 Oct: 0o12 Dec: 10 Hex: 0x0A)
	LF = 0x0A
	// VT is ascii code of "Vertical Tab".
	//
	// (Bin: 0b0000_1011 Oct: 0o13 Dec: 11 Hex: 0x0B)
	VT = 0x0B
	// FF is ascii code of "Form Feed".
	//
	// (Bin: 0b0000_1100 Oct: 0o14 Dec: 12 Hex: 0x0C)
	FF = 0x0C
	// CR is ascii code of "Carriage Return".
	//
	// (Bin: 0b0000_1101 Oct: 0o15 Dec: 13 Hex: 0x0D)
	CR = 0x0D
	// SO is ascii code of "Shift Out".
	//
	// (Bin: 0b0000_1110 Oct: 0o16 Dec: 14 Hex: 0x0E)
	SO = 0x0E
	// SI is ascii code of "Shift In".
	//
	// (Bin: 0b0000_1111 Oct: 0o17 Dec: 15 Hex: 0x0F)
	SI = 0x0F
	// DLE is ascii code of "Data Link Escape".
	//
	// (Bin: 0b0001_0000 Oct: 0o20 Dec: 16 Hex: 0x10)
	DLE = 0x10
	// DC1 is ascii code of "Device Control 1 (often XON)".
	//
	// (Bin: 0b0001_0001 Oct: 0o21 Dec: 17 Hex: 0x11)
	DC1 = 0x11
	// DC2 is ascii code of "Device Control 2".
	//
	// (Bin: 0b0001_0010 Oct: 0o22 Dec: 18 Hex: 0x12)
	DC2 = 0x12
	// DC3 is ascii code of "Device Control 3 (often XOFF)".
	//
	// (Bin: 0b0001_0011 Oct: 0o23 Dec: 19 Hex: 0x13)
	DC3 = 0x13
	// DC4 is ascii code of "Device Control 4".
	//
	// (Bin: 0b0001_0100 Oct: 0o24 Dec: 20 Hex: 0x14)
	DC4 = 0x14
	// NAK is ascii code of "Negative Acknowledgement".
	//
	// (Bin: 0b0001_0101 Oct: 0o25 Dec: 21 Hex: 0x15)
	NAK = 0x15
	// SYN is ascii code of "Synchronous Idle".
	//
	// (Bin: 0b0001_0110 Oct: 0o26 Dec: 22 Hex: 0x16)
	SYN = 0x16
	// ETB is ascii code of "End of Transmission Block".
	//
	// (Bin: 0b0001_0111 Oct: 0o27 Dec: 23 Hex: 0x17)
	ETB = 0x17
	// CAN is ascii code of "Cancel".
	//
	// (Bin: 0b0001_1000 Oct: 0o30 Dec: 24 Hex: 0x18)
	CAN = 0x18
	// EM is ascii code of "End of Medium".
	// (Bin: 0b0001_1001 Oct: 0o31 Dec: 25 Hex: 0x19)
	EM = 0x19
	// SUB is ascii code of "Substitute".
	//
	// (Bin: 0b0001_1010 Oct: 0o32 Dec: 26 Hex: 0x1A)
	SUB = 0x1A
	// ESC is ascii code of "Escape".
	//
	// (Bin: 0b0001_1011 Oct: 0o33 Dec: 27 Hex: 0x1B)
	ESC = 0x1B
	// FS is ascii code of "File Separator".
	//
	// (Bin: 0b0001_1100 Oct: 0o34 Dec: 28 Hex: 0x1C)
	FS = 0x1C
	// GS is ascii code of "Group Separator".
	//
	// (Bin: 0b0001_1101 Oct: 0o35 Dec: 29 Hex: 0x1D)
	GS = 0x1D
	// RS is ascii code of "Record Separator".
	//
	// (Bin: 0b0001_1110 Oct: 0o36 Dec: 30 Hex: 0x1E)
	RS = 0x1E
	// US is ascii code of "Unit Separator".
	//
	// (Bin: 0b0001_1111 Oct: 0o37 Dec: 31 Hex: 0x1F)
	US = 0x1F
	// DEL is ascii code of "Delete".
	//
	// (Bin: 0b0111_1111 Oct: 0o177 Dec: 127 Hex: 0x7F)
	DEL = 0x7F
)
