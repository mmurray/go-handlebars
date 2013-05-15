package main

import (
  "github.com/murz/go-handlebars/handlebars"
  "fmt"
  "bytes"
)

func main() {
  var buf bytes.Buffer
  t := handlebars.Must(handlebars.New("test").Parse("text {{bar \"peaches\"}} "))
  t.Execute(&buf, map[string]interface{} {
  	"bar": func(str string) string {
  		return "boooo"+str+"oooo"
  	},
  	// "bar": map[string]interface{} {

  	// 	"blah": func(str string) string{
  	// 		return  "baz"+str+"z"
  	// 	},
  	// },
  })
  fmt.Println("result: ", buf.String())
}
