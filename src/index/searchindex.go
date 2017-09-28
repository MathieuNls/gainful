package index

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/mathieunls/gainful/src/indexable"
	"github.com/mathieunls/gpbt/src"
)

type SearchIndex struct {
	sa      *Index
	bst     gpbt.NavigableTree
	newTree func(keys []int, values []indexable.HasStringIndex, sorted bool)
}

func NewSearchIndex(values []indexable.HasStringIndex) *SearchIndex {

	fs := &SearchIndex{}
	fs.newTree = func(keys []int, values []indexable.HasStringIndex, sorted bool) {

		var interfaceSlice = make([]interface{}, len(values))
		for i, d := range values {
			interfaceSlice[i] = d
		}
		fs.bst = gpbt.NewTree(keys, interfaceSlice, sorted)
	}
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

//Lookup searches for the search string with respect to
//start and end which represent indexes in the suffix tree.
//They can be used to build pagination
//n limits to n the number of the matches in the suffix tree that will
//be investigated further. Each inxdexed object can have many matches.
//Consequently, n is not necessary equal to len(results).
//sort is a function that can be defined to sort the results
//according to user preferences.
//
//Lookup returns the elapsed time for the query, the number of item found
//The last index for the
func (fs *SearchIndex) Lookup(
	search string,
	start int,
	end int,
	n int,
	sort func([]indexable.HasStringIndex) []indexable.HasStringIndex) (
	elapsedTime int64,
	resultsCount int,
	lastIndex int,
	results []indexable.HasStringIndex,
) {

	startTime := time.Now()

	if len(search) > 0 && n != 0 {

		tmp := fs.sa.Lookup([]byte(search), n, 0)
		fmt.Println(tmp)
		matches := fs.findPara(tmp)

		return time.Since(startTime).Nanoseconds(), len(matches), 0, matches

	}

	return 0, 0, 0, nil
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

//FindSequential is a sequential version of Lookup and allows
//to check the sanity and efficiency of the || algorithm
func (fs *SearchIndex) FindSequential(search string, start int, n int) []indexable.HasStringIndex {

	offsets := fs.sa.Lookup([]byte(search), n, start)
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
