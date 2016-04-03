package main

import (
	"strconv"
)

func makePage(in [][]question) string {
	header := `<!DOCTYPE html>
<html lang="en">
<head>

  <!-- Basic Page Needs
  –––––––––––––––––––––––––––––––––––––––––––––––––– -->
  <meta charset="utf-8">
  <title>Your page title here :)</title>
  <meta name="description" content="">
  <meta name="author" content="">

  <!-- Mobile Specific Metas
  –––––––––––––––––––––––––––––––––––––––––––––––––– -->
  <meta name="viewport" content="width=device-width, initial-scale=1">

  <!-- FONT
  –––––––––––––––––––––––––––––––––––––––––––––––––– -->
  <link href="//fonts.googleapis.com/css?family=Raleway:400,300,600" rel="stylesheet" type="text/css">

  <!-- CSS
  –––––––––––––––––––––––––––––––––––––––––––––––––– -->
  <link rel="stylesheet" href="css/normalize.css">
  <link rel="stylesheet" href="css/skeleton.css">
  <link rel="stylesheet" href="css/mystyles.css">

  <!-- Favicon
  –––––––––––––––––––––––––––––––––––––––––––––––––– -->
<!--  <link rel="icon" type="image/png" href="images/favicon.png">-->

</head>
<body>

  <!-- Primary Page Layout
  –––––––––––––––––––––––––––––––––––––––––––––––––– -->

 <div class="container">
  <div class="row">
      <div class="twelve columns">
      <h1> Bienvenido a COSACO</h1>
      </div>
     </div>
   <div class="row">
       <div class="two columns"><a href="", target="_blank">¡Contáctenos!</a></div>
      <div class="ten columns">`
	footer := `     </div>
    </div>
</div>

<!-- End Document
  –––––––––––––––––––––––––––––––––––––––––––––––––– -->
</body>
    <script src="myjs.js"></script>
</html>`

	output := header

	for i := range in {
		output += `<div class="group">`
		for j, elem := range in[i] {
			if j == 0 {
				output += `<div class="name">`
				output += in[i][j].Resp
				output += `</div>`
				output += `<div class="data">
                <div class="data-primary">`
			} else if j != 0 && j < 4 {
				if len(in[i][j].Subs) > 0 {
					output += `<div class="data-question-set">`
					for _, elem := range in[i][j].Subs {
						output += `<div class="data-element">`
						output += `<div class="data-name">`
						output += elem.Stem
						output += `</div>`
						output += `<div class="data-value">`
						output += elem.Resp
						output += `</div>`
						output += `</div>`
					}
					output += `</div>`
				} else {
					output += `<div class="data-element">`
					output += `<div class="data-name">`
					output += elem.Stem
					output += `</div>`
					output += `<div class="data-value">`
					output += elem.Resp
					output += `</div>`
					output += `</div>`
				}
			} else if j == 4 {
				output += `</div>`
				output += `<button class="showData"  data-id="`
				index := i
				output += strconv.Itoa(index)
				output += `">Mas</button>`
				output += `<div id="`
				output += strconv.Itoa(index)
				output += `" class="data-secondary hidden">`
				if len(elem.Subs) > 0 {
					output += `<div class="data-question-set">`
					for _, elem := range elem.Subs {

						output += `<div class="data-question">
                                <div class="data-key">`
						output += elem.Stem
						output += `</div>`
						output += `<div class="data-value">`
						output += elem.Resp
						output += `</div>`
						output += `</div>`
					}
				} else {
					output += `<div class="data-question-set">`

					output += `<div class="data-question">
                                <div class="data-key">`
					output += elem.Stem
					output += `</div>`
					output += `<div class="data-value">`
					output += elem.Resp
					output += `</div>`
					output += `</div>`
				}

			} else if j > 4 {
				if len(elem.Subs) > 0 {
					output += `<div class="data-question-set">`
					for _, elem := range elem.Subs {

						output += `<div class="data-question">
                                <div class="data-key">`
						output += elem.Stem
						output += `</div>`
						output += `<div class="data-value">`
						output += elem.Resp
						output += `</div>`
						output += `</div>`
					}
				} else {
					output += `<div class="data-question-set">`

					output += `<div class="data-question">
                                <div class="data-key">`
					output += elem.Stem
					output += `</div>`
					output += `<div class="data-value">`
					output += elem.Resp
					output += `</div>`
					output += `</div>`
				}
				output += `</div>`

			}

		}
		output += `</div>`
		output += `</div>`

	}
	return output + footer
}
