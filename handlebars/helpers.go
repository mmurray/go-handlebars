package handlebars

import (
  "fmt"
)

func init() {

        RegisterHelper("content_for", func(params ...interface{}) string {
          if len(params) == 2 {
            RegisterContent(fmt.Sprintf("%v", params[0]), fmt.Sprintf("%v", params[1]))
          }
          return ""
        })

        RegisterHelper("yield", func(params ...interface{}) string {
          fmt.Println("yield..", len(params))
          if len(params) != 2 {
            return ""
          }
          if tmpl, ok := ContentFor(fmt.Sprintf("%v", params[0])); ok {
            return tmpl.Render()
          } else {
            return fmt.Sprintf("%v", params[1])
          }
          return ""
        })

}
