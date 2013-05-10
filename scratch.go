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
		if fn, ok := params[1].(func() string); ok {
			return "test: " + fn();
		}
		return "fail"
	})
	ctx := map[string]interface{}{
		"posts": []interface{}{
			map[string]string{"url": "/hello-world", "body": "Hello World!"},
		},
	}
	fmt.Printf("result:\n%v\n",
		handlebars.Render(`<ul>{{#posts}}<li>{{#bar "foo"}}hello{{/bar}}</li>{{/posts}}</ul>`, ctx))
 
}
