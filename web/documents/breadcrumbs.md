# Breadcrumbs

Badges are often used to inform the user of a value or changing score.

Breadcrumbs are used as navigational hierarchies to indicate current location.

You can find the original Spectre [documentation here](https://picturepan2.github.io/spectre/components/breadcrumbs.html)

```go
el := []app.UI{
  app.A().Href("#").Text("home"),
  app.A().Href("#").Text("settings"),
  app.A().Href("#").Text("change avatar"),
}
return components.Breadcrumbs(el)
```

When you click on any of the links, either you can set the location to change the page or if you want to change the content you can do something similar to changing tabs

```go
type BreadcrumbPage struct {
	app.Compo
	activeTab string
}

func (h *BreadcrumbPage) OnMount(ctx app.Context) {
	h.activeTab = "change-avatar"
}
func (h *BreadcrumbPage) Render() app.UI {
	el := []app.UI{
		app.A().Text("home").Href("#home").OnClick(func(ctx app.Context, e app.Event) {
			h.activeTab = "home"
		}),
		app.A().Text("settings").Href("#settings").OnClick(func(ctx app.Context, e app.Event) {
			h.activeTab = "settings"
		}),
		app.A().Text("change avatar").Href("#change-avatar").OnClick(func(ctx app.Context, e app.Event) {
			h.activeTab = "change-avatar"
		}),
	}
	return pages.Page("Breadcrumbs", "breadcrumbs help show and navigate around a site", "/web/documents/breadcrumbs.md",
		components.Breadcrumbs(el),
		app.If(h.activeTab == "home",
			components.Card("Home", "Home of user", "https://picturepan2.github.io/spectre/img/osx-yosemite.jpg",
				app.Span().Body(
					app.Span().Text("Home of a user to deal with"),
				),
				app.Span().Body(
					app.Button().Class("btn btn-primary").Text("Card Button"),
				),
			),
		),
		app.If(h.activeTab == "settings",
			components.Card("Settings", "Settings area", "https://picturepan2.github.io/spectre/img/osx-yosemite.jpg",
				app.Span().Body(
					app.Span().Text("User settings"),
				),
				app.Span().Body(
					app.Button().Class("btn btn-primary").Text("Card Button"),
				),
			),
		),
		app.If(h.activeTab == "change-avatar",
			components.Card("Avatar Page", "User avatar page", "https://picturepan2.github.io/spectre/img/osx-yosemite.jpg",
				app.Span().Body(
					app.Span().Text("Change your avatar"),
				),
				app.Span().Body(
					app.Button().Class("btn btn-primary").Text("Card Button"),
				),
			),
		),
	)
}
```


