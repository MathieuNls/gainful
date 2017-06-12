package index

import (
	"index/suffixarray"
	"runtime"
	"strings"

	"github.com/mathieunls/gainful/src/binary"
	"github.com/mathieunls/gainful/src/indexable"
)

type Index struct {
	sa  *suffixarray.Index
	bst *binary.Tree
}

func New(values []indexable.HasStringIndex) *Index {

	keys := make([]int, len(values))
	stringValues := make([]string, len(values))

	fs := &Index{}

	currLength := 0

	for i := 0; i < len(values); i++ {
		keys[i] = currLength
		str := values[i].StringIndex()
		currLength += len(str) + 0
		stringValues[i] = str
	}

	fs.bst = binary.FromKeys(keys, values, true)

	joinedStrings := "" + strings.Join(stringValues, "")

	fs.sa = suffixarray.New([]byte(joinedStrings))

	return fs
}

func (fs *Index) Find(search string, n int) []indexable.HasStringIndex {

	offsets := fs.sa.Lookup([]byte(search), n)
	results := []indexable.HasStringIndex{}
	knownKeys := make(map[int]struct{})

	keys := make(chan int, len(offsets))
	resultsChan := make(chan *binary.Node, len(offsets))

	for i := 0; i < runtime.NumCPU(); i++ {
		go fs.bstLookupWorker(keys, resultsChan)
	}

	for _, off := range offsets {

		keys <- off
	}
	close(keys)

	for _ = range offsets {

		node := <-resultsChan

		if node != nil {
			if _, present := knownKeys[node.Key]; !present {
				knownKeys[node.Key] = struct{}{}
				results = append(results, node.Value)
			}
		}
	}

	return results
}

func (fs *Index) FindSequential(search string, n int) []indexable.HasStringIndex {

	offsets := fs.sa.Lookup([]byte(search), n)
	results := []indexable.HasStringIndex{}
	knownKeys := make(map[int]struct{})

	for _, off := range offsets {

		node, err := fs.bst.FloorKey(off)

		if _, present := knownKeys[node.Key]; !present && err == nil {
			knownKeys[node.Key] = struct{}{}
			results = append(results, node.Value)
		}
	}

	return results
}

func (fs *Index) bstLookupWorker(keys <-chan int, results chan *binary.Node) {

	for key := range keys {
		node, _ := fs.bst.FloorKey(key)

		results <- node
	}
}
