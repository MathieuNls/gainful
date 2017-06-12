package gainful

import (
	"index/suffixarray"
	"runtime"
	"strings"
)

type Index struct {
	sa  *suffixarray.Index
	bst *bst
}

type Indexable interface {
	StringIndex() string
}

func NewIndex(values []Indexable) *Index {

	keys := make([]int, len(values))
	stringValues := make([]string, len(values))

	fs := &Index{}

	currLength := 1

	for i := 0; i < len(values); i++ {
		keys[i] = currLength
		str := values[i].StringIndex()
		currLength += len(str) + 1
		stringValues[i] = str
	}

	fs.bst = FromKeys(keys, values, true)

	joinedStrings := "\x00" + strings.Join(stringValues, "\x00")

	fs.sa = suffixarray.New([]byte(joinedStrings))

	return fs
}

func (fs *Index) Find(search string, n int) []Indexable {

	offsets := fs.sa.Lookup([]byte(search), n)
	results := []Indexable{}
	knownKeys := make(map[int]struct{})

	keys := make(chan int, len(offsets))
	resultsChan := make(chan *binaryNode, len(offsets))

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
			if _, present := knownKeys[node.key]; !present {
				knownKeys[node.key] = struct{}{}
				results = append(results, node.value)
			}
		}
	}

	return results
}

func (fs *Index) FindSequential(search string, n int) []Indexable {

	offsets := fs.sa.Lookup([]byte(search), n)
	results := []Indexable{}
	knownKeys := make(map[int]struct{})

	for _, off := range offsets {

		node, err := fs.bst.FloorKey(off)

		if _, present := knownKeys[node.key]; !present && err == nil {
			knownKeys[node.key] = struct{}{}
			results = append(results, node.value)
		}
	}

	return results
}

func (fs *Index) bstLookupWorker(keys <-chan int, results chan *binaryNode) {

	for key := range keys {
		node, _ := fs.bst.FloorKey(key)

		results <- node
	}
}
