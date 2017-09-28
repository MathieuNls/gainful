package index

import (
	"github.com/mathieunls/gainful/src/indexable"
	gpbt "github.com/mathieunls/gpbt/src"
)

type FastIndex struct {
	SearchIndex
}

func NewFastIndex(values []indexable.HasStringIndex, threads int) *FastIndex {

	fs := &FastIndex{}
	fs.SearchIndex.newTree = func(keys []int, values []indexable.HasStringIndex, sorted bool) {

		var interfaceSlice = make([]interface{}, len(values))
		for i, d := range values {
			interfaceSlice[i] = d
		}
		fs.SearchIndex.bst = gpbt.NewParralelTree(keys, interfaceSlice, threads)
		fs.SearchIndex.bst.Print()
	}
	fs.init(values)

	return fs
}
