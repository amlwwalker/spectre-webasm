package components

import (
	"github.com/amlwwalker/spectre-webasm/pkg/spectre/components"
	"github.com/amlwwalker/spectre-webasm/pkg/spectre/elements"
	"github.com/amlwwalker/spectre-webasm/pkg/spectre/layouts"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func NewButtonPage() Page {
	p := NewPage("Buttons", "Buttons submit forms, or execute actions", "/documents/buttons.md", []app.UI{})
	const stateName = "button-click"
	p.States[stateName] = ""
	onClick := func(ctx app.Context, e app.Event) {
		app.Window().Call("alert", "setting state now...")
		ctx.SetState(stateName, "button was clicked. State was obsererved and monitored.")
	}
	p.SetBody(layouts.FlexBox("75%",
		app.Div().Class("column").Body(components.Button("block button", "btn btn-block", onClick)),
		app.Div().Class("column").Body(components.Button("primary button", "btn  btn-primary", onClick)),
		app.Div().Class("column").Body(components.Button("link button", "btn  btn-link", onClick)),
		app.Div().Class("column").Body(components.Button("succes button", "btn  btn-success", onClick)),
		app.Div().Class("column").Body(components.Button("error button", "btn  btn-error", onClick)),
		app.Div().Class("column").Body(components.Button("large error button", "btn  btn-error btn-lg", onClick)),
		app.Div().Class("column").Body(components.Button("small success button", "btn  btn-success btn-sm", onClick)),
	),
		app.Div().Body(
			layouts.FlexBox("75%",
				app.P().Text(p.States[stateName]),
			),
		),
	)
	return p
}
func NewTablePage() Page {
	t := []any{
		struct {
			A string
			B string
			C string
		}{
			A: "a",
			B: "b",
			C: "c1",
		},
		struct {
			C string
			D string
			E string
		}{
			C: "c2",
			D: "d",
			E: "e",
		},
	}
	p := NewPage("Tables", "Tables are used to display tabular data", "/documents/tables.md", []app.UI{},
		elements.Table(t, "table-striped"),
	)
	return p
}
func NewBreadcrumbPage() Page {
	p := NewPage("Breadcrumbs", "breadcrumbs help show and navigate around a site", "/documents/breadcrumbs.md", []app.UI{})
	p.Properties["activeTab"] = "home"
	el := []app.UI{
		app.A().Text("home").Href("#home").OnClick(func(ctx app.Context, e app.Event) {
			p.Properties["activeTab"] = "home"
		}),
		app.A().Text("settings").Href("#settings").OnClick(func(ctx app.Context, e app.Event) {
			p.Properties["activeTab"] = "settings"
		}),
		app.A().Text("change avatar").Href("#change-avatar").OnClick(func(ctx app.Context, e app.Event) {
			p.Properties["activeTab"] = "change-avatar"
		}),
	}
	p.SetBody(
		components.Breadcrumbs(el),
		app.If(p.Properties["activeTab"] == "home",
			components.Card("Home", "Home of user", "https://picturepan2.github.io/spectre/img/osx-yosemite.jpg",
				app.Span().Body(
					app.Span().Text("Home of a user to deal with"),
				),
				app.Span().Body(
					app.Button().Class("btn btn-primary").Text("Card Button"),
				),
			),
		),
		app.If(p.Properties["activeTab"] == "settings",
			components.Card("Settings", "Settings area", "https://picturepan2.github.io/spectre/img/osx-yosemite.jpg",
				app.Span().Body(
					app.Span().Text("User settings"),
				),
				app.Span().Body(
					app.Button().Class("btn btn-primary").Text("Card Button"),
				),
			),
		),
		app.If(p.Properties["activeTab"] == "change-avatar",
			components.Card("Avatar NewPage", "User avatar page", "https://picturepan2.github.io/spectre/img/osx-yosemite.jpg",
				app.Span().Body(
					app.Span().Text("Change your avatar"),
				),
				app.Span().Body(
					app.Button().Class("btn btn-primary").Text("Card Button"),
				),
			),
		),
	)
	return p
}
func NewTabsPage() Page {
	p := NewPage("Tabs", "Tabs group information together onto different tabs based on content or context", "/documents/tabs.md", []app.UI{})
	p.Properties["activeTab"] = "profile-tab" //pass in Properties potentialy?
	p.SetBody(app.Div().Class("column col-12").Body(
		components.Tabs(
			components.Tab("Profile", p.Properties["activeTab"] == "profile-tab", func(ctx app.Context, e app.Event) {
				p.Properties["activeTab"] = "profile-tab"
				p.Update()
			}),
			components.Tab("Files", p.Properties["activeTab"] == "files-tab", func(ctx app.Context, e app.Event) {
				p.Properties["activeTab"] = "files-tab"
				p.Update()
			}),
			components.Tab("Tasks", p.Properties["activeTab"] == "tasks-tab", func(ctx app.Context, e app.Event) {
				p.Properties["activeTab"] = "tasks-tab"
				p.Update()
			}),
		),
		app.If(p.Properties["activeTab"] == "profile-tab",
			components.Card("Profile", "Profile of user", "https://picturepan2.github.io/spectre/img/osx-yosemite.jpg",
				app.Span().Body(
					app.Span().Text("Profile of a user to deal with"),
				),
				app.Span().Body(
					app.Button().Class("btn btn-primary").Text("Card Button"),
				),
			),
		),
		app.If(p.Properties["activeTab"] == "files-tab",
			components.Card("Files", "Files that need handling", "https://picturepan2.github.io/spectre/img/osx-yosemite.jpg",
				app.Span().Body(
					app.Span().Text("Files that the user hasn't processed yet"),
				),
				app.Span().Body(
					app.Button().Class("btn btn-primary").Text("Card Button"),
				),
			),
		),
		app.If(p.Properties["activeTab"] == "tasks-tab",
			components.Card("Tasks", "Tasks that need doing", "https://picturepan2.github.io/spectre/img/osx-yosemite.jpg",
				app.Span().Body(
					app.Span().Text("Tasks the user is yet to deal with"),
				),
				app.Span().Body(
					app.Button().Class("btn btn-primary").Text("Card Button"),
				),
			),
		),
	))
	return p
}
func NewPanelPage() Page {
	avengersTiles := []app.UI{
		components.Tile(
			"The Avengers",
			"Earth's Mightiest Heroes joined forces to take on threats that were too big for any one hero to tackle...",
			"Join",
			"https://picturepan2.github.io/spectre/img/avatar-1.png"),
		components.Tile(
			"The Avengers",
			"Earth's Mightiest Heroes joined forces to take on threats that were too big for any one hero to tackle...",
			"Join",
			"https://picturepan2.github.io/spectre/img/avatar-1.png"),
		components.Tile(
			"The Avengers",
			"Earth's Mightiest Heroes joined forces to take on threats that were too big for any one hero to tackle...",
			"Join",
			"https://picturepan2.github.io/spectre/img/avatar-1.png"),
	}

	bannerProfile := []app.UI{
		components.Tile(
			"Email",
			"bruce.banner@hulk.com",
			"Edit",
			""),
		components.Tile(
			"Skype",
			"bruce.banner",
			"Edit",
			""),
		components.Tile(
			"Location",
			"Dayton, Idaho",
			"Edit",
			""),
	}
	return NewPage("Panels", "Panels are flexible view container with auto-expand content section.", "/documents/panels.md", []app.UI{},
		layouts.TwoColumn(
			components.Panel(app.Span().Text("Comments").Class("h4"), nil, nil, app.Span().Body(
				app.Range(avengersTiles).Slice(func(i int) app.UI {
					return avengersTiles[i]
				}),
			), app.Span().Text("footer content")).Style("max-height", "40vh"),
			components.Panel(app.Span().Text("Bruce Banner").Class("h4"),
				app.Span().Text("THE HULK").Class("h5"),
				components.Tabs(
					components.Tab("Profile", true, nil),
					components.Tab("Files", false, nil),
					components.Tab("Tasks", false, nil),
				),
				app.Span().Body(
					app.Range(bannerProfile).Slice(func(i int) app.UI {
						return bannerProfile[i]
					}),
				),
				components.Button("Save", "btn-block", nil),
			).Style("max-height", "40vh"),
		),
	)
}
