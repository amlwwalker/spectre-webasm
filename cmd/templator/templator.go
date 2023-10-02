package main

import (
	"fmt"
	"github.com/amlwwalker/spectre-webasm/pkg/spectre/components"
	"github.com/amlwwalker/spectre-webasm/pkg/spectre/layouts"
	"github.com/lithammer/shortuuid"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/yosssi/gohtml"
	"golang.org/x/net/html"
	"log"
	"strings"
)

const template = `<div class="container">
		<navbar data-brand="Spectre.css" data-doc-link="Docs" data-git-link="Github" data-img-src="https://picturepan2.github.io/spectre/img/spectre-logo.svg"></navbar>
		<sidebar data-link1-href="#link1" data-link1-text="Link 1 Text" data-link2-href="#link2" data-link2-text="Link 2 Text">
			<accordionmenu></accordionmenu>
			<hero data-title="hello world" data-description="the description"></hero>
			<p>content</p>
		</sidebar>
	</div>`

// Element interface
type Element interface {
	String() string
	Generate() app.UI
}

// Base element holding common attributes
type BaseElement struct {
	Class string
	ID    string
}
type Text struct {
	Content string
}

func (t Text) String() string {
	return t.Content
}
func (t Text) Generate() app.UI {
	return app.Raw(string(t.String()))
}

type HTMLElement struct {
	TagName  string
	Attrs    []html.Attribute
	Children []Element
}

var attrToMethod = map[string]func(interface{}, string) interface{}{
	"id": func(ui interface{}, value string) interface{} {
		switch v := ui.(type) {
		case app.HTMLDiv:
			return v.ID(value)
		case app.HTMLHeader:
			return v.ID(value)
			// ... add more types as required
		}
		return ui
	},
	"class": func(ui interface{}, value string) interface{} {
		switch v := ui.(type) {
		case app.HTMLDiv:
			return v.Class(value)
		case app.HTMLHeader:
			return v.Class(value)
			// ... add more types as required
		}
		return ui
	},
	"href": func(ui interface{}, value string) interface{} {
		switch v := ui.(type) {
		case app.HTMLA: // If there's a separate interface for navigable elements
			return v.Href(value)
		}
		return ui
	},
	// ... add more mappings as required
}

func createUIFromTagName(tagName string) app.UI {
	switch tagName {
	case "div":
		return app.Div()
	case "a":
		return app.A()
	case "img":
		return app.Img()
	// ... add more mappings
	default:
		// Handle unsupported tags or return a default like app.Div()
		return app.Div()
	}
}
func (e HTMLElement) String() string {
	attributes := ""
	for _, attr := range e.Attrs {
		attributes += fmt.Sprintf(" %s=\"%s\"", attr.Key, attr.Val)
	}

	content := "<" + e.TagName + attributes + ">"
	for _, child := range e.Children {
		content += child.String()
	}
	content += "</" + e.TagName + ">"
	return content
}
func (e HTMLElement) Generate() app.UI {
	// Create the raw start tag with attributes
	attributes := ""
	for _, attr := range e.Attrs {
		attributes += fmt.Sprintf(" %s=\"%s\"", attr.Key, attr.Val)
	}
	startTag := app.Raw("<" + e.TagName + attributes + ">")

	// Convert child elements
	var childComponents []app.UI
	for _, child := range e.Children {
		if comp, ok := child.(Element); ok {
			childComponents = append(childComponents, comp.Generate())
		} else {
			childComponents = append(childComponents, app.Raw(child.String()))
		}
	}

	// Create the raw end tag
	endTag := app.Raw("</" + e.TagName + ">")

	// Combine the start tag, child components, and end tag
	return app.Div().Body(append([]app.UI{startTag}, append(childComponents, endTag)...)...)
}

type Div struct {
	BaseElement
	Children []Element
}

func (d Div) String() string {
	attributes := ""
	if d.Class != "" {
		attributes += fmt.Sprintf(" class=\"%s\"", d.Class)
	}
	if d.ID != "" {
		attributes += fmt.Sprintf(" id=\"%s\"", d.ID)
	}

	content := "<div" + attributes + ">"
	for _, child := range d.Children {
		content += child.String()
	}
	content += "</div>"
	return content
}
func (d Div) Generate() app.UI {
	// Start with an empty div
	divComponent := app.Div()

	// Add class if present
	if d.Class != "" {
		divComponent = divComponent.Class(d.Class)
	}
	//fmt.Println("added classes ", app.HTMLString(divComponent))
	// Add ID if present
	if d.ID != "" {
		divComponent = divComponent.ID(d.ID)
	}

	// Convert child elements
	var childComponents []app.UI
	for _, child := range d.Children {
		if comp, ok := child.(Element); ok {
			childComponents = append(childComponents, comp.Generate())
		} else {
			childComponents = append(childComponents, app.Raw(child.String()))
		}
	}

	// Add child components to div
	divComponent = divComponent.Body(childComponents...)
	//fmt.Println("generated div ", app.HTMLString(divComponent))
	return divComponent
}

type Navbar struct {
	BaseElement
	Brand    string
	DocLink  string
	GitLink  string
	ImageSrc string
}

func (n Navbar) String() string {
	return fmt.Sprintf(`
<header class="navbar">
  <section class="navbar-section">
    <a href="#" class="navbar-brand mr-2">%s</a>
    <a href="#" class="btn btn-link">%s</a>
    <a href="#" class="btn btn-link">%s</a>
  </section>
  <section class="navbar-center">
    <img src="%s" />
  </section>
  <section class="navbar-section">
    <div class="input-group input-inline">
      <input class="form-input" type="text" placeholder="search">
      <button class="btn btn-primary input-group-btn">Search</button>
    </div>
  </section>
</header>`, n.Brand, n.DocLink, n.GitLink, n.ImageSrc)
}
func (n Navbar) Generate() app.UI {
	return app.Header().Class("navbar").Body(
		app.Section().Class("navbar-section").Body(
			app.A().Href("#").Class("navbar-brand mr-2").Text(n.Brand),
			app.A().Href("#").Class("btn btn-link").Text(n.DocLink),
			app.A().Href("#").Class("btn btn-link").Text(n.GitLink),
		),
		app.Section().Class("navbar-center").Body(
			app.Img().Src(n.ImageSrc),
		),
		app.Section().Class("navbar-section").Body(
			app.Div().Class("input-group input-inline").Body(
				app.Input().Class("form-input").Type("text").Placeholder("search"),
				app.Button().Class("btn btn-primary input-group-btn").Text("Search"),
			),
		),
	)
}

type Accordion struct {
	BaseElement
	Menus []AccordionMenu
}

func (a Accordion) String() string {
	var menus string
	for _, menu := range a.Menus {
		menus += menu.String() + "\n"
	}
	return fmt.Sprintf(`<div class="accordion-container">
%s
</div>`, menus)
}
func (a Accordion) Generate() app.UI {
	var menus []app.UI
	for _, menu := range a.Menus {
		menus = append(menus, menu.Generate())
	}
	return components.Accordion(menus...)
}

type AccordionMenu struct {
	BaseElement
	PathID       string
	MenuTitle    string
	MenuElements []string
}

func (am AccordionMenu) String() string {
	uuid := shortuuid.New()

	var elements string
	for _, el := range am.MenuElements {
		elements += fmt.Sprintf(`<li class="menu-item"><a href="/%s/%s">%s</a></li>`, am.PathID, urlLink(el), el) + "\n"
	}

	return fmt.Sprintf(`
<div class="accordion">
	<input id="accordion-%s-%s" type="checkbox" name="docs-accordion-checkbox" hidden="true">
	<label class="accordion header c-hand" for="accordion-%s-%s">%s</label>
	<div class="accordion-body">
		<ul class="menu menu-nav">
%s
		</ul>
	</div>
</div>`, am.PathID, uuid, am.PathID, uuid, am.MenuTitle, elements)
}
func (am AccordionMenu) Generate() app.UI {
	return components.AccordionMenu(am.PathID, am.MenuTitle, am.MenuElements)
}

// A mock function for urlLink, replace with the actual logic if it's different
func urlLink(input string) string {
	return strings.ToLower(strings.ReplaceAll(input, " ", "-"))
}

type Hero struct {
	BaseElement
	Title       string
	Description string
}

func (h Hero) String() string {
	content := "<div class=\"hero bg-gray\">"
	content += "<div class=\"hero-body\">"
	content += fmt.Sprintf("<h1>%s</h1>", h.Title)
	content += fmt.Sprintf("<p>%s</p>", h.Description)
	content += "</div>"
	content += "</div>"
	return content
}
func (h Hero) Generate() app.UI {
	return layouts.Hero(h.Title, h.Description)
}

type FlexBox struct {
	Width string
	Body  Element
}

func (f FlexBox) String() string {
	return fmt.Sprintf(`
<div class="container" style="max-width: %s;">
	<div class="columns">
		%s
	</div>
</div>
`, f.Width, f.Body.String())
}
func (f FlexBox) Generate() app.UI {
	var bodyComponent app.UI
	if comp, ok := f.Body.(Element); ok {
		bodyComponent = comp.Generate()
	} else {
		bodyComponent = app.Raw(f.Body.String())
	}

	return layouts.FlexBox(f.Width, bodyComponent)
}

type Sidebar struct {
	BaseElement
	Links []Link
	Body  []Element
}

type Link struct {
	Href string
	Text string
}

func (s Sidebar) String() string {
	// Construct the sidebar links
	sidebarContent := ""
	for _, link := range s.Links {
		sidebarContent += fmt.Sprintf(`<a href="%s">%s</a>`, link.Href, link.Text)
	}

	// Construct the body content
	bodyContent := ""
	for _, element := range s.Body {
		bodyContent += element.String()
	}

	sidebarHtml := fmt.Sprintf(`
<div class="docs-container off-canvas off-canvas-sidebar-show" style="min-height: 100vh;">
	<a class="off-canvas-toggle btn btn-primary btn-action" href="#sidebar-id">
		<i class="icon icon-menu"></i>
	</a>
	<div id="sidebar-id" class="docs-sidebar off-canvas-sidebar flex-centered">
		<div class="docs-nav">
			%s
		</div>
	</div>
	<a class="off-canvas-overlay" href="#close"></a>
	<div class="off-canvas-content">
		%s
	</div>
</div>
`, sidebarContent, bodyContent)

	return FlexBox{
		Width: "100%",
		Body:  HTMLStringElement(sidebarHtml), // An adapter to treat raw HTML string as an Element
	}.String()
}
func (s Sidebar) Generate() app.UI {
	// Construct the sidebar links for Go-app
	var sidebarComponents []app.UI
	for _, link := range s.Links {
		sidebarComponents = append(sidebarComponents, app.A().Href(link.Href).Text(link.Text))
	}

	// Construct the body content for Go-app
	var bodyComponents []app.UI
	for _, element := range s.Body {
		if comp, ok := element.(Element); ok {
			bodyComponents = append(bodyComponents, comp.Generate())
		} else {
			bodyComponents = append(bodyComponents, app.Raw(element.String()))
		}
	}

	return layouts.SideBar(
		app.Div().Body(sidebarComponents...),
		app.Div().Body(bodyComponents...),
	)
}

// An adapter to treat raw HTML string as an Element
type HTMLStringElement string

func (h HTMLStringElement) String() string {
	return string(h)
}
func (h HTMLStringElement) Generate() app.UI {
	return app.Raw(string(h))
}

// ... [Other Element Structs like Navbar, Sidebar, etc. here]

func getAttribute(n *html.Node, key string) string {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func processNode(n *html.Node) Element {
	if n.Type == html.DocumentNode {
		// For the root document node, just process its children
		return HTMLElement{
			TagName:  n.Data,
			Attrs:    n.Attr,
			Children: processChildren(n),
		}
	} else if n.Type == html.ElementNode {
		base := BaseElement{
			Class: getAttribute(n, "class"),
			ID:    getAttribute(n, "id"),
		}
		fmt.Println("processing ", n.Data)
		switch n.Data {
		case "div":
			fmt.Println("found div")
			return Div{BaseElement: base, Children: processChildren(n)}
		case "navbar":
			return Navbar{
				BaseElement: base,
				Brand:       getAttribute(n, "data-brand"),
				DocLink:     getAttribute(n, "data-doc-link"),
				GitLink:     getAttribute(n, "data-git-link"),
				ImageSrc:    getAttribute(n, "data-img-src"),
			}
		case "sidebar":
			// Extract links from data attributes (you can extend this logic for more links)
			link1 := Link{Href: getAttribute(n, "data-link1-href"), Text: getAttribute(n, "data-link1-text")}
			link2 := Link{Href: getAttribute(n, "data-link2-href"), Text: getAttribute(n, "data-link2-text")}
			// ... add more links as necessary

			return Sidebar{
				BaseElement: BaseElement{
					Class: getAttribute(n, "class"),
					ID:    getAttribute(n, "id"),
				},
				Links: []Link{link1, link2}, // ... add more links as necessary
				Body:  processChildren(n),
			}
		case "accordionmenu":
			return Accordion{
				Menus: []AccordionMenu{
					{
						PathID:       "getting-started",
						MenuTitle:    "Getting Started",
						MenuElements: []string{"Installation", "Custom Version", "Browser-Support"},
					},
					//... other predefined menus
				},
			}
		case "hero":
			title := getAttribute(n, "data-title")
			desc := getAttribute(n, "data-description")
			return Hero{
				BaseElement: BaseElement{
					Class: getAttribute(n, "class"),
					ID:    getAttribute(n, "id"),
				},
				Title:       title,
				Description: desc,
			}
			// Add cases for other tags (AccordionMenu, Hero, etc.)
			// Add cases for other specific tags (AccordionMenu, Hero, etc.)
		default:
			// Fallback to the generic HTMLElement
			return HTMLElement{
				TagName:  n.Data,
				Attrs:    n.Attr,
				Children: processChildren(n),
			}
		}

	} else if n.Type == html.TextNode {
		return Text{Content: n.Data}
	}
	return nil
}

func processChildren(n *html.Node) []Element {
	var children []Element
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childElement := processNode(c)
		if childElement != nil {
			children = append(children, childElement)
		}
	}
	return children
}

func main() {
	var input = `
<div class="container">
		<navbar data-brand="Spectre.css" data-doc-link="Docs" data-git-link="Github" data-img-src="https://picturepan2.github.io/spectre/img/spectre-logo.svg"></navbar>
		<sidebar data-link1-href="#link1" data-link1-text="Link 1 Text" data-link2-href="#link2" data-link2-text="Link 2 Text">
			<accordionmenu></accordionmenu>
			<hero data-title="hello world" data-description="the description"></hero>
			<p>content</p>
		</sidebar>
	</div>`
	htmlContent := input

	r := strings.NewReader(htmlContent)
	doc, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	root := processNode(doc)
	//fmt.Println("root node is ", root.String())
	fmt.Println("generated node is ", root.Generate())
	//"github.com/yosssi/gohtml"
	fmt.Println(gohtml.Format(app.HTMLString(root.Generate())))
}
