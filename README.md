#webhandle [![Build Status](https://travis-ci.org/xyproto/webhandle.svg?branch=master)](https://travis-ci.org/xyproto/webhandle) [![GoDoc](https://godoc.org/github.com/xyproto/webhandle?status.svg)](http://godoc.org/github.com/xyproto/webhandle)

One way to serve webpages with [onthefly](https://github.com/xyproto/onthefly).

This is an experimental package. Other packages are likely to be better at the same tasks.

Online API Documentation
------------------------

[godoc.org](http://godoc.org/github.com/xyproto/webhandle)

Features and limitations
------------------------

* Webhandle can take a `*onthefly.Page` and publish both the HTML and CSS together, by listening to HTTP GET requests.
* There are also a few helper functions.

<!--
Example
-------

``` go
package main

import (
	"io/ioutil"

	"net/http"
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
	return webhandle.Message("root page", "hello: "+val)
}

func sampleSVG() string {
	return onthefly.SampleSVG2().String()
}

func main() {
	webhandle.PublishPage("/", "/main.css", onthefly.SamplePage)

	web.Get("/error", errorlog)
	web.Get("/hello/(.*)", hello)
	web.Get("/svg", sampleSVG)
	web.Get("/(.*)", notFound)

	web.Run(":3000")
}
```
-->

General information
-------------------

* Version: 0.1.1
* License: MIT
* Alexander F. Rødseth

