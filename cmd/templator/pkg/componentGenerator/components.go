package componentGenerator

import (
	"fmt"
	"github.com/amlwwalker/spectre-webasm/pkg/spectre/components"
	"github.com/amlwwalker/spectre-webasm/pkg/spectre/layouts"
	"github.com/lithammer/shortuuid"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Sidebar struct {
	BaseElement
	Links []Link
	Body  []Element
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
