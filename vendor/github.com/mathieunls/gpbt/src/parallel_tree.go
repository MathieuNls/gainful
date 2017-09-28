package gpbt

import (
	"errors"
	"runtime"
	"sync"
)

//ParralelTree is a binary tree where each node is itself another binary tree
//All the keys can be accessed via Fetch and FloorKey in a transparent manner.
//From the outside worlds, this behaves exactly the same as binary.tree
type ParralelTree struct {
	Tree
	threads int
}

//NewParralelTree constructs a threads threads parrallel tree with keys and values.
// /!\ The keys have to be sorted /!\
//The time complexity of the construction is k + n/k where k is the number of threads
//and n the number of keys
func NewParralelTree(keys []int, values []interface{}, threads int) NavigableTree {

	tree := &ParralelTree{}

	if threads == -1 {
		tree.threads = runtime.NumCPU()
	} else if threads > len(values) {
		tree.threads = len(values)
	} else {
		tree.threads = threads
	}

	//Store the sub b-trees
	trees := make([]interface{}, tree.threads)
	//Store the root value of the sub b-trees
	splittedTreesRootValues := make([]int, tree.threads)

	keysPerThread := len(keys) / tree.threads

	//If the len(keys) / tree.threads isn't a round division
	//we'll need add the rest of element
	//This stores how many *rest* element are left to add
	keysOutsideThreadBoundaries := len(keys) % tree.threads

	if keysPerThread > 0 {

		var wg sync.WaitGroup
		wg.Add(tree.threads)

		start := 0
		currentThread := 0

		//while they are keys to add
		for start < len(keys) {

			end := start + keysPerThread - 1

			//if there are still key outside the rounded division
			if keysOutsideThreadBoundaries > 0 {
				end++
				keysOutsideThreadBoundaries--
			}

			go func(wg *sync.WaitGroup, start int, end int, thread int) {

				localRoot := tree.fromSortedKeys(keys, values, start, end, nil)
				trees[thread] = &Tree{
					root: localRoot,
				}
				splittedTreesRootValues[thread] = localRoot.Key
				wg.Done()
			}(&wg, start, end, currentThread)

			start = end + 1
			currentThread++
		}

		wg.Wait()

		tree.root = tree.fromSortedKeys(splittedTreesRootValues, trees, 0, tree.threads-1, nil)
	}

	return tree
}

//Fetch fetches a key
//Complexity is O(log(k) + log(n)/k) where N is a the number
//of keys and k the number of threads
func (tree *ParralelTree) Fetch(key int) (*Node, error) {

	//Fetch the mother tree
	//Here we use floorKey as the mother tree only
	//contains a subset of the all the keys.
	//Consequently, we identify the nearest mother key
	//and use it to drill down further
	node, err := tree.floorKey(key, tree.root)
	if err == nil {
		//Fetch the subtree brought back by the mother tree
		return node.Value.(*Tree).fetch(key, node.Value.(*Tree).root)
	}
	return nil, errors.New("Key not found")
}

//FloorKey returns the neareast lowest node with regards to key
//Complexity is O(log(k) + log(n)/k) where N is a the number
//of keys and k the number of threads
func (tree *ParralelTree) FloorKey(key int) (*Node, error) {

	//Fetch the mother tree
	node, err := tree.floorKey(key, tree.root)
	if err == nil {
		//Fetch the subtree brought back by the mother tree
		return node.Value.(*Tree).floorKey(key, node.Value.(*Tree).root)
	}
	return nil, errors.New("Key not found")
}

//Add adds a new key at the right place
//Complexity is O(log(k) + log(n)/k) where N is a the number
//of keys and k the number of threads
func (tree *ParralelTree) Add(key int, value interface{}) {
	node, err := tree.floorKey(key, tree.root)
	if err == nil {
		//Fetch the subtree brought back by the mother tree
		node.Value.(*Tree).Add(key, value)
	}
}
