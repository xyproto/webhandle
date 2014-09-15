package webhandle

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/drbawb/mustache"
	"github.com/hoisie/web"
	"github.com/xyproto/instapage"
	. "github.com/xyproto/onthefly"
)

type (
	// Various functiomn signatures for handling requests
	WebHandle                 (func(ctx *web.Context, val string) string)
	SimpleContextHandle       (func(ctx *web.Context) string)
	OldTemplateValueGenerator func(*web.Context) TemplateValues
)

// Create a web.go compatible function that returns a string that is the HTML for this page
func OldGenerateHTML(page *Page) func(*web.Context) string {
	return func(ctx *web.Context) string {
		return page.GetXML(true)
	}
}

// Create a web.go compatible function that returns a string that is the HTML for this page
func OldGenerateHTMLwithTemplate(page *Page, tvg OldTemplateValueGenerator) func(*web.Context) string {
	return func(ctx *web.Context) string {
		values := tvg(ctx)
		return mustache.Render(page.GetXML(true), values)
	}
}

// Create a web.go compatible function that returns a string that is the CSS for this page
func OldGenerateCSS(page *Page) func(*web.Context) string {
	return func(ctx *web.Context) string {
		ctx.ContentType("css")
		return page.GetCSS()
	}
}

// Create a web.go compatible function that returns a string that is the XML for this page
func OldGenerateXML(page *Page) func(*web.Context) string {
	return func(ctx *web.Context) string {
		ctx.ContentType("xml")
		return page.GetXML(false)
	}
}

// Creates a page based on the contents of "error.log". Useful for showing compile errors while creating an application.
func OldGenerateErrorHandle(errorfilename string) SimpleContextHandle {
	return func(ctx *web.Context) string {
		data, err := ioutil.ReadFile(errorfilename)
		if err != nil {
			return instapage.Message("Good", "No errors")
		}
		errors := strings.Replace(string(data), "\n", "</br>", -1)
		return instapage.Message("Errors", errors)
	}
}

// Handles pages that are not found
func NotFound(ctx *web.Context, val string) string {
	ctx.NotFound(instapage.Message("No", "Page not found"))
	return ""
}

// Takes a filename and returns a function that can handle the request
func OldFile(filename string) func(ctx *web.Context) string {
	var extension string
	if strings.Contains(filename, ".") {
		extension = filepath.Ext(filename)
	}
	return func(ctx *web.Context) string {
		if extension != "" {
			ctx.ContentType(extension)
		}
		imagebytes, _ := ioutil.ReadFile(filename)
		buf := bytes.NewBuffer(imagebytes)
		return buf.String()
	}
}

// Takes an url and a filename and offers that file at the given url
func OldPublishFile(url, filename string) {
	web.Get(url, File(filename))
}

// Takes a filename and offers that file at the root url
func OldPublishRootFile(filename string) {
	web.Get("/"+filename, File(filename))
}

// Expose the HTML and CSS generated by a page building function to the two given urls
func OldPublishPage(htmlurl, cssurl string, buildfunction func(string) *Page) {
	page := buildfunction(cssurl)
	web.Get(htmlurl, GenerateHTML(page))
	web.Get(cssurl, GenerateCSS(page))
}

// Wrap a SimpleContextHandle so that the output is cached (with an id)
// Do not cache functions with side-effects! (that sets the mimetype for instance)
// The safest thing for now is to only cache images.
func OldCacheWrapper(id string, f SimpleContextHandle) SimpleContextHandle {
	return func(ctx *web.Context) string {
		if _, ok := globalStringCache[id]; !ok {
			globalStringCache[id] = f(ctx)
		}
		return globalStringCache[id]
	}
}

func OldPublish(url, filename string, cache bool) {
	if cache {
		web.Get(url, OldCacheWrapper(url, OldFile(filename)))
	} else {
		web.Get(url, File(filename))
	}
}