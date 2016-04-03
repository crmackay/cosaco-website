package main

import (
	"encoding/json"
	"fmt"
)

var testJSON = `[
    {
        "question": "this is a question",
        "response": "this is a response",
        "subquestions": ""
    },
    {
        "question": "this is another question",
        "response": "this is another response",
        "subquestions":
        [
            {
                "question": "this is a subquestion",
                "response": "this is a subresponse",
                "subquestions": ""
            },
            {
                "question": "this is a subquestion",
                "response": "this is a subresponse",
                "subquestions": ""
            }
        ]
    }
]`

type question struct {
	Ques string     `json:"question"`
	Resp string     `json:"response"`
	Subs []question `json:"subquestions"`
}

func main() {

	var t []question
	in := []byte(testJSON)
	json.Unmarshal(in, &t)

	fmt.Println(t)

	for i := range t {
		fmt.Println(t[i].Ques)
		fmt.Println(t[i].Resp)
		fmt.Println(t[i].Subs)
	}

}
