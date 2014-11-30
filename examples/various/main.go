package main

import (
	"io/ioutil"

	"github.com/hoisie/web"
	"github.com/xyproto/instapage"
	"github.com/xyproto/onthefly"
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

	web.Run(":9080")
}
