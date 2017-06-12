package index

import (
	"fmt"
	"testing"

)

func TestNewFastIndex(t *testing.T) {

	f := &FastIndex{} 

	fmt.Println(f.Lookup([]byte("plopt"), -1))

}
