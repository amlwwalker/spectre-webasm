package components

import (
	"amlwwalker/go-app-tuts/pkg/markdown"
	"amlwwalker/go-app-tuts/pkg/spectre/components"
	"amlwwalker/go-app-tuts/pkg/spectre/layouts"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type page struct {
	app.Compo
}

func (p *page) Render() app.UI {
	return app.Span()
}

// Page
func Page(title, description, src string, body ...app.UI) app.HTMLDiv {
	return app.Div().Class("container").Body(
		layouts.NavBar(),
		layouts.SideBar(
			components.Accordion(
				components.AccordionMenu("getting-started", "Getting Started", []string{"Installation", "Custom Version", "Browser-Support"}),
				components.AccordionMenu("elements", "Elements", []string{"Typography", "Tables", "Buttons", "Forms", "Icons.css", "Labels", "Code", "Media"}),
				components.AccordionMenu("layouts", "Layouts", []string{"Flexbox grid", "Responsive", "Hero", "Navbar"}),
				components.AccordionMenu("components", "Components", []string{"Accordions", "Avatars", "Badges", "Bars", "Breadcrumbs", "Cards", "Chips", "Empty States", "Menu", "Modals", "Nav", "Pagination", "Panels", "Popovers", "Steps", "Tabs", "Tiles", "Toasts", "Tooltips"}),
			), app.Div().Class("docs-content").Class("content").Body(
				layouts.Hero(title, description),
				app.Span().Body(
					body...,
				),
				markdown.NewRemoteMarkdownDoc().Src(src),
			),
		),
		&page{},
	)
}
