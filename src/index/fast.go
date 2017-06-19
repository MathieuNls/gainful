package index

import (
	"github.com/mathieunls/gainful/src/indexable"
	gpbt "github.com/mathieunls/gpbt/src"
)

type FastIndex struct {
	SearchIndex
}

func NewFastIndex(values []indexable.HasStringIndex) *FastIndex {

	fs := &FastIndex{}
	fs.init(values)
	return fs
}

func (fs *FastIndex) newTree(keys []int, values []indexable.HasStringIndex, sorted bool) {
	var interfaceSlice = make([]interface{}, len(values))
	for i, d := range values {
		interfaceSlice[i] = d
	}
	fs.bst = gpbt.NewParralelTree(keys, interfaceSlice, -1)
}

func (fs *FastIndex) Lookup(search string, n int) []indexable.HasStringIndex {

	results := []indexable.HasStringIndex{}
	s := []byte(search)

	if len(s) > 0 && n != 0 {
		matches := fs.SearchIndex.sa.lookupAll(s)
		if n < 0 || len(matches) < n {
			n = len(matches)
		}
		// 0 <= n <= len(matches)
		if n > 0 {

			return fs.SearchIndex.findPara(matches)
		}
	}
	return results
}
