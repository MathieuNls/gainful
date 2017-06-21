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
	return fs
}

func (fs *FastIndex) newTree(keys []int, values []indexable.HasStringIndex, sorted bool) {
	var interfaceSlice = make([]interface{}, len(values))
	for i, d := range values {
		interfaceSlice[i] = d
	}
	fs.SearchIndex.bst = gpbt.NewParralelTree(keys, interfaceSlice, -1)
}
