[![Build Status](https://travis-ci.org/MathieuNls/gpbt.png)](https://travis-ci.org/MathieuNls/gpbt)
[![GoDoc](https://godoc.org/github.com/MathieuNls/gpbt?status.png)](https://godoc.org/github.com/MathieuNls/gpbt)
[![codecov](https://codecov.io/gh/MathieuNls/gpbt/branch/master/graph/badge.svg)](https://codecov.io/gh/MathieuNls/gpbt)

# Go Parallel Binary Trees 

This package provides a parallel binary tree where the the *Fetch* and *Add* complexities are  *O(log(k) + log(n)/k)* where N is a the number of keys and k the number of threads.

It works by creating a root-tree where each node is the root of a sub-tree. This divides the search space by k for the *Fetch* and *Add* operations as shown below.

![gpbt](https://user-images.githubusercontent.com/7218861/27291126-61fb35e8-54dd-11e7-8e50-5b1e1ac98a32.png)


To fetch the key 149 (on the right most sub-tree), we first search for the floor key of 149 which gives us 150. 
Then, we look for 149 inside the sub-tree. 

# Usage

```go
  
//Dumb data
ints := make([]int, 30)
values := make([]interface{}, 30)

for i := 0; i < 30; i++ {
	ints[i] = i * 3
	values[i] = strconv.Itoa(i * 3)
}
  
//New p-tree with 12 threads
tree := NewParralelTree(ints, values, 12)
  
//Fetch the node (beloging to a sub-tree) with key = 3
node, err = tree.Fetch(3)
  
//Add 99-"99"
tree.Add(99, "99")
```
