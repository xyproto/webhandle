webhandle
=========

Serve webpages with [instapage](https://github.com/xyproto/instapage), [web.go](https://github.com/hoisie/web) and [onthefly](https://github.com/xyproto/onthefly).

I'm planning to support [negroni](https://github.com/codegangsta/negroni) or [martini](https://github.com/go-martini/martini) instead of [web.go](https://github.com/hoisie/web) in the future.

Example
-------

``` go
package main

import (
	"io/ioutil"

	"github.com/hoisie/web"
	"github.com/xyproto/onthefly"
	"github.com/xyproto/instapage"
	"github.com/xyproto/webhandle"
)

func notFound(ctx *web.Context, message string) {
	ctx.NotFound("Page not found")
}

func errorlog() string {
	data, err := ioutil.ReadFile("error.log")
	if err != nil {
		return "No errors\n"
	}
	return "Errors:\n" + string(data) + "\n"
}

func hello(val string) string {
	return instapage.Message("root page", "hello: "+val)
}

func exampleSVG() string {
	return onthefly.SampleSVG2().String()
}

func main() {
	webhandle.PublishPage("/", "/main.css", onthefly.SamplePage)

	web.Get("/error", errorlog)
	web.Get("/hello/(.*)", hello)

	web.Get("/svg", exampleSVG)

	web.Get("/(.*)", notFound)

	web.Run(":3000")
}
```
