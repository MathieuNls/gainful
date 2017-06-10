package gainful

import "testing"

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

	index.bst.Print()

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

	index.bst.Print()

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
