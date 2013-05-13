package main

import (
  "github.com/murz/go-handlebars/handlebars"
  "fmt"
  "bytes"
)

func main() {
  var buf bytes.Buffer
  t := handlebars.Must(handlebars.New("test").Parse("text {{foo}}"))
  t.Execute(&buf, map[string]string{
    "foo": "bar",
  })
  fmt.Println("result: ", buf.String())
}
