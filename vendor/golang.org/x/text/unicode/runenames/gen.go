// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
<<<<<<< HEAD
	"log"
	"strings"
	"unicode"

	"golang.org/x/text/internal/gen"
	"golang.org/x/text/internal/ucd"
)

// snippet is a slice of data; data is the concatenation of all of the names.
type snippet struct {
	offset int
	length int
	s      string
}

func makeTable0EntryDirect(rOffset, rLength, dOffset, dLength int) uint64 {
	if rOffset >= 1<<bitsRuneOffset {
		log.Fatalf("makeTable0EntryDirect: rOffset %d is too large", rOffset)
	}
	if rLength >= 1<<bitsRuneLength {
		log.Fatalf("makeTable0EntryDirect: rLength %d is too large", rLength)
	}
	if dOffset >= 1<<bitsDataOffset {
		log.Fatalf("makeTable0EntryDirect: dOffset %d is too large", dOffset)
	}
	if dLength >= 1<<bitsRuneLength {
		log.Fatalf("makeTable0EntryDirect: dLength %d is too large", dLength)
	}
	return uint64(rOffset)<<shiftRuneOffset |
		uint64(rLength)<<shiftRuneLength |
		uint64(dOffset)<<shiftDataOffset |
		uint64(dLength)<<shiftDataLength |
		1 // Direct bit.
}

func makeTable0EntryIndirect(rOffset, rLength, dBase, t1Offset int) uint64 {
	if rOffset >= 1<<bitsRuneOffset {
		log.Fatalf("makeTable0EntryIndirect: rOffset %d is too large", rOffset)
	}
	if rLength >= 1<<bitsRuneLength {
		log.Fatalf("makeTable0EntryIndirect: rLength %d is too large", rLength)
	}
	if dBase >= 1<<bitsDataBase {
		log.Fatalf("makeTable0EntryIndirect: dBase %d is too large", dBase)
	}
	if t1Offset >= 1<<bitsTable1Offset {
		log.Fatalf("makeTable0EntryIndirect: t1Offset %d is too large", t1Offset)
	}
	return uint64(rOffset)<<shiftRuneOffset |
		uint64(rLength)<<shiftRuneLength |
		uint64(dBase)<<shiftDataBase |
		uint64(t1Offset)<<shiftTable1Offset |
		0 // Direct bit.
}

func makeTable1Entry(x int) uint16 {
	if x < 0 || 0xffff < x {
		log.Fatalf("makeTable1Entry: entry %d is out of range", x)
	}
	return uint16(x)
}

var (
	data     []byte
	snippets = make([]snippet, 1+unicode.MaxRune)
)

func main() {
	gen.Init()

	names, counts := parse()
	appendRepeatNames(names, counts)
	appendUniqueNames(names, counts)

	table0, table1 := makeTables()

	gen.Repackage("gen_bits.go", "bits.go", "runenames")

	w := gen.NewCodeWriter()
	w.WriteVar("table0", table0)
	w.WriteVar("table1", table1)
	w.WriteConst("data", string(data))
	w.WriteGoFile("tables.go", "runenames")
}

func parse() (names []string, counts map[string]int) {
	names = make([]string, 1+unicode.MaxRune)
	counts = map[string]int{}
	ucd.Parse(gen.OpenUCDFile("UnicodeData.txt"), func(p *ucd.Parser) {
		r, s := p.Rune(0), p.String(ucd.Name)
		if s == "" {
			return
		}
		if s[0] == '<' {
			const first = ", First>"
			if i := strings.Index(s, first); i >= 0 {
				s = s[:i] + ">"
			}
		}
		names[r] = s
		counts[s]++
	})
	return names, counts
}

func appendRepeatNames(names []string, counts map[string]int) {
	alreadySeen := map[string]snippet{}
	for r, s := range names {
		if s == "" || counts[s] == 1 {
			continue
		}
		if s[0] != '<' {
			log.Fatalf("Repeated name %q does not start with a '<'", s)
		}

		if z, ok := alreadySeen[s]; ok {
			snippets[r] = z
			continue
		}

		z := snippet{
			offset: len(data),
			length: len(s),
			s:      s,
		}
		data = append(data, s...)
		snippets[r] = z
		alreadySeen[s] = z
	}
}

func appendUniqueNames(names []string, counts map[string]int) {
	for r, s := range names {
		if s == "" || counts[s] != 1 {
			continue
		}
		if s[0] == '<' {
			log.Fatalf("Unique name %q starts with a '<'", s)
		}

		z := snippet{
			offset: len(data),
			length: len(s),
			s:      s,
		}
		data = append(data, s...)
		snippets[r] = z
	}
}

func makeTables() (table0 []uint64, table1 []uint16) {
	for i := 0; i < len(snippets); {
		zi := snippets[i]
		if zi == (snippet{}) {
			i++
			continue
		}

		// Look for repeat names. If we have one, we only need a table0 entry.
		j := i + 1
		for ; j < len(snippets) && zi == snippets[j]; j++ {
		}
		if j > i+1 {
			table0 = append(table0, makeTable0EntryDirect(i, j-i, zi.offset, zi.length))
			i = j
			continue
		}

		// Otherwise, we have a run of unique names. We need one table0 entry
		// and two or more table1 entries.
		base := zi.offset &^ (1<<dataBaseUnit - 1)
		t1Offset := len(table1) + 1
		table1 = append(table1, makeTable1Entry(zi.offset-base))
		table1 = append(table1, makeTable1Entry(zi.offset+zi.length-base))
		for ; j < len(snippets) && snippets[j] != (snippet{}); j++ {
			zj := snippets[j]
			if data[zj.offset] == '<' {
				break
			}
			table1 = append(table1, makeTable1Entry(zj.offset+zj.length-base))
		}
		table0 = append(table0, makeTable0EntryIndirect(i, j-i, base>>dataBaseUnit, t1Offset))
		i = j
	}
	return table0, table1
=======
	"bytes"
	"log"
	"sort"
	"strings"

	"golang.org/x/text/internal/gen"
	"golang.org/x/text/internal/gen/bitfield"
	"golang.org/x/text/internal/ucd"
)

var (
	// computed by computeDirectOffsets
	directOffsets = map[string]int{}
	directData    bytes.Buffer

	// computed by computeEntries
	entries    []entry
	singleData bytes.Buffer
	index      []uint16
)

type entry struct {
	start    rune `bitfield:"21,startRune"`
	numRunes int  `bitfield:"16"`
	end      rune
	index    int  `bitfield:"16"`
	base     int  `bitfield:"6"`
	direct   bool `bitfield:""`
	name     string
}

func main() {
	gen.Init()

	w := gen.NewCodeWriter()
	defer w.WriteVersionedGoFile("tables.go", "runenames")

	gen.WriteUnicodeVersion(w)

	computeDirectOffsets()
	computeEntries()

	if err := bitfield.Gen(w, entry{}, nil); err != nil {
		log.Fatal(err)
	}

	type entry uint64 // trick the generation code to use the entry type
	packed := []entry{}
	for _, e := range entries {
		e.numRunes = int(e.end - e.start + 1)
		v, err := bitfield.Pack(e, nil)
		if err != nil {
			log.Fatal(err)
		}
		packed = append(packed, entry(v))
	}

	index = append(index, uint16(singleData.Len()))

	w.WriteVar("entries", packed)
	w.WriteVar("index", index)
	w.WriteConst("directData", directData.String())
	w.WriteConst("singleData", singleData.String())
}

func computeDirectOffsets() {
	counts := map[string]int{}

	p := ucd.New(gen.OpenUCDFile("UnicodeData.txt"), ucd.KeepRanges)
	for p.Next() {
		start, end := p.Range(0)
		counts[getName(p)] += int(end-start) + 1
	}

	direct := []string{}
	for k, v := range counts {
		if v > 1 {
			direct = append(direct, k)
		}
	}
	sort.Strings(direct)

	for _, s := range direct {
		directOffsets[s] = directData.Len()
		directData.WriteString(s)
	}
}

func computeEntries() {
	p := ucd.New(gen.OpenUCDFile("UnicodeData.txt"), ucd.KeepRanges)
	for p.Next() {
		start, end := p.Range(0)

		last := entry{}
		if len(entries) > 0 {
			last = entries[len(entries)-1]
		}

		name := getName(p)
		if index, ok := directOffsets[name]; ok {
			if last.name == name && last.end+1 == start {
				entries[len(entries)-1].end = end
				continue
			}
			entries = append(entries, entry{
				start:  start,
				end:    end,
				index:  index,
				base:   len(name),
				direct: true,
				name:   name,
			})
			continue
		}

		if start != end {
			log.Fatalf("Expected start == end, found %x != %x", start, end)
		}

		offset := singleData.Len()
		base := offset >> 16
		index = append(index, uint16(offset))
		singleData.WriteString(name)

		if last.base == base && last.end+1 == start {
			entries[len(entries)-1].end = start
			continue
		}

		entries = append(entries, entry{
			start: start,
			end:   end,
			index: len(index) - 1,
			base:  base,
			name:  name,
		})
	}
}

func getName(p *ucd.Parser) string {
	s := p.String(ucd.Name)
	if s == "" {
		return ""
	}
	if s[0] == '<' {
		const first = ", First>"
		if i := strings.Index(s, first); i >= 0 {
			s = s[:i] + ">"
		}

	}
	return s
>>>>>>> v0.0.4
}
