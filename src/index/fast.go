package index

import (
	"bytes"
	"sort"

	"github.com/mathieunls/gainful/src/binary"
	"github.com/mathieunls/gainful/src/indexable"
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
	fs.bst = binary.NewParralelTree(keys, values, -1)
}

func (x *Index) lookupAllPara(s []byte) []int {
	// find matching suffix index range [i:j]
	// find the first index where s would be the prefix
	i := sort.Search(len(x.sa), func(i int) bool { return bytes.Compare(x.at(i), s) >= 0 })
	// starting at i, find the first index at which s is not a prefix
	j := i + sort.Search(len(x.sa)-i, func(j int) bool { return !bytes.HasPrefix(x.at(j+i), s) })
	return x.sa[i:j]
}

func (fs *FastIndex) Lookup(search string, n int) []indexable.HasStringIndex {

	// https: //stackoverflow.com/questions/8423873/parallel-binary-search
	// https://golang.org/src/sort/search.go

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
