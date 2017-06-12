package index

import (
	"index/suffixarray"

	"github.com/mathieunls/gainful/src/binary"
)

type FastIndex struct {
	suffixarray.Index
	bst *binary.Tree
}

func (f *FastIndex) oneFunc() {

}

func (f *FastIndex) lookupAll(s []byte) []int {
	return []int{1}
}

func (f *FastIndex) Lookup(s []byte, n int) (result []int) {
	if len(s) > 0 && n != 0 {
		matches := f.lookupAll(s)
		if n < 0 || len(matches) < n {
			n = len(matches)
		}
		// 0 <= n <= len(matches)
		if n > 0 {
			result = make([]int, n)
			copy(result, matches)
		}
	}
	return
}
