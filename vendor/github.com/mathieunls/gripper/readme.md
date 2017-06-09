[![Build Status](https://travis-ci.org/MathieuNls/gripper.png)](https://travis-ci.org/MathieuNls/gripper)
[![GoDoc](https://godoc.org/github.com/MathieuNls/gripper?status.png)](https://godoc.org/github.com/MathieuNls/gripper)
[![codecov](https://codecov.io/gh/MathieuNls/gripper/branch/master/graph/badge.svg)](https://codecov.io/gh/MathieuNls/gripper)

# Go Performances Plotter

```go

import "github.com/mathieunls/gripper"

func main(){
    //max size of the dataset
    max := 10
    //increment dataset by 
    increment := 1
    //how many tries for each dataset size
    //we keep the median
    retries := 10
    //The name to be displayed on the graph
    label := "nothing"

    gripper.PerfPlotter().
    Analyze(
        myFunctionToAnalyze,
        max,
        increment,
        retries,
        label,
    ).
    Analyze(
        anotherFunction,
        max,
        increment,
        retries,
        "more nothing",
    ).
    Plot("x", "y", "x/y", "testing.png")
}

func myFunctionToAnalyze(datasetSize int){

}

func anotherFunction(datasetSize int){

}
```

will create 

![testing](https://user-images.githubusercontent.com/7218861/26978307-0e3def36-4cf9-11e7-8527-3bf9936d8ae9.png)
