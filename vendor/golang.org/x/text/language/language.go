// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
//go:generate go run gen.go gen_common.go -output tables.go
//go:generate go run gen_index.go
=======
//go:generate go run gen.go -output tables.go
>>>>>>> v0.0.4

package language

// TODO: Remove above NOTE after:
// - verifying that tables are dropped correctly (most notably matcher tables).

import (
<<<<<<< HEAD
	"errors"
	"fmt"
	"strings"
)

const (
	// maxCoreSize is the maximum size of a BCP 47 tag without variants and
	// extensions. Equals max lang (3) + script (4) + max reg (3) + 2 dashes.
	maxCoreSize = 12

	// max99thPercentileSize is a somewhat arbitrary buffer size that presumably
	// is large enough to hold at least 99% of the BCP 47 tags.
	max99thPercentileSize = 32

	// maxSimpleUExtensionSize is the maximum size of a -u extension with one
	// key-type pair. Equals len("-u-") + key (2) + dash + max value (8).
	maxSimpleUExtensionSize = 14
=======
	"strings"

	"golang.org/x/text/internal/language"
	"golang.org/x/text/internal/language/compact"
>>>>>>> v0.0.4
)

// Tag represents a BCP 47 language tag. It is used to specify an instance of a
// specific language or locale. All language tag values are guaranteed to be
// well-formed.
<<<<<<< HEAD
type Tag struct {
	lang   langID
	region regionID
	// TODO: we will soon run out of positions for script. Idea: instead of
	// storing lang, region, and script codes, store only the compact index and
	// have a lookup table from this code to its expansion. This greatly speeds
	// up table lookup, speed up common variant cases.
	// This will also immediately free up 3 extra bytes. Also, the pVariant
	// field can now be moved to the lookup table, as the compact index uniquely
	// determines the offset of a possible variant.
	script   scriptID
	pVariant byte   // offset in str, includes preceding '-'
	pExt     uint16 // offset of first extension, includes preceding '-'

	// str is the string representation of the Tag. It will only be used if the
	// tag has variants or extensions.
	str string
}

=======
type Tag compact.Tag

func makeTag(t language.Tag) (tag Tag) {
	return Tag(compact.Make(t))
}

func (t *Tag) tag() language.Tag {
	return (*compact.Tag)(t).Tag()
}

func (t *Tag) isCompact() bool {
	return (*compact.Tag)(t).IsCompact()
}

// TODO: improve performance.
func (t *Tag) lang() language.Language { return t.tag().LangID }
func (t *Tag) region() language.Region { return t.tag().RegionID }
func (t *Tag) script() language.Script { return t.tag().ScriptID }

>>>>>>> v0.0.4
// Make is a convenience wrapper for Parse that omits the error.
// In case of an error, a sensible default is returned.
func Make(s string) Tag {
	return Default.Make(s)
}

// Make is a convenience wrapper for c.Parse that omits the error.
// In case of an error, a sensible default is returned.
func (c CanonType) Make(s string) Tag {
	t, _ := c.Parse(s)
	return t
}

// Raw returns the raw base language, script and region, without making an
// attempt to infer their values.
func (t Tag) Raw() (b Base, s Script, r Region) {
<<<<<<< HEAD
	return Base{t.lang}, Script{t.script}, Region{t.region}
}

// equalTags compares language, script and region subtags only.
func (t Tag) equalTags(a Tag) bool {
	return t.lang == a.lang && t.script == a.script && t.region == a.region
=======
	tt := t.tag()
	return Base{tt.LangID}, Script{tt.ScriptID}, Region{tt.RegionID}
>>>>>>> v0.0.4
}

// IsRoot returns true if t is equal to language "und".
func (t Tag) IsRoot() bool {
<<<<<<< HEAD
	if int(t.pVariant) < len(t.str) {
		return false
	}
	return t.equalTags(und)
}

// private reports whether the Tag consists solely of a private use tag.
func (t Tag) private() bool {
	return t.str != "" && t.pVariant == 0
=======
	return compact.Tag(t).IsRoot()
>>>>>>> v0.0.4
}

// CanonType can be used to enable or disable various types of canonicalization.
type CanonType int

const (
	// Replace deprecated base languages with their preferred replacements.
	DeprecatedBase CanonType = 1 << iota
	// Replace deprecated scripts with their preferred replacements.
	DeprecatedScript
	// Replace deprecated regions with their preferred replacements.
	DeprecatedRegion
	// Remove redundant scripts.
	SuppressScript
	// Normalize legacy encodings. This includes legacy languages defined in
	// CLDR as well as bibliographic codes defined in ISO-639.
	Legacy
	// Map the dominant language of a macro language group to the macro language
	// subtag. For example cmn -> zh.
	Macro
	// The CLDR flag should be used if full compatibility with CLDR is required.
	// There are a few cases where language.Tag may differ from CLDR. To follow all
	// of CLDR's suggestions, use All|CLDR.
	CLDR

	// Raw can be used to Compose or Parse without Canonicalization.
	Raw CanonType = 0

	// Replace all deprecated tags with their preferred replacements.
	Deprecated = DeprecatedBase | DeprecatedScript | DeprecatedRegion

	// All canonicalizations recommended by BCP 47.
	BCP47 = Deprecated | SuppressScript

	// All canonicalizations.
	All = BCP47 | Legacy | Macro

	// Default is the canonicalization used by Parse, Make and Compose. To
	// preserve as much information as possible, canonicalizations that remove
	// potentially valuable information are not included. The Matcher is
	// designed to recognize similar tags that would be the same if
	// they were canonicalized using All.
	Default = Deprecated | Legacy

	canonLang = DeprecatedBase | Legacy | Macro

	// TODO: LikelyScript, LikelyRegion: suppress similar to ICU.
)

// canonicalize returns the canonicalized equivalent of the tag and
// whether there was any change.
<<<<<<< HEAD
func (t Tag) canonicalize(c CanonType) (Tag, bool) {
=======
func canonicalize(c CanonType, t language.Tag) (language.Tag, bool) {
>>>>>>> v0.0.4
	if c == Raw {
		return t, false
	}
	changed := false
	if c&SuppressScript != 0 {
<<<<<<< HEAD
		if t.lang < langNoIndexOffset && uint8(t.script) == suppressScript[t.lang] {
			t.script = 0
=======
		if t.LangID.SuppressScript() == t.ScriptID {
			t.ScriptID = 0
>>>>>>> v0.0.4
			changed = true
		}
	}
	if c&canonLang != 0 {
		for {
<<<<<<< HEAD
			if l, aliasType := normLang(t.lang); l != t.lang {
				switch aliasType {
				case langLegacy:
					if c&Legacy != 0 {
						if t.lang == _sh && t.script == 0 {
							t.script = _Latn
						}
						t.lang = l
						changed = true
					}
				case langMacro:
=======
			if l, aliasType := t.LangID.Canonicalize(); l != t.LangID {
				switch aliasType {
				case language.Legacy:
					if c&Legacy != 0 {
						if t.LangID == _sh && t.ScriptID == 0 {
							t.ScriptID = _Latn
						}
						t.LangID = l
						changed = true
					}
				case language.Macro:
>>>>>>> v0.0.4
					if c&Macro != 0 {
						// We deviate here from CLDR. The mapping "nb" -> "no"
						// qualifies as a typical Macro language mapping.  However,
						// for legacy reasons, CLDR maps "no", the macro language
						// code for Norwegian, to the dominant variant "nb". This
						// change is currently under consideration for CLDR as well.
<<<<<<< HEAD
						// See http://unicode.org/cldr/trac/ticket/2698 and also
						// http://unicode.org/cldr/trac/ticket/1790 for some of the
						// practical implications. TODO: this check could be removed
						// if CLDR adopts this change.
						if c&CLDR == 0 || t.lang != _nb {
							changed = true
							t.lang = l
						}
					}
				case langDeprecated:
					if c&DeprecatedBase != 0 {
						if t.lang == _mo && t.region == 0 {
							t.region = _MD
						}
						t.lang = l
=======
						// See https://unicode.org/cldr/trac/ticket/2698 and also
						// https://unicode.org/cldr/trac/ticket/1790 for some of the
						// practical implications. TODO: this check could be removed
						// if CLDR adopts this change.
						if c&CLDR == 0 || t.LangID != _nb {
							changed = true
							t.LangID = l
						}
					}
				case language.Deprecated:
					if c&DeprecatedBase != 0 {
						if t.LangID == _mo && t.RegionID == 0 {
							t.RegionID = _MD
						}
						t.LangID = l
>>>>>>> v0.0.4
						changed = true
						// Other canonicalization types may still apply.
						continue
					}
				}
<<<<<<< HEAD
			} else if c&Legacy != 0 && t.lang == _no && c&CLDR != 0 {
				t.lang = _nb
=======
			} else if c&Legacy != 0 && t.LangID == _no && c&CLDR != 0 {
				t.LangID = _nb
>>>>>>> v0.0.4
				changed = true
			}
			break
		}
	}
	if c&DeprecatedScript != 0 {
<<<<<<< HEAD
		if t.script == _Qaai {
			changed = true
			t.script = _Zinh
		}
	}
	if c&DeprecatedRegion != 0 {
		if r := normRegion(t.region); r != 0 {
			changed = true
			t.region = r
=======
		if t.ScriptID == _Qaai {
			changed = true
			t.ScriptID = _Zinh
		}
	}
	if c&DeprecatedRegion != 0 {
		if r := t.RegionID.Canonicalize(); r != t.RegionID {
			changed = true
			t.RegionID = r
>>>>>>> v0.0.4
		}
	}
	return t, changed
}

// Canonicalize returns the canonicalized equivalent of the tag.
func (c CanonType) Canonicalize(t Tag) (Tag, error) {
<<<<<<< HEAD
	t, changed := t.canonicalize(c)
	if changed {
		t.remakeString()
	}
	return t, nil
=======
	// First try fast path.
	if t.isCompact() {
		if _, changed := canonicalize(c, compact.Tag(t).Tag()); !changed {
			return t, nil
		}
	}
	// It is unlikely that one will canonicalize a tag after matching. So do
	// a slow but simple approach here.
	if tag, changed := canonicalize(c, t.tag()); changed {
		tag.RemakeString()
		return makeTag(tag), nil
	}
	return t, nil

>>>>>>> v0.0.4
}

// Confidence indicates the level of certainty for a given return value.
// For example, Serbian may be written in Cyrillic or Latin script.
// The confidence level indicates whether a value was explicitly specified,
// whether it is typically the only possible value, or whether there is
// an ambiguity.
type Confidence int

const (
	No    Confidence = iota // full confidence that there was no match
	Low                     // most likely value picked out of a set of alternatives
	High                    // value is generally assumed to be the correct match
	Exact                   // exact match or explicitly specified value
)

var confName = []string{"No", "Low", "High", "Exact"}

func (c Confidence) String() string {
	return confName[c]
}

<<<<<<< HEAD
// remakeString is used to update t.str in case lang, script or region changed.
// It is assumed that pExt and pVariant still point to the start of the
// respective parts.
func (t *Tag) remakeString() {
	if t.str == "" {
		return
	}
	extra := t.str[t.pVariant:]
	if t.pVariant > 0 {
		extra = extra[1:]
	}
	if t.equalTags(und) && strings.HasPrefix(extra, "x-") {
		t.str = extra
		t.pVariant = 0
		t.pExt = 0
		return
	}
	var buf [max99thPercentileSize]byte // avoid extra memory allocation in most cases.
	b := buf[:t.genCoreBytes(buf[:])]
	if extra != "" {
		diff := len(b) - int(t.pVariant)
		b = append(b, '-')
		b = append(b, extra...)
		t.pVariant = uint8(int(t.pVariant) + diff)
		t.pExt = uint16(int(t.pExt) + diff)
	} else {
		t.pVariant = uint8(len(b))
		t.pExt = uint16(len(b))
	}
	t.str = string(b)
}

// genCoreBytes writes a string for the base languages, script and region tags
// to the given buffer and returns the number of bytes written. It will never
// write more than maxCoreSize bytes.
func (t *Tag) genCoreBytes(buf []byte) int {
	n := t.lang.stringToBuf(buf[:])
	if t.script != 0 {
		n += copy(buf[n:], "-")
		n += copy(buf[n:], t.script.String())
	}
	if t.region != 0 {
		n += copy(buf[n:], "-")
		n += copy(buf[n:], t.region.String())
	}
	return n
}

// String returns the canonical string representation of the language tag.
func (t Tag) String() string {
	if t.str != "" {
		return t.str
	}
	if t.script == 0 && t.region == 0 {
		return t.lang.String()
	}
	buf := [maxCoreSize]byte{}
	return string(buf[:t.genCoreBytes(buf[:])])
=======
// String returns the canonical string representation of the language tag.
func (t Tag) String() string {
	return t.tag().String()
>>>>>>> v0.0.4
}

// MarshalText implements encoding.TextMarshaler.
func (t Tag) MarshalText() (text []byte, err error) {
<<<<<<< HEAD
	if t.str != "" {
		text = append(text, t.str...)
	} else if t.script == 0 && t.region == 0 {
		text = append(text, t.lang.String()...)
	} else {
		buf := [maxCoreSize]byte{}
		text = buf[:t.genCoreBytes(buf[:])]
	}
	return text, nil
=======
	return t.tag().MarshalText()
>>>>>>> v0.0.4
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (t *Tag) UnmarshalText(text []byte) error {
<<<<<<< HEAD
	tag, err := Raw.Parse(string(text))
	*t = tag
=======
	var tag language.Tag
	err := tag.UnmarshalText(text)
	*t = makeTag(tag)
>>>>>>> v0.0.4
	return err
}

// Base returns the base language of the language tag. If the base language is
// unspecified, an attempt will be made to infer it from the context.
// It uses a variant of CLDR's Add Likely Subtags algorithm. This is subject to change.
func (t Tag) Base() (Base, Confidence) {
<<<<<<< HEAD
	if t.lang != 0 {
		return Base{t.lang}, Exact
	}
	c := High
	if t.script == 0 && !(Region{t.region}).IsCountry() {
		c = Low
	}
	if tag, err := addTags(t); err == nil && tag.lang != 0 {
		return Base{tag.lang}, c
=======
	if b := t.lang(); b != 0 {
		return Base{b}, Exact
	}
	tt := t.tag()
	c := High
	if tt.ScriptID == 0 && !tt.RegionID.IsCountry() {
		c = Low
	}
	if tag, err := tt.Maximize(); err == nil && tag.LangID != 0 {
		return Base{tag.LangID}, c
>>>>>>> v0.0.4
	}
	return Base{0}, No
}

// Script infers the script for the language tag. If it was not explicitly given, it will infer
// a most likely candidate.
// If more than one script is commonly used for a language, the most likely one
// is returned with a low confidence indication. For example, it returns (Cyrl, Low)
// for Serbian.
// If a script cannot be inferred (Zzzz, No) is returned. We do not use Zyyy (undetermined)
// as one would suspect from the IANA registry for BCP 47. In a Unicode context Zyyy marks
// common characters (like 1, 2, 3, '.', etc.) and is therefore more like multiple scripts.
<<<<<<< HEAD
// See http://www.unicode.org/reports/tr24/#Values for more details. Zzzz is also used for
=======
// See https://www.unicode.org/reports/tr24/#Values for more details. Zzzz is also used for
>>>>>>> v0.0.4
// unknown value in CLDR.  (Zzzz, Exact) is returned if Zzzz was explicitly specified.
// Note that an inferred script is never guaranteed to be the correct one. Latin is
// almost exclusively used for Afrikaans, but Arabic has been used for some texts
// in the past.  Also, the script that is commonly used may change over time.
// It uses a variant of CLDR's Add Likely Subtags algorithm. This is subject to change.
func (t Tag) Script() (Script, Confidence) {
<<<<<<< HEAD
	if t.script != 0 {
		return Script{t.script}, Exact
	}
	sc, c := scriptID(_Zzzz), No
	if t.lang < langNoIndexOffset {
		if scr := scriptID(suppressScript[t.lang]); scr != 0 {
			// Note: it is not always the case that a language with a suppress
			// script value is only written in one script (e.g. kk, ms, pa).
			if t.region == 0 {
				return Script{scriptID(scr)}, High
			}
			sc, c = scr, High
		}
	}
	if tag, err := addTags(t); err == nil {
		if tag.script != sc {
			sc, c = tag.script, Low
		}
	} else {
		t, _ = (Deprecated | Macro).Canonicalize(t)
		if tag, err := addTags(t); err == nil && tag.script != sc {
			sc, c = tag.script, Low
=======
	if scr := t.script(); scr != 0 {
		return Script{scr}, Exact
	}
	tt := t.tag()
	sc, c := language.Script(_Zzzz), No
	if scr := tt.LangID.SuppressScript(); scr != 0 {
		// Note: it is not always the case that a language with a suppress
		// script value is only written in one script (e.g. kk, ms, pa).
		if tt.RegionID == 0 {
			return Script{scr}, High
		}
		sc, c = scr, High
	}
	if tag, err := tt.Maximize(); err == nil {
		if tag.ScriptID != sc {
			sc, c = tag.ScriptID, Low
		}
	} else {
		tt, _ = canonicalize(Deprecated|Macro, tt)
		if tag, err := tt.Maximize(); err == nil && tag.ScriptID != sc {
			sc, c = tag.ScriptID, Low
>>>>>>> v0.0.4
		}
	}
	return Script{sc}, c
}

// Region returns the region for the language tag. If it was not explicitly given, it will
// infer a most likely candidate from the context.
// It uses a variant of CLDR's Add Likely Subtags algorithm. This is subject to change.
func (t Tag) Region() (Region, Confidence) {
<<<<<<< HEAD
	if t.region != 0 {
		return Region{t.region}, Exact
	}
	if t, err := addTags(t); err == nil {
		return Region{t.region}, Low // TODO: differentiate between high and low.
	}
	t, _ = (Deprecated | Macro).Canonicalize(t)
	if tag, err := addTags(t); err == nil {
		return Region{tag.region}, Low
=======
	if r := t.region(); r != 0 {
		return Region{r}, Exact
	}
	tt := t.tag()
	if tt, err := tt.Maximize(); err == nil {
		return Region{tt.RegionID}, Low // TODO: differentiate between high and low.
	}
	tt, _ = canonicalize(Deprecated|Macro, tt)
	if tag, err := tt.Maximize(); err == nil {
		return Region{tag.RegionID}, Low
>>>>>>> v0.0.4
	}
	return Region{_ZZ}, No // TODO: return world instead of undetermined?
}

<<<<<<< HEAD
// Variant returns the variants specified explicitly for this language tag.
// or nil if no variant was specified.
func (t Tag) Variants() []Variant {
	v := []Variant{}
	if int(t.pVariant) < int(t.pExt) {
		for x, str := "", t.str[t.pVariant:t.pExt]; str != ""; {
			x, str = nextToken(str)
			v = append(v, Variant{x})
		}
=======
// Variants returns the variants specified explicitly for this language tag.
// or nil if no variant was specified.
func (t Tag) Variants() []Variant {
	if !compact.Tag(t).MayHaveVariants() {
		return nil
	}
	v := []Variant{}
	x, str := "", t.tag().Variants()
	for str != "" {
		x, str = nextToken(str)
		v = append(v, Variant{x})
>>>>>>> v0.0.4
	}
	return v
}

// Parent returns the CLDR parent of t. In CLDR, missing fields in data for a
// specific language are substituted with fields from the parent language.
// The parent for a language may change for newer versions of CLDR.
<<<<<<< HEAD
func (t Tag) Parent() Tag {
	if t.str != "" {
		// Strip the variants and extensions.
		t, _ = Raw.Compose(t.Raw())
		if t.region == 0 && t.script != 0 && t.lang != 0 {
			base, _ := addTags(Tag{lang: t.lang})
			if base.script == t.script {
				return Tag{lang: t.lang}
			}
		}
		return t
	}
	if t.lang != 0 {
		if t.region != 0 {
			maxScript := t.script
			if maxScript == 0 {
				max, _ := addTags(t)
				maxScript = max.script
			}

			for i := range parents {
				if langID(parents[i].lang) == t.lang && scriptID(parents[i].maxScript) == maxScript {
					for _, r := range parents[i].fromRegion {
						if regionID(r) == t.region {
							return Tag{
								lang:   t.lang,
								script: scriptID(parents[i].script),
								region: regionID(parents[i].toRegion),
							}
						}
					}
				}
			}

			// Strip the script if it is the default one.
			base, _ := addTags(Tag{lang: t.lang})
			if base.script != maxScript {
				return Tag{lang: t.lang, script: maxScript}
			}
			return Tag{lang: t.lang}
		} else if t.script != 0 {
			// The parent for an base-script pair with a non-default script is
			// "und" instead of the base language.
			base, _ := addTags(Tag{lang: t.lang})
			if base.script != t.script {
				return und
			}
			return Tag{lang: t.lang}
		}
	}
	return und
=======
//
// Parent returns a tag for a less specific language that is mutually
// intelligible or Und if there is no such language. This may not be the same as
// simply stripping the last BCP 47 subtag. For instance, the parent of "zh-TW"
// is "zh-Hant", and the parent of "zh-Hant" is "und".
func (t Tag) Parent() Tag {
	return Tag(compact.Tag(t).Parent())
>>>>>>> v0.0.4
}

// returns token t and the rest of the string.
func nextToken(s string) (t, tail string) {
	p := strings.Index(s[1:], "-")
	if p == -1 {
		return s[1:], ""
	}
	p++
	return s[1:p], s[p:]
}

// Extension is a single BCP 47 extension.
type Extension struct {
	s string
}

// String returns the string representation of the extension, including the
// type tag.
func (e Extension) String() string {
	return e.s
}

// ParseExtension parses s as an extension and returns it on success.
func ParseExtension(s string) (e Extension, err error) {
<<<<<<< HEAD
	scan := makeScannerString(s)
	var end int
	if n := len(scan.token); n != 1 {
		return Extension{}, errSyntax
	}
	scan.toLower(0, len(scan.b))
	end = parseExtension(&scan)
	if end != len(s) {
		return Extension{}, errSyntax
	}
	return Extension{string(scan.b)}, nil
=======
	ext, err := language.ParseExtension(s)
	return Extension{ext}, err
>>>>>>> v0.0.4
}

// Type returns the one-byte extension type of e. It returns 0 for the zero
// exception.
func (e Extension) Type() byte {
	if e.s == "" {
		return 0
	}
	return e.s[0]
}

// Tokens returns the list of tokens of e.
func (e Extension) Tokens() []string {
	return strings.Split(e.s, "-")
}

// Extension returns the extension of type x for tag t. It will return
// false for ok if t does not have the requested extension. The returned
// extension will be invalid in this case.
func (t Tag) Extension(x byte) (ext Extension, ok bool) {
<<<<<<< HEAD
	for i := int(t.pExt); i < len(t.str)-1; {
		var ext string
		i, ext = getExtension(t.str, i)
		if ext[0] == x {
			return Extension{ext}, true
		}
	}
	return Extension{}, false
=======
	if !compact.Tag(t).MayHaveExtensions() {
		return Extension{}, false
	}
	e, ok := t.tag().Extension(x)
	return Extension{e}, ok
>>>>>>> v0.0.4
}

// Extensions returns all extensions of t.
func (t Tag) Extensions() []Extension {
<<<<<<< HEAD
	e := []Extension{}
	for i := int(t.pExt); i < len(t.str)-1; {
		var ext string
		i, ext = getExtension(t.str, i)
=======
	if !compact.Tag(t).MayHaveExtensions() {
		return nil
	}
	e := []Extension{}
	for _, ext := range t.tag().Extensions() {
>>>>>>> v0.0.4
		e = append(e, Extension{ext})
	}
	return e
}

// TypeForKey returns the type associated with the given key, where key and type
// are of the allowed values defined for the Unicode locale extension ('u') in
<<<<<<< HEAD
// http://www.unicode.org/reports/tr35/#Unicode_Language_and_Locale_Identifiers.
// TypeForKey will traverse the inheritance chain to get the correct value.
func (t Tag) TypeForKey(key string) string {
	if start, end, _ := t.findTypeForKey(key); end != start {
		return t.str[start:end]
	}
	return ""
}

var (
	errPrivateUse       = errors.New("cannot set a key on a private use tag")
	errInvalidArguments = errors.New("invalid key or type")
)

// SetTypeForKey returns a new Tag with the key set to type, where key and type
// are of the allowed values defined for the Unicode locale extension ('u') in
// http://www.unicode.org/reports/tr35/#Unicode_Language_and_Locale_Identifiers.
// An empty value removes an existing pair with the same key.
func (t Tag) SetTypeForKey(key, value string) (Tag, error) {
	if t.private() {
		return t, errPrivateUse
	}
	if len(key) != 2 {
		return t, errInvalidArguments
	}

	// Remove the setting if value is "".
	if value == "" {
		start, end, _ := t.findTypeForKey(key)
		if start != end {
			// Remove key tag and leading '-'.
			start -= 4

			// Remove a possible empty extension.
			if (end == len(t.str) || t.str[end+2] == '-') && t.str[start-2] == '-' {
				start -= 2
			}
			if start == int(t.pVariant) && end == len(t.str) {
				t.str = ""
				t.pVariant, t.pExt = 0, 0
			} else {
				t.str = fmt.Sprintf("%s%s", t.str[:start], t.str[end:])
			}
		}
		return t, nil
	}

	if len(value) < 3 || len(value) > 8 {
		return t, errInvalidArguments
	}

	var (
		buf    [maxCoreSize + maxSimpleUExtensionSize]byte
		uStart int // start of the -u extension.
	)

	// Generate the tag string if needed.
	if t.str == "" {
		uStart = t.genCoreBytes(buf[:])
		buf[uStart] = '-'
		uStart++
	}

	// Create new key-type pair and parse it to verify.
	b := buf[uStart:]
	copy(b, "u-")
	copy(b[2:], key)
	b[4] = '-'
	b = b[:5+copy(b[5:], value)]
	scan := makeScanner(b)
	if parseExtensions(&scan); scan.err != nil {
		return t, scan.err
	}

	// Assemble the replacement string.
	if t.str == "" {
		t.pVariant, t.pExt = byte(uStart-1), uint16(uStart-1)
		t.str = string(buf[:uStart+len(b)])
	} else {
		s := t.str
		start, end, hasExt := t.findTypeForKey(key)
		if start == end {
			if hasExt {
				b = b[2:]
			}
			t.str = fmt.Sprintf("%s-%s%s", s[:start], b, s[end:])
		} else {
			t.str = fmt.Sprintf("%s%s%s", s[:start], value, s[end:])
		}
	}
	return t, nil
}

// findKeyAndType returns the start and end position for the type corresponding
// to key or the point at which to insert the key-value pair if the type
// wasn't found. The hasExt return value reports whether an -u extension was present.
// Note: the extensions are typically very small and are likely to contain
// only one key-type pair.
func (t Tag) findTypeForKey(key string) (start, end int, hasExt bool) {
	p := int(t.pExt)
	if len(key) != 2 || p == len(t.str) || p == 0 {
		return p, p, false
	}
	s := t.str

	// Find the correct extension.
	for p++; s[p] != 'u'; p++ {
		if s[p] > 'u' {
			p--
			return p, p, false
		}
		if p = nextExtension(s, p); p == len(s) {
			return len(s), len(s), false
		}
	}
	// Proceed to the hyphen following the extension name.
	p++

	// curKey is the key currently being processed.
	curKey := ""

	// Iterate over keys until we get the end of a section.
	for {
		// p points to the hyphen preceding the current token.
		if p3 := p + 3; s[p3] == '-' {
			// Found a key.
			// Check whether we just processed the key that was requested.
			if curKey == key {
				return start, p, true
			}
			// Set to the next key and continue scanning type tokens.
			curKey = s[p+1 : p3]
			if curKey > key {
				return p, p, true
			}
			// Start of the type token sequence.
			start = p + 4
			// A type is at least 3 characters long.
			p += 7 // 4 + 3
		} else {
			// Attribute or type, which is at least 3 characters long.
			p += 4
		}
		// p points past the third character of a type or attribute.
		max := p + 5 // maximum length of token plus hyphen.
		if len(s) < max {
			max = len(s)
		}
		for ; p < max && s[p] != '-'; p++ {
		}
		// Bail if we have exhausted all tokens or if the next token starts
		// a new extension.
		if p == len(s) || s[p+2] == '-' {
			if curKey == key {
				return start, p, true
			}
			return p, p, true
		}
	}
}

// CompactIndex returns an index, where 0 <= index < NumCompactTags, for tags
// for which data exists in the text repository. The index will change over time
// and should not be stored in persistent storage. Extensions, except for the
// 'va' type of the 'u' extension, are ignored. It will return 0, false if no
// compact tag exists, where 0 is the index for the root language (Und).
func CompactIndex(t Tag) (index int, ok bool) {
	// TODO: perhaps give more frequent tags a lower index.
	// TODO: we could make the indexes stable. This will excluded some
	//       possibilities for optimization, so don't do this quite yet.
	b, s, r := t.Raw()
	if len(t.str) > 0 {
		if strings.HasPrefix(t.str, "x-") {
			// We have no entries for user-defined tags.
			return 0, false
		}
		if uint16(t.pVariant) != t.pExt {
			// There are no tags with variants and an u-va type.
			if t.TypeForKey("va") != "" {
				return 0, false
			}
			t, _ = Raw.Compose(b, s, r, t.Variants())
		} else if _, ok := t.Extension('u'); ok {
			// Strip all but the 'va' entry.
			variant := t.TypeForKey("va")
			t, _ = Raw.Compose(b, s, r)
			t, _ = t.SetTypeForKey("va", variant)
		}
		if len(t.str) > 0 {
			// We have some variants.
			for i, s := range specialTags {
				if s == t {
					return i + 1, true
				}
			}
			return 0, false
		}
	}
	// No variants specified: just compare core components.
	// The key has the form lllssrrr, where l, s, and r are nibbles for
	// respectively the langID, scriptID, and regionID.
	key := uint32(b.langID) << (8 + 12)
	key |= uint32(s.scriptID) << 12
	key |= uint32(r.regionID)
	x, ok := coreTags[key]
	return int(x), ok
}

// Base is an ISO 639 language code, used for encoding the base language
// of a language tag.
type Base struct {
	langID
=======
// https://www.unicode.org/reports/tr35/#Unicode_Language_and_Locale_Identifiers.
// TypeForKey will traverse the inheritance chain to get the correct value.
func (t Tag) TypeForKey(key string) string {
	if !compact.Tag(t).MayHaveExtensions() {
		if key != "rg" && key != "va" {
			return ""
		}
	}
	return t.tag().TypeForKey(key)
}

// SetTypeForKey returns a new Tag with the key set to type, where key and type
// are of the allowed values defined for the Unicode locale extension ('u') in
// https://www.unicode.org/reports/tr35/#Unicode_Language_and_Locale_Identifiers.
// An empty value removes an existing pair with the same key.
func (t Tag) SetTypeForKey(key, value string) (Tag, error) {
	tt, err := t.tag().SetTypeForKey(key, value)
	return makeTag(tt), err
}

// NumCompactTags is the number of compact tags. The maximum tag is
// NumCompactTags-1.
const NumCompactTags = compact.NumCompactTags

// CompactIndex returns an index, where 0 <= index < NumCompactTags, for tags
// for which data exists in the text repository.The index will change over time
// and should not be stored in persistent storage. If t does not match a compact
// index, exact will be false and the compact index will be returned for the
// first match after repeatedly taking the Parent of t.
func CompactIndex(t Tag) (index int, exact bool) {
	id, exact := compact.LanguageID(compact.Tag(t))
	return int(id), exact
}

var root = language.Tag{}

// Base is an ISO 639 language code, used for encoding the base language
// of a language tag.
type Base struct {
	langID language.Language
>>>>>>> v0.0.4
}

// ParseBase parses a 2- or 3-letter ISO 639 code.
// It returns a ValueError if s is a well-formed but unknown language identifier
// or another error if another error occurred.
func ParseBase(s string) (Base, error) {
<<<<<<< HEAD
	if n := len(s); n < 2 || 3 < n {
		return Base{}, errSyntax
	}
	var buf [3]byte
	l, err := getLangID(buf[:copy(buf[:], s)])
	return Base{l}, err
}

// Script is a 4-letter ISO 15924 code for representing scripts.
// It is idiomatically represented in title case.
type Script struct {
	scriptID
=======
	l, err := language.ParseBase(s)
	return Base{l}, err
}

// String returns the BCP 47 representation of the base language.
func (b Base) String() string {
	return b.langID.String()
}

// ISO3 returns the ISO 639-3 language code.
func (b Base) ISO3() string {
	return b.langID.ISO3()
}

// IsPrivateUse reports whether this language code is reserved for private use.
func (b Base) IsPrivateUse() bool {
	return b.langID.IsPrivateUse()
}

// Script is a 4-letter ISO 15924 code for representing scripts.
// It is idiomatically represented in title case.
type Script struct {
	scriptID language.Script
>>>>>>> v0.0.4
}

// ParseScript parses a 4-letter ISO 15924 code.
// It returns a ValueError if s is a well-formed but unknown script identifier
// or another error if another error occurred.
func ParseScript(s string) (Script, error) {
<<<<<<< HEAD
	if len(s) != 4 {
		return Script{}, errSyntax
	}
	var buf [4]byte
	sc, err := getScriptID(script, buf[:copy(buf[:], s)])
	return Script{sc}, err
}

// Region is an ISO 3166-1 or UN M.49 code for representing countries and regions.
type Region struct {
	regionID
=======
	sc, err := language.ParseScript(s)
	return Script{sc}, err
}

// String returns the script code in title case.
// It returns "Zzzz" for an unspecified script.
func (s Script) String() string {
	return s.scriptID.String()
}

// IsPrivateUse reports whether this script code is reserved for private use.
func (s Script) IsPrivateUse() bool {
	return s.scriptID.IsPrivateUse()
}

// Region is an ISO 3166-1 or UN M.49 code for representing countries and regions.
type Region struct {
	regionID language.Region
>>>>>>> v0.0.4
}

// EncodeM49 returns the Region for the given UN M.49 code.
// It returns an error if r is not a valid code.
func EncodeM49(r int) (Region, error) {
<<<<<<< HEAD
	rid, err := getRegionM49(r)
=======
	rid, err := language.EncodeM49(r)
>>>>>>> v0.0.4
	return Region{rid}, err
}

// ParseRegion parses a 2- or 3-letter ISO 3166-1 or a UN M.49 code.
// It returns a ValueError if s is a well-formed but unknown region identifier
// or another error if another error occurred.
func ParseRegion(s string) (Region, error) {
<<<<<<< HEAD
	if n := len(s); n < 2 || 3 < n {
		return Region{}, errSyntax
	}
	var buf [3]byte
	r, err := getRegionID(buf[:copy(buf[:], s)])
	return Region{r}, err
}

// IsCountry returns whether this region is a country or autonomous area. This
// includes non-standard definitions from CLDR.
func (r Region) IsCountry() bool {
	if r.regionID == 0 || r.IsGroup() || r.IsPrivateUse() && r.regionID != _XK {
		return false
	}
	return true
=======
	r, err := language.ParseRegion(s)
	return Region{r}, err
}

// String returns the BCP 47 representation for the region.
// It returns "ZZ" for an unspecified region.
func (r Region) String() string {
	return r.regionID.String()
}

// ISO3 returns the 3-letter ISO code of r.
// Note that not all regions have a 3-letter ISO code.
// In such cases this method returns "ZZZ".
func (r Region) ISO3() string {
	return r.regionID.ISO3()
}

// M49 returns the UN M.49 encoding of r, or 0 if this encoding
// is not defined for r.
func (r Region) M49() int {
	return r.regionID.M49()
}

// IsPrivateUse reports whether r has the ISO 3166 User-assigned status. This
// may include private-use tags that are assigned by CLDR and used in this
// implementation. So IsPrivateUse and IsCountry can be simultaneously true.
func (r Region) IsPrivateUse() bool {
	return r.regionID.IsPrivateUse()
}

// IsCountry returns whether this region is a country or autonomous area. This
// includes non-standard definitions from CLDR.
func (r Region) IsCountry() bool {
	return r.regionID.IsCountry()
>>>>>>> v0.0.4
}

// IsGroup returns whether this region defines a collection of regions. This
// includes non-standard definitions from CLDR.
func (r Region) IsGroup() bool {
<<<<<<< HEAD
	if r.regionID == 0 {
		return false
	}
	return int(regionInclusion[r.regionID]) < len(regionContainment)
=======
	return r.regionID.IsGroup()
>>>>>>> v0.0.4
}

// Contains returns whether Region c is contained by Region r. It returns true
// if c == r.
func (r Region) Contains(c Region) bool {
<<<<<<< HEAD
	return r.regionID.contains(c.regionID)
}

func (r regionID) contains(c regionID) bool {
	if r == c {
		return true
	}
	g := regionInclusion[r]
	if g >= nRegionGroups {
		return false
	}
	m := regionContainment[g]

	d := regionInclusion[c]
	b := regionInclusionBits[d]

	// A contained country may belong to multiple disjoint groups. Matching any
	// of these indicates containment. If the contained region is a group, it
	// must strictly be a subset.
	if d >= nRegionGroups {
		return b&m != 0
	}
	return b&^m == 0
}

var errNoTLD = errors.New("language: region is not a valid ccTLD")

=======
	return r.regionID.Contains(c.regionID)
}

>>>>>>> v0.0.4
// TLD returns the country code top-level domain (ccTLD). UK is returned for GB.
// In all other cases it returns either the region itself or an error.
//
// This method may return an error for a region for which there exists a
// canonical form with a ccTLD. To get that ccTLD canonicalize r first. The
// region will already be canonicalized it was obtained from a Tag that was
// obtained using any of the default methods.
func (r Region) TLD() (Region, error) {
<<<<<<< HEAD
	// See http://en.wikipedia.org/wiki/Country_code_top-level_domain for the
	// difference between ISO 3166-1 and IANA ccTLD.
	if r.regionID == _GB {
		r = Region{_UK}
	}
	if (r.typ() & ccTLD) == 0 {
		return Region{}, errNoTLD
	}
	return r, nil
=======
	tld, err := r.regionID.TLD()
	return Region{tld}, err
>>>>>>> v0.0.4
}

// Canonicalize returns the region or a possible replacement if the region is
// deprecated. It will not return a replacement for deprecated regions that
// are split into multiple regions.
func (r Region) Canonicalize() Region {
<<<<<<< HEAD
	if cr := normRegion(r.regionID); cr != 0 {
		return Region{cr}
	}
	return r
=======
	return Region{r.regionID.Canonicalize()}
>>>>>>> v0.0.4
}

// Variant represents a registered variant of a language as defined by BCP 47.
type Variant struct {
	variant string
}

// ParseVariant parses and returns a Variant. An error is returned if s is not
// a valid variant.
func ParseVariant(s string) (Variant, error) {
<<<<<<< HEAD
	s = strings.ToLower(s)
	if _, ok := variantIndex[s]; ok {
		return Variant{s}, nil
	}
	return Variant{}, mkErrInvalid([]byte(s))
=======
	v, err := language.ParseVariant(s)
	return Variant{v.String()}, err
>>>>>>> v0.0.4
}

// String returns the string representation of the variant.
func (v Variant) String() string {
	return v.variant
}
