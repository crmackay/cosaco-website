package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// process template

// import data

// assemble template

// output html

func main() {
	inBytes, err := ioutil.ReadFile("dominican.csv")
	if err != nil {
		fmt.Println("error reading file")
	}

	myJSON := cSVtoPlainJSON(inBytes)

	// fmt.Println(myJSON)
	myStructs := plainJSONtoStruct([]byte(myJSON))
	// fmt.Println(myStructs)

	cleanStructs := make([][]question, len(myStructs))

	for i := range myStructs {
		cleanStructs[i] = collapseSubs(myStructs[i])
	}

	// for i := range cleanStructs {
	// 	for j := range cleanStructs[i] {
	// 		//fmt.Println(cleanStructs[i][j].Subs)
	// 	}
	// }
	myList := cleanStructs
	page := makePage(myList)

	fmt.Println(page)
}

func collapseSubs(in []question) []question {
	//fmt.Println(in)
	var result []question
	var currPart string
	var stemNum int
	for i := range in {
		part := strings.Split(in[i].Stem, ".")[0]
		// fmt.Println(currPart, part)
		if part == currPart {
			if len(result[stemNum-1].Subs) == 0 {
				result[stemNum-1].Subs = append(result[stemNum-1].Subs, result[stemNum-1])
				result[stemNum-1].Stem = part
				result[stemNum-1].Resp = ""
			}

			result[stemNum-1].Subs = append(result[stemNum-1].Subs, in[i])
		} else {
			result = append(result, in[i])
			currPart = part
			stemNum++
		}
	}
	// fmt.Println(result)
	return result
}
