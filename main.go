package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/adrg/frontmatter"
	"github.com/amlwwalker/spectre-webasm/pkg/markdown"
	"github.com/amlwwalker/spectre-webasm/pkg/spectre/components"
	"github.com/amlwwalker/spectre-webasm/pkg/spectre/layouts"
	pages "github.com/amlwwalker/spectre-webasm/pkg/spectre/pages"
	spectre_components "github.com/amlwwalker/spectre-webasm/pkg/spectre/pages/components"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

//
//// homePage pages
//type homePage struct {
//	app.Compo
//}
//
//// homePage render method
//func (h *homePage) Render() app.UI {
//	res, err := http.Get("/links")
//	if err != nil {
//		return app.Div()
//	}
//	all, err := io.ReadAll(res.Body)
//	if err != nil {
//		return app.Div()
//	}
//	var links []link
//	json.Unmarshal(all, &links)
//	fmt.Println("links", links)
//	return app.Div().Class("container").Body(
//		layouts.NavBar(),
//		layouts.SideBar(
//			createSideBar(links),
//			app.Div().Class("docs-content").Class("content").Body(
//				//components.Hero("Spectre (wasm)", "Using the Spectre css library in wasm with Go!"),
//				layouts.MediaHero("Spectre (wasm)", "Using the Spectre css library in wasm with Go!", "cbB3QEwWMlA"),
//				//markdown.NewRemoteMarkdownDoc().Src("/documents/home.md"),
//				components.Steps(),
//				app.Div().Class("docs-content").Class("content").Body(
//					//components.Hero("Spectre (wasm)", "Using the Spectre css library in wasm with Go!"),
//					layouts.MediaHero("Spectre (wasm)", "Using the Spectre css library in wasm with Go!", "cbB3QEwWMlA"),
//					//markdown.NewRemoteMarkdownDoc().Src("/documents/home.md"),
//					components.Steps(),
//				),
//				components.CustomComponent("ComponentA", map[string]interface{}{
//					"data": "Some data for ComponentA",
//				}),
//				components.Toast("Lorem ipsum dolor sit amet, consectetur adipiscing elit.", "warning", true, true),
//			),
//		),
//	)
//}

// The main function is the entry point where the app is configured and started.
// It is executed in 2 different environments: A client (the web browser) and a
// server.
func main() {
	// The first thing to do is to associate the hello component with a path.
	//
	// This is done by calling the Route() function,  which tells go-app what
	// component to display for a given path, on both client and server-side.

	homePage := pages.NewPage("Home", "Home Page", "/documents/home.md", []app.UI{},
		components.Steps(),
		//components.CustomComponent("ComponentA", map[string]interface{}{
		//	"data": "Some data for ComponentA",
		//}),
		//components.Toast("Lorem ipsum dolor sit amet, consectetur adipiscing elit.", "warning", true, true)
	)
	//fmt.Println("homepage ", homePage.Title, homePage.Src)
	//fmt.Println("rendering page 1", homePage)
	installationPage := pages.NewPage("Getting Started", "Getting Started with Spectre for Web Assembly (in Go)", "/documents/installation.md", []app.UI{}, []app.UI{}...)
	customVersionPage := pages.NewPage("Custom Versions", "Customising Spectre in Web Assembly", "/documents/customVersion.md", []app.UI{}, []app.UI{}...)
	//
	heroPage := pages.NewPage("Hero", "Heros are large title blocks", "/documents/heros.md", []app.UI{},
		app.H2().Text("An example hero"),
		layouts.Hero("An example title", "an example description"),
	)
	//
	badgePage := pages.NewPage("Badges", "Badges are often used as unread number indicators", "/documents/badges.md", []app.UI{},
		components.NotificationBadge("notification", 7),
		components.ButtonBadge("button", 7, nil),
		components.FigureBadge("https://picturepan2.github.io/spectre/img/avatar-3.png", 7, nil),
	)
	accordionPage := pages.NewPage("Accordions", "Accordions are used to toggle sections of content", "/documents/accordions.md", []app.UI{},
		components.Accordion(
			components.AccordionMenu("getting-started", "Getting Started", []string{"Installation", "Custom Version", "Browser-Support"}),
			components.AccordionMenu("elements", "Elements", []string{"Typography", "Tables", "Buttons", "Forms", "Icons.css", "Labels", "Code", "Media"}),
			components.AccordionMenu("layouts", "Layouts", []string{"Flexbox grid", "Responsive", "Hero", "Navbar"}),
		),
	)
	cardPage := pages.NewPage("Cards", "Cards hold succinct information related to A specific topic", "/documents/cards.md", []app.UI{},
		layouts.TwoColumn(
			components.Card("Apple", "Apple desktop background", "https://picturepan2.github.io/spectre/img/osx-yosemite.jpg",
				app.Span().Body(
					app.Span().Text("To make A contribution to the world by making tools for the mind that advance humankind."),
				),
				app.Span().Body(
					app.Button().Class("btn btn-primary").Text("Card Button"),
				),
			),
			components.Card("Microsoft", "Microsoft desktop background", "https://picturepan2.github.io/spectre/img/osx-el-capitan.jpg",
				app.Span().Body(
					app.Span().Text("Empower every person and every organization on the planet to achieve more."),
				),
				app.Span().Body(
					app.Button().Class("btn btn-primary").Text("Card Button"),
				),
			),
		),
	)
	panelPage := pages.NewPanelPage()
	tilePage := pages.NewPage("Tiles", "Tiles are repeatable or embeddable information blocks.", "/documents/tiles.md", []app.UI{},
		app.Div().Class("columns").Body(
			app.Div().Class("column col-9 col-sm-12").Body(
				components.Tile("The Avengers", "Earth's Mightiest Heroes joined forces to take on threats that were too big for any one hero to tackle...", "Join", "https://picturepan2.github.io/spectre/img/avatar-1.png"),
				components.Tile("The Avengers", "Earth's Mightiest Heroes joined forces to take on threats that were too big for any one hero to tackle...", "Join", "https://picturepan2.github.io/spectre/img/avatar-1.png"),
			),
		),
	)
	avatarPage := pages.NewPage("Avatars", "Avatars are personal icons next to details or people", "/documents/avatars.md", []app.UI{},
		layouts.TwoColumn(
			components.Avatar("avatar-xl", "https://picturepan2.github.io/spectre/img/avatar-1.png", "AW", "#5755d9"),
			components.Avatar("avatar-lg", "", "AW", "#5755d9"),
			components.Avatar("avatar-sm", "https://picturepan2.github.io/spectre/img/avatar-1.png", "AW", "#5755d9"),
			components.Avatar("avatar-xs", "", "AW", "#5755d9"),
		),
	)
	tabsPage := pages.NewTabsPage()
	breadCrumbsPage := pages.NewBreadcrumbPage()
	chipsPage := pages.NewPage("Chips", "Chips are complex entities in small blocks.", "/documents/chips.md", []app.UI{},
		layouts.FlexBox("80%",
			components.Chip("Crime", "", "", false),
			components.Chip("Crime", "avatar-sm", "https://picturepan2.github.io/spectre/img/avatar-1.png", false),
			components.Chip("Crime", "avatar-sm", "https://picturepan2.github.io/spectre/img/avatar-1.png", true),
		),
	)
	emptyStatePage := pages.NewPage("Empty States", "Empty states/blank slates are commonly used as placeholders for first time use, empty data and error screens.", "/documents/empty.md", []app.UI{},
		app.Br(),
		components.EmptyState("icon-people", "An Empty State", "An example empty state", "btn-primary", "does nothing", nil),
		app.Br(),
		components.EmptyState("icon-3x icon-mail", "You have no new messages", "Click the button to start a conversation", "btn-primary", "Send a message", nil),
	)
	tablePage := pages.NewTablePage()
	buttonPage := pages.NewButtonPage()
	customComponentPage := pages.NewPage("React Custom Components", "Custom React Components are possible!", "/documents/custom-components.md", []app.UI{},
		components.CustomComponent("ComponentA", map[string]interface{}{
			"field": "data",
			"data":  "Some data for ComponentA",
		}),
	)
	app.RouteFunc("/", pages.NewPageComposer(&homePage))
	app.RouteFunc("/getting-started/installation", pages.NewPageComposer(&installationPage))
	app.RouteFunc("/getting-started/customVersion", pages.NewPageComposer(&customVersionPage))
	//
	app.RouteFunc("/layouts/heros", pages.NewPageComposer(&heroPage))

	app.Route("/components/modals", &spectre_components.ModalPage{})
	app.RouteFunc("/components/badges", pages.NewPageComposer(&badgePage))
	app.RouteFunc("/components/accordions", pages.NewPageComposer(&accordionPage))

	app.RouteFunc("/components/cards", pages.NewPageComposer(&cardPage))
	app.RouteFunc("/components/panels", pages.NewPageComposer(&panelPage))
	app.RouteFunc("/components/tiles", pages.NewPageComposer(&tilePage))
	app.RouteFunc("/components/avatars", pages.NewPageComposer(&avatarPage))
	app.RouteFunc("/components/tabs", pages.NewPageComposer(&tabsPage))
	app.RouteFunc("/components/breadcrumbs", pages.NewPageComposer(&breadCrumbsPage))
	app.RouteFunc("/components/chips", pages.NewPageComposer(&chipsPage))
	app.RouteFunc("/components/empty", pages.NewPageComposer(&emptyStatePage))

	app.RouteFunc("/elements/tables", pages.NewPageComposer(&tablePage))
	app.RouteFunc("/elements/buttons", pages.NewPageComposer(&buttonPage))

	app.RouteFunc("/react/customComponents", pages.NewPageComposer(&customComponentPage))
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

	////todo - this is having issues wi
	//err := app.GenerateStaticWebsite("./static", &app.Handler{
	//	Name:        "Hello",
	//	Description: "An Hello World! example",
	//})
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//return
	// Finally, launching the server that serves the app is done by using the Go
	// standard HTTP package.
	//
	// The Handler is an HTTP handler that serves the client and all its
	// required resources to make it work into a web browser. Here it is
	// configured to handle requests with a path that starts with "/".

	http.HandleFunc("/retrieve/", retrieveMarkdown)
	http.HandleFunc("/links", func(w http.ResponseWriter, r *http.Request) {
		links, err := pages.GenerateLinks("./documents")
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to generate links: %v", err), http.StatusInternalServerError)
			return
		}

		for _, l := range links {
			fmt.Println(l.Name, l.URL, l.Group)
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(links); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}
	})

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
		Scripts: []string{
			"/web/dist/app.bundle.js",
		},
		RawHeaders:         []string{`<script src="/web/js/prism.js" data-manual></script>`},
		CacheableResources: []string{"/web/js/prism.js"},
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

type matter struct {
	Group string   `yaml:"group"` //if no group it goes top level
	Name  string   `yaml:"name"`  //the file name currently
	Roles []string `yaml:"roles"` //the permissions
}

type link struct {
	Name  string
	Group string // You can decide how to populate this as it wasn't clear from the description
	URL   string
}

func retrieveMarkdown(w http.ResponseWriter, r *http.Request) {
	// Extract the path after /retrieve/
	path := strings.TrimPrefix(r.URL.Path, "/retrieve/")

	// Read the file from the given path
	content, err := os.ReadFile(path)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	role := "user"
	var m matter
	rest, err := frontmatter.Parse(bytes.NewReader(content), &m)
	if err != nil {
		http.Error(w, "Error in file", http.StatusNotFound)
		return
	}
	result, err := pages.CompareRoles(role, m.Roles)
	if err != nil {
		http.Error(w, "No Role", http.StatusBadRequest)
		return
	}
	if result < 0 {
		fmt.Println("user does not have access to this file")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Return the file content as plain text (or markdown type if you have one)
	w.Header().Set("Content-Type", "text/plain")
	w.Write(rest)
}
