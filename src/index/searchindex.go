package index

import (
	"runtime"
	"strings"

	"github.com/mathieunls/gainful/src/indexable"
	gpbt "github.com/mathieunls/gpbt/src"
)

type SearchIndex struct {
	sa  *Index
	bst gpbt.NavigableTree
}

func NewSearchIndex(values []indexable.HasStringIndex) *SearchIndex {

	fs := &SearchIndex{}
	fs.init(values)

	return fs
}

func (fs *SearchIndex) init(values []indexable.HasStringIndex) {
	keys := make([]int, len(values))
	stringValues := make([]string, len(values))
	currLength := 0

	for i := 0; i < len(values); i++ {
		keys[i] = currLength
		str := values[i].StringIndex()
		currLength += len(str) + 0
		stringValues[i] = str
	}

	fs.newTree(keys, values, true)

	joinedStrings := "" + strings.Join(stringValues, "")

	fs.sa = New([]byte(joinedStrings))
}

//newTree purpose is to be overrided by other Indexes
//with other trees implementation
func (fs *SearchIndex) newTree(keys []int, values []indexable.HasStringIndex, sorted bool) {

	var interfaceSlice = make([]interface{}, len(values))
	for i, d := range values {
		interfaceSlice[i] = d
	}
	fs.bst = gpbt.NewTree(keys, interfaceSlice, sorted)
}

func (fs *SearchIndex) Lookup(
	search string,
	n int,
	sort func([]indexable.HasStringIndex) []indexable.HasStringIndex) []indexable.HasStringIndex {

	results := []indexable.HasStringIndex{}
	s := []byte(search)

	if len(s) > 0 && n != 0 {
		matches := fs.sa.lookupAll(s)
		results = fs.findPara(matches)

		if n == -1 {
			n = len(results)
		}

		if sort == nil {
			return results[0:n]
		}
		return sort(results)[0:n]

	}
	return results
}

func (fs *SearchIndex) findPara(offsets []int) []indexable.HasStringIndex {
	results := []indexable.HasStringIndex{}
	knownKeys := make(map[int]struct{})

	keys := make(chan int, len(offsets))
	resultsChan := make(chan *gpbt.Node, len(offsets))

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
				results = append(results, node.Value.(indexable.HasStringIndex))
			}
		}
	}

	return results
}

func (fs *SearchIndex) FindSequential(search string, n int) []indexable.HasStringIndex {

	offsets := fs.sa.Lookup([]byte(search), n)
	results := []indexable.HasStringIndex{}
	knownKeys := make(map[int]struct{})

	for _, off := range offsets {

		node, err := fs.bst.FloorKey(off)

		if _, present := knownKeys[node.Key]; !present && err == nil {
			knownKeys[node.Key] = struct{}{}
			results = append(results, node.Value.(indexable.HasStringIndex))
		}
	}

	return results
}

func (fs *SearchIndex) bstLookupWorker(keys <-chan int, results chan *gpbt.Node) {

	for key := range keys {
		node, _ := fs.bst.FloorKey(key)

		results <- node
	}
}
