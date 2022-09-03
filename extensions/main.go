package main

import (
	"amlwwalker/go-app-tuts/extensions/pkg/ace"
	"amlwwalker/go-app-tuts/pkg/markdown"
	"amlwwalker/go-app-tuts/pkg/spectre/layouts"
	"fmt"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// homePage pages
type homePage struct {
	app.Compo
	activeTab string
}

// homePage render method
func (h *homePage) Render() app.UI {
	icons := []string{
		"icon-cross",
		"icon-check",
		"icon-stop",
		"icon-shutdown",
		"icon-refresh",
		"icon-search",
		"icon-flag",
		"icon-bookmark",
		"icon-edit",
		"icon-delete",
		"icon-share",
		"icon-download",
		"icon-upload",
		"icon-copy",
		"icon-arrow-up",
		"icon-arrow-right",
		"icon-arrow-down",
		"icon-arrow-left",
		"icon-upward",
		"icon-forward",
		"icon-downward",
		"icon-back",
		"icon-caret",
		"icon-menu",
		"icon-apps",
		"icon-more-horiz",
		"icon-more-vert",
	}
	return app.Div().Class("container").Body(
		layouts.NavBar(),
		layouts.SideBar(
			layouts.FlexBox("100%",
				app.Range(icons).Slice(func(i int) app.UI {
					return app.Div().Class("column").Body(
						app.I().Class("icon icon-2x").Class(icons[i]),
					).Style("padding", "0.5rem")
				}),
			),
			app.Div().Class("docs-content").Class("content").Body(
				app.Div().Class("column col-12").Body(
					ace.NewEditor(),
				),
			),
		),
		//components.Toast("Lorem ipsum dolor sit amet, consectetur adipiscing elit.", "warning", true, true),
	)
}
func printMessage(this app.Value, inputs[]app.Value) interface{} {
	fmt.Println("printing message")
	app.Window().Call("alert", "hello world")
	message := "hello world"//inputs[0].String()
	document := app.Window().Get("document")
	p := document.Call("createElement", "p")
	p.Set("innerHTML", message)
	document.Get("body").Call("appendChild", p)
	return nil
}
// The main function is the entry point where the app is configured and started.
// It is executed in 2 different environments: A client (the web browser) and a
// server.
func main() {
	app.Window().Set("go_printMessage", app.FuncOf(printMessage))
	// The first thing to do is to associate the hello component with a path.
	//
	// This is done by calling the Route() function,  which tells go-app what
	// component to display for a given path, on both client and server-side.
	app.Route("/", &homePage{})
	app.Handle(markdown.GetMarkdown, markdown.HandleGetMarkdown)
	// Once the routes set up, the next thing to do is to either launch the app
	// or the server that serves the app.
	//
	// When executed on the client-side, the RunWhenOnBrowser() function
	// launches the app,  starting a loop that listens for app events and
	// executes client instructions. Since it is a blocking call, the code below
	// it will never be executed.
	//
	// When executed on the server-side, RunWhenOnBrowser() does nothing, which
	// lets room for server implementation without the need for precompiling
	// instructions.
	app.RunWhenOnBrowser()

	// Finally, launching the server that serves the app is done by using the Go
	// standard HTTP package.
	//
	// The Handler is an HTTP handler that serves the client and all its
	// required resources to make it work into a web browser. Here it is
	// configured to handle requests with a path that starts with "/".
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
		Styles: []string{
			"/web/css/spectre/spectre.min.css", // Loads hello.css file.
			"/web/css/spectre/spectre-exp.min.css",
			"/web/css/spectre/spectre-icons.min.css",
			"/web/css/docs.css",
			"/web/css/popover-styles.css",
			"/web/css/editor.css",
		},
		Scripts: []string{
			//"https://cdnjs.cloudflare.com/ajax/libs/ace/1.9.6/ace.js",
			"/web/js/ace-src/ace.js",
			"/web/js/ace-src/theme-solarized_dark.js",
		},
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
