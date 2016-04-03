package main

import (
	"fmt"
	"testing"
)

func TestCSVtoplainJSON(t *testing.T) {

	testInput := `column 1,column 2,column 3
this is a piece,of some more,data to read
and then,this is more,data for reading`

	result := cSVtoPlainJSON(testInput)
	fmt.Println(result)
	t.Fail()

}
