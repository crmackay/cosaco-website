package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func cSVtoPlainJSON(in []byte) string {

	r := csv.NewReader(strings.NewReader(string(in)))
	out := ""
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	out += "["
	for j := 1; j < len(records); j++ {
		out += `[`
		for i := 0; i < len(records[0]); i++ {
			out += `{`
			out += `"question":"`
			out += records[0][i]
			out += `","response":"`
			out += records[j][i]
			out += `"}`
			if i < len(records[0])-1 {
				out += `,`
			}

		}
		out += `]`
		if j < len(records)-1 {
			out += `,`
		}
	}
	out += "]"
	return out
}

type question struct {
	Stem string     `json:"question"`
	Resp string     `json:"response"`
	Subs []question `json:"subquestions"`
}

func plainJSONtoStruct(in []byte) [][]question {

	var qs = new([][]question)
	//fmt.Println(string(in[0:20]))
	err := json.Unmarshal(in, qs)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(*qs)
	return *qs

}

func plainJSONtoMap(in string) []map[string]interface{} {
	var f []interface{}
	inJSON := []byte(in)
	err := json.Unmarshal(inJSON, &f)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(f)

	maps := make([]map[string]interface{}, len(f))

	for _, elem := range f {
		m := elem.(map[string]interface{})
		maps = append(maps, m)
	}

	// for _, elem := range f {
	//
	// 	for k, v := range elem {
	// 		switch vT := v.(type) {
	// 		case string:
	// 			fmt.Println(k, "is string", vT)
	// 		case []interface{}:
	// 			fmt.Println(k, "is an array:")
	// 			for i, u := range vT {
	// 				fmt.Println(i, u)
	// 			}
	// 		}
	// 	}
	// }

	return maps
}

//put answer together with same number into the same thing:

// first get column name: to response map...
// { response1: {"column 1": "this is a piece","column 2": "of some more","column 3": "data to read",},
//   response2: {"column 1": "and then","column 2": "this is more","column 3": "data for reading",},
// }
//then create a sub map[string]string for responses with
// same starting number (eg"3. ")

// strip off the number and name, store the name as "name"

// creat another map[string]string of responses
//
//

// {
//     "actividades": {
//         "name": "Actividades",
//         "responeses": {
//             "Attencion a la salud general": "",
//             "Salud Materna-Infantil":
//             "Sida/VIH Sanidad":
//             "Cuidados Emergencia":
//             "Prevencion (General)":
//             "Educacion d Sanidad General":
//         }
//     }
// }

// m := f.(map[string]interface{})
