package index

import (
	"testing"

	"strings"

	lorem "github.com/drhodes/golorem"
	"github.com/mathieunls/gainful/src/indexable"
)

func TestNew(t *testing.T) {

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

	indexables := make([]indexable.HasStringIndex, len(values))

	for i := 0; i < len(values); i++ {
		indexables[i] = indexable.New(values[i])
	}

	i := NewSearchIndex(indexables)

	//a word
	_, _, _, results := i.Lookup("sky", 0, -1, -1, nil)

	if len(results) != 1 || results[0].StringIndex() != values[0] {
		t.Error("expected", values[0], "got", results)
	}

	//word in middle with many matches
	_, _, _, results = i.Lookup("where", 0, -1, -1, nil)

	if len(results) != 2 || results[0].StringIndex() != values[4] || results[1].StringIndex() != values[7] {
		t.Error("expected", values[4], "got", results)
	}

	//reduce the matches
	_, _, _, results = i.Lookup("where", 0, -1, 1, nil)

	if len(results) != 1 || results[0].StringIndex() != values[4] {
		t.Error("expected", values[4], "got", results)
	}

	//end words
	_, _, _, results = i.Lookup("lie.", 0, -1, -1, nil)
	if len(results) != 2 || results[0].StringIndex() != values[4] || results[1].StringIndex() != values[7] {
		t.Error("expected", values[4], "got", results)
	}

	//start words
	_, _, _, results = i.Lookup("In", 0, -1, -1, nil)

	if len(results) != 2 || results[0].StringIndex() != values[4] || results[1].StringIndex() != values[7] {
		t.Error("expected", values[4], "got", results)
	}

	//One sentence with multiple match - erased
	_, _, _, results = i.Lookup("One", 0, -1, -1, nil)

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

	indexables := make([]indexable.HasStringIndex, len(values))

	for i := 0; i < len(values); i++ {
		indexables[i] = indexable.New(values[i])
	}

	i := NewSearchIndex(indexables)

	//a word
	results := i.FindSequential("sky", 10, -1)

	if len(results) != 1 || results[0].StringIndex() != values[0] {
		t.Error("expected", values[0], "got", results)
	}

	//word in middle with many matches
	results = i.FindSequential("where", 0, -1)

	if len(results) != 2 || results[0].StringIndex() != values[4] || results[1].StringIndex() != values[7] {
		t.Error("expected", values[4], "got", results)
	}

	//reduce the matches
	results = i.FindSequential("where", 0, 1)

	if len(results) != 1 || results[0].StringIndex() != values[4] {
		t.Error("expected", values[4], "got", results)
	}

	//end words
	results = i.FindSequential("lie.", 0, -1)
	if len(results) != 2 || results[0].StringIndex() != values[4] || results[1].StringIndex() != values[7] {
		t.Error("expected", values[4], "got", results)
	}

	//start words
	results = i.FindSequential("In", 0, -1)

	if len(results) != 2 || results[0].StringIndex() != values[4] || results[1].StringIndex() != values[7] {
		t.Error("expected", values[4], "got", results)
	}

	//One sentence with multiple match - erased
	results = i.FindSequential("One", 0, -1)

	if len(results) != 3 {
		t.Error("expected 3 got", results)
	}

}

/*
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
				f := NewSearchIndex(dataset(size))
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
				f := NewSearchIndex(dataset(size))
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
*/
func dumb(data []interface{}) {

	indexables := data[0].([]indexable.HasStringIndex)
	word := data[1].(string)
	results := []indexable.HasStringIndex{}

	for i := 0; i < len(indexables); i++ {
		if strings.Contains(indexables[i].StringIndex(), word) {
			results = append(results, indexables[i])
		}
	}
}

func sequential(data []interface{}) {

	f := data[0].(*SearchIndex)
	word := data[1].(string)
	f.FindSequential(word, 0, -1)
	f = nil
}

func parralel(data []interface{}) {
	f := data[0].(*SearchIndex)
	word := data[1].(string)
	f.Lookup(word, 0, -1, -1, nil)
	f = nil
}

func dataset(size int) []indexable.HasStringIndex {

	indexables := make([]indexable.HasStringIndex, size)

	for i := 0; i < size; i++ {
		indexables[i] = indexable.New(lorem.Sentence(10, 100))
	}
	return indexables
}
