package main

import (
	"github.com/murz/go-handlebars/handlebars"
	"fmt"
)

func main() {
	handlebars.RegisterHelper("link_to", func(params ...interface{}) string {
		if str, ok := params[0].(string); ok {
		  return fmt.Sprintf("OK!! %v", str)
		}
		if ctx, ok := params[0].(map[string]string); ok {
		  return fmt.Sprintf("<a href='%v'>%v</a>", ctx["url"], ctx["body"]) 
		}
		return ""
	})


	ctx := map[string]interface{}{
                "foo2": map[string]string{
                  "bar": "barbarbar!",
                },
                "foo": "foooo!",
                "bar2": "barrrr",
                "foobar": func() string {
                  return "blaaaah"
                },
		"posts": []interface{}{
			map[string]string{"url": "/hello-world", "body": "Hello World!"},
		},
                "fp": func(s string) string {
                  return "yayyyyyy:"+s+":yyyyyy"
                },
	}
        fmt.Printf("input: %v\n\n", ctx)
	fmt.Printf("result:\n%v\n",
//		handlebars.Render(`<ul>{{#posts}}<li>{{#yield "foo"}}hello{{/yield}}</li>{{/posts}}</ul>`, ctx))
                handlebars.RenderFile("about.html.hbs", ctx))

}
