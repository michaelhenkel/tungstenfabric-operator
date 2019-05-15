// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
//go:generate go run gen.go gen_bits.go
=======
//go:generate go run gen.go
>>>>>>> v0.0.4

// Package runenames provides rune names from the Unicode Character Database.
// For example, the name for '\u0100' is "LATIN CAPITAL LETTER A WITH MACRON".
//
<<<<<<< HEAD
// See http://www.unicode.org/Public/UCD/latest/ucd/UnicodeData.txt
=======
// See https://www.unicode.org/Public/UCD/latest/ucd/UnicodeData.txt
>>>>>>> v0.0.4
package runenames

import (
	"sort"
)

// Name returns the name for r.
func Name(r rune) string {
<<<<<<< HEAD
	i := sort.Search(len(table0), func(j int) bool {
		e := table0[j]
		rOffset := rune(e >> shiftRuneOffset)
		return r < rOffset
=======
	i := sort.Search(len(entries), func(j int) bool {
		return entries[j].startRune() > r
>>>>>>> v0.0.4
	})
	if i == 0 {
		return ""
	}
<<<<<<< HEAD

	e := table0[i-1]
	rOffset := rune(e >> shiftRuneOffset)
	rLength := rune(e>>shiftRuneLength) & maskRuneLength
	if r >= rOffset+rLength {
		return ""
	}

	if (e>>shiftDirect)&maskDirect != 0 {
		o := int(e>>shiftDataOffset) & maskDataOffset
		n := int(e>>shiftDataLength) & maskDataLength
		return data[o : o+n]
	}

	base := uint32(e>>shiftDataBase) & maskDataBase
	base <<= dataBaseUnit
	j := rune(e>>shiftTable1Offset) & maskTable1Offset
	j += r - rOffset
	d0 := base + uint32(table1[j-1]) // dataOffset
	d1 := base + uint32(table1[j-0]) // dataOffset + dataLength
	return data[d0:d1]
}
=======
	e := entries[i-1]

	offset := int(r - e.startRune())
	if offset >= e.numRunes() {
		return ""
	}

	if e.direct() {
		o := e.index()
		n := e.len()
		return directData[o : o+n]
	}

	start := int(index[e.index()+offset])
	end := int(index[e.index()+offset+1])
	base1 := e.base() << 16
	base2 := base1
	if start > end {
		base2 += 1 << 16
	}
	return singleData[start+base1 : end+base2]
}

func (e entry) len() int { return e.base() }
>>>>>>> v0.0.4
