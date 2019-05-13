package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/xyproto/onthefly"
	"github.com/xyproto/webhandle"
)

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
	gmux := mux.NewRouter()

	webhandle.PublishPage(gmux, "/", "/main.css", onthefly.SamplePage)

	gmux.HandleFunc("/error", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "%s", errorlog())
	})

	gmux.HandleFunc("/hello/{name}", func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		name := vars["name"]
		fmt.Fprintf(w, "%s", hello(name))
	})

	gmux.HandleFunc("/svg", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "%s", sampleSVG())
	})

	n := negroni.Classic()
	n.UseHandler(gmux)

	fmt.Println("Serving on port 3000")
	http.ListenAndServe(":3000", n)
}
