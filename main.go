package main

import (
	"github.com/amlwwalker/spectre-webasm/pkg/markdown"
	"github.com/amlwwalker/spectre-webasm/pkg/spectre/components"
	"github.com/amlwwalker/spectre-webasm/pkg/spectre/layouts"
	spectre_components "github.com/amlwwalker/spectre-webasm/pkg/spectre/pages/components"
	"github.com/amlwwalker/spectre-webasm/pkg/spectre/pages/elements"
	getting_started "github.com/amlwwalker/spectre-webasm/pkg/spectre/pages/getting-started"
	page_layouts "github.com/amlwwalker/spectre-webasm/pkg/spectre/pages/layouts"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// homePage pages
type homePage struct {
	app.Compo
}

// homePage render method
func (h *homePage) Render() app.UI {
	return app.Div().Class("container").Body(
		layouts.NavBar(),
		layouts.SideBar(
				components.Accordion(
					components.AccordionMenu("getting-started", "Getting Started", []string{"Installation", "Custom Version", "Browser-Support"}),
					components.AccordionMenu("elements", "Elements", []string{"Typography", "Tables", "Buttons", "Forms", "Icons.css", "Labels", "Code", "Media"}),
					components.AccordionMenu("layouts", "Layouts", []string{"Flexbox grid", "Responsive", "Hero", "Navbar"}),
					components.AccordionMenu("components", "Components", []string{"Accordions", "Avatars", "Badges", "Bars", "Breadcrumbs", "Cards", "Chips", "Empty States", "Menu", "Modals", "Nav", 				"Pagination", "Panels", "Popovers", "Steps", "Tabs", "Tiles", "Toasts", "Tooltips"}),
				),
				app.Div().Class("docs-content").Class("content").Body(
			//components.Hero("Spectre (wasm)", "Using the Spectre css library in wasm with Go!"),
			layouts.MediaHero("Spectre (wasm)", "Using the Spectre css library in wasm with Go!", "cbB3QEwWMlA"),
			//markdown.NewRemoteMarkdownDoc().Src("/web/documents/home.md"),
			components.Steps(),
		),
		),
		components.Toast("Lorem ipsum dolor sit amet, consectetur adipiscing elit.", "warning", true, true),
	)
}

// The main function is the entry point where the app is configured and started.
// It is executed in 2 different environments: A client (the web browser) and a
// server.
func main() {
	// The first thing to do is to associate the hello component with a path.
	//
	// This is done by calling the Route() function,  which tells go-app what
	// component to display for a given path, on both client and server-side.
	app.Route("/", &homePage{})
	app.Route("/getting-started/installation", &getting_started.InstallationPage{})
	app.Route("/getting-started/customVersion", &getting_started.CustomVersionPage{})

	app.Route("/layouts/hero", &page_layouts.HeroPage{})


	app.Route("/components/modals", &spectre_components.ModalPage{})
	app.Route("/components/badges", &spectre_components.BadgePage{})
	app.Route("/components/accordions", &spectre_components.AccordionPage{})
	app.Route("/components/cards", &spectre_components.CardPage{})
	app.Route("/components/panels", &spectre_components.PanelPage{})
	app.Route("/components/tiles", &spectre_components.TilePage{})
	app.Route("/components/avatars", &spectre_components.AvatarPage{})
	app.Route("/components/tabs", &spectre_components.TabsPage{})
	app.Route("/components/breadcrumbs", &spectre_components.BreadcrumbPage{})
	app.Route("/components/chips", &spectre_components.ChipsPage{})
	app.Route("/components/emptyStates", &spectre_components.EmptyStatePage{})

	app.Route("/elements/tables", &elements.TablePage{})
	app.Route("/elements/buttons", &elements.ButtonPage{})

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
	app.Window().Get("Prism").Call("highlightAll")
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
			"/web/css/prism.css",
		},
		RawHeaders: []string{`<script src="/web/js/prism.js" data-manual></script>`},
		CacheableResources: []string{"/web/js/prism.js"},
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
