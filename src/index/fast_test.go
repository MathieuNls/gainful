package index

import (
	"testing"

	"github.com/mathieunls/gainful/src/indexable"
)

func TestNewFastIndex(t *testing.T) {
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

	for index := 4; index < 22; index = index + 2 {
		fi := NewFastIndex(indexables, index)
		i := NewSearchIndex(indexables)
		_, _, _, results := i.Lookup("sky", 0, -1, -1, nil)

		if len(results) != 1 || results[0].StringIndex() != values[0] {
			t.Error("expected", values[0], "got", results, "with", index, "threads")
		}

		_, _, _, results = fi.Lookup("sky", 0, -1, -1, nil)

		if len(results) != 1 || results[0].StringIndex() != values[0] {
			t.Error("expected", values[0], "got", results[0], "with", index, "threads")
		}
	}
}
