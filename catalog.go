// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"de": &dictionary{index: deIndex, data: deData},
		"en": &dictionary{index: enIndex, data: enData},
	}
	fallback := language.MustParse("en")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"%.2f":                  18,
	"Amount":                13,
	"Daily tracked":         15,
	"Daily tracked time:":   11,
	"Daily un-tracked":      14,
	"Daily untracked time:": 12,
	"Date":                  0,
	"Duration":              3,
	"End":                   2,
	"Exact amount":          16,
	"Exact duration":        17,
	"Exact tracked time:":   10,
	"Notes":                 4,
	"Project":               7,
	"Rounded duration":      6,
	"Start":                 1,
	"Time range:":           8,
	"Total":                 5,
	"Tracked time:":         9,
}

var deIndex = []uint32{ // 20 elements
	0x00000000, 0x00000006, 0x0000000d, 0x00000012,
	0x00000018, 0x00000024, 0x0000002b, 0x0000003a,
	0x00000042, 0x0000004f, 0x0000005e, 0x00000073,
	0x00000089, 0x000000a2, 0x000000a9, 0x000000bc,
	0x000000cd, 0x000000dc, 0x000000e9, 0x000000f1,
} // Size: 104 bytes

const deData string = "" + // Size: 241 bytes
	"\x02Datum\x02Beginn\x02Ende\x02Dauer\x02Anmerkungen\x02Gesamt\x02Gerunde" +
	"te Zeit\x02Projekt\x02Zeitbereich:\x02Erfasste Zeit:\x02Exakt erfasste Z" +
	"eit:\x02Täglich erfasst Zeit\x02Täglich unerfasste Zeit\x02Betrag\x02Täg" +
	"lich unerfasst\x02Täglich erfasst\x02Exakter Betrag\x02Exakte Dauer\x02%" +
	".2[1]f"

var enIndex = []uint32{ // 20 elements
	0x00000000, 0x00000005, 0x0000000b, 0x0000000f,
	0x00000018, 0x0000001e, 0x00000024, 0x00000035,
	0x0000003d, 0x00000049, 0x00000057, 0x0000006b,
	0x0000007f, 0x00000095, 0x0000009c, 0x000000ad,
	0x000000bb, 0x000000c8, 0x000000d7, 0x000000df,
} // Size: 104 bytes

const enData string = "" + // Size: 223 bytes
	"\x02Date\x02Start\x02End\x02Duration\x02Notes\x02Total\x02Rounded durati" +
	"on\x02Project\x02Time range:\x02Tracked time:\x02Exact tracked time:\x02" +
	"Daily tracked time:\x02Daily untracked time:\x02Amount\x02Daily un-track" +
	"ed\x02Daily tracked\x02Exact amount\x02Exact duration\x02%.2[1]f"

	// Total table size 672 bytes (0KiB); checksum: 9B1B0EBC
