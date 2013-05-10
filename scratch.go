package main

import (
	"github.com/murz/go-handlebars/handlebars"
	"fmt"
)

func main() {
	handlebars.RegisterHelper("foo", func(params ...interface{}) string {
          return fmt.Sprintf("<h2>%v</h2> <p>%v</p>", params[0], params[1])
        })
	handlebars.RegisterHelper("link_to", func(params ...interface{}) string {
		if str, ok := params[0].(string); ok {
		  return fmt.Sprintf("OK!! %v", str)
		}
		if ctx, ok := params[0].(map[string]string); ok {
		  return fmt.Sprintf("<a href='%v'>%v</a>", ctx["url"], ctx["body"]) 
		}
		return ""
	})
	handlebars.RegisterHelper("bar", func(params ...interface{}) string {
		if str, ok := params[1].(string); ok {
			return "test: " + str;
		}
		return "fail"
	})

        handlebars.RegisterHelper("content_for", func(params ...interface{}) string {
          if len(params) == 2 {
            handlebars.RegisterContent(fmt.Sprintf("%v", params[0]), fmt.Sprintf("%v", params[1]))
          }
          return ""
        })
        handlebars.RegisterHelper("yield", func(params ...interface{}) string {
          if len(params) != 2 {
            return ""
          }
          if tmpl, ok := handlebars.ContentFor(fmt.Sprintf("%v", params[0])); ok {
            return tmpl.Render()
          } else {
            return fmt.Sprintf("%v", params[1])
          }
          return ""
        })


	ctx := map[string]interface{}{
		"posts": []interface{}{
			map[string]string{"url": "/hello-world", "body": "Hello World!"},
		},
	}
	fmt.Printf("result:\n%v\n",
//		handlebars.Render(`<ul>{{#posts}}<li>{{#yield "foo"}}hello{{/yield}}</li>{{/posts}}</ul>`, ctx))
                handlebars.RenderFile("about.html.hbs", ctx))

}
