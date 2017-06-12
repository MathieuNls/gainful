package gainful

import (
	"testing"

	"strings"

	"fmt"

	lorem "github.com/drhodes/golorem"
	"github.com/mathieunls/gripper"
)

func TestNewIndex(t *testing.T) {

	values := []string{
		"Three Rings for the Elven-kings under the sky,",
		"Seven for the Dwarf-lords in their halls of stone,",
		"Nine for Mortal Men doomed to die,",
		"One for the Dark Lord on his dark throne",
		"In the Land of Mordor where the Shadows lie.",
		"One Ring to rule them all, One Ring to find them,",
		"One Ring to bring them all and in the darkness bind them",
		"In the Land of Mordor where the Shadows lie.",
	}

	indexables := make([]Indexable, len(values))

	for i := 0; i < len(values); i++ {
		indexables[i] = newIndexable(values[i])
	}

	index := NewIndex(indexables)

	//a word
	results := index.Find("sky", -1)

	if len(results) != 1 || results[0].StringIndex() != values[0] {
		t.Error("expected", values[0], "got", results)
	}

	//word in middle with many matches
	results = index.Find("where", -1)

	if len(results) != 2 || results[0].StringIndex() != values[4] || results[1].StringIndex() != values[7] {
		t.Error("expected", values[4], "got", results)
	}

	//reduce the matches
	results = index.Find("where", 1)

	if len(results) != 1 || results[0].StringIndex() != values[4] {
		t.Error("expected", values[4], "got", results)
	}

	//end words
	results = index.Find("lie.", -1)
	if len(results) != 2 || results[0].StringIndex() != values[4] || results[1].StringIndex() != values[7] {
		t.Error("expected", values[4], "got", results)
	}

	//start words
	results = index.Find("In", -1)

	if len(results) != 2 || results[0].StringIndex() != values[4] || results[1].StringIndex() != values[7] {
		t.Error("expected", values[4], "got", results)
	}

	//One sentence with multiple match - erased
	results = index.Find("One", -1)

	if len(results) != 3 {
		t.Error("expected 3 got", results)
	}

}

func TestNewSequential(t *testing.T) {

	values := []string{
		"Three Rings for the Elven-kings under the sky,",
		"Seven for the Dwarf-lords in their halls of stone,",
		"Nine for Mortal Men doomed to die,",
		"One for the Dark Lord on his dark throne",
		"In the Land of Mordor where the Shadows lie.",
		"One Ring to rule them all, One Ring to find them,",
		"One Ring to bring them all and in the darkness bind them",
		"In the Land of Mordor where the Shadows lie.",
	}

	indexables := make([]Indexable, len(values))

	for i := 0; i < len(values); i++ {
		indexables[i] = newIndexable(values[i])
	}

	index := NewIndex(indexables)

	//a word
	results := index.FindSequential("sky", -1)

	if len(results) != 1 || results[0].StringIndex() != values[0] {
		t.Error("expected", values[0], "got", results)
	}

	//word in middle with many matches
	results = index.FindSequential("where", -1)

	if len(results) != 2 || results[0].StringIndex() != values[4] || results[1].StringIndex() != values[7] {
		t.Error("expected", values[4], "got", results)
	}

	//reduce the matches
	results = index.FindSequential("where", 1)

	if len(results) != 1 || results[0].StringIndex() != values[4] {
		t.Error("expected", values[4], "got", results)
	}

	//end words
	results = index.FindSequential("lie.", -1)
	if len(results) != 2 || results[0].StringIndex() != values[4] || results[1].StringIndex() != values[7] {
		t.Error("expected", values[4], "got", results)
	}

	//start words
	results = index.FindSequential("In", -1)

	if len(results) != 2 || results[0].StringIndex() != values[4] || results[1].StringIndex() != values[7] {
		t.Error("expected", values[4], "got", results)
	}

	//One sentence with multiple match - erased
	results = index.FindSequential("One", -1)

	if len(results) != 3 {
		t.Error("expected 3 got", results)
	}

}

func TestPerformances(t *testing.T) {

	max := 10
	increment := 1
	retries := 10

	gripper.PerfPlotter().
		AnalyzeWithGeneratedData(
			func(size int) []interface{} {
				r := make([]interface{}, 2)
				r[0] = dataset(size)
				r[1] = lorem.Word(5, 15)
				fmt.Println("running dumb with", size)
				return r
			},
			dumb,
			max,
			increment,
			retries,
			"classical",
		).
		AnalyzeWithGeneratedData(
			func(size int) []interface{} {
				r := make([]interface{}, 2)
				f := NewIndex(dataset(size))
				r[0] = f
				r[1] = lorem.Word(5, 15)
				fmt.Println("running sequential with", size)
				return r
			},
			sequential,
			max,
			increment,
			retries,
			"suffix bst",
		).
		AnalyzeWithGeneratedData(
			func(size int) []interface{} {
				r := make([]interface{}, 2)
				f := NewIndex(dataset(size))
				r[0] = f
				r[1] = lorem.Word(5, 15)
				fmt.Println("running parralel with", size)
				return r
			},
			parralel,
			max,
			increment,
			retries,
			"suffix bst w/ k=4",
		).
		Plot("data size", "ms", "Time complexity", "testing2.png")

}

func dumb(data []interface{}) {

	indexables := data[0].([]Indexable)
	word := data[1].(string)
	results := []Indexable{}

	for i := 0; i < len(indexables); i++ {
		if strings.Contains(indexables[i].StringIndex(), word) {
			results = append(results, indexables[i])
		}
	}
}

func sequential(data []interface{}) {

	f := data[0].(*Index)
	word := data[1].(string)
	f.FindSequential(word, -1)
	f = nil
}

func parralel(data []interface{}) {
	f := data[0].(*Index)
	word := data[1].(string)
	f.Find(word, -1)
	f = nil
}

func dataset(size int) []Indexable {

	indexables := make([]Indexable, size)

	for i := 0; i < size; i++ {
		indexables[i] = newIndexable(lorem.Sentence(10, 100))
	}
	return indexables
}
