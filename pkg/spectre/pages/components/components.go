package components

import (
	"fmt"
	"github.com/amlwwalker/spectre-webasm/pkg/markdown"
	"github.com/amlwwalker/spectre-webasm/pkg/spectre/components"
	"github.com/amlwwalker/spectre-webasm/pkg/spectre/layouts"
	pages "github.com/amlwwalker/spectre-webasm/pkg/spectre/pages"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// AvatarPage
type AvatarPage struct {
	app.Compo
}

func (h *AvatarPage) Render() app.UI {
	return pages.Page("Avatars", "Avatars are personal icons next to details or people", "web/documents/avatars.md",
		layouts.TwoColumn(
			components.Avatar("avatar-xl", "https://picturepan2.github.io/spectre/img/avatar-1.png", "AW", "#5755d9"),
			components.Avatar("avatar-lg", "", "AW", "#5755d9"),
			components.Avatar("avatar-sm", "https://picturepan2.github.io/spectre/img/avatar-1.png", "AW", "#5755d9"),
			components.Avatar("avatar-xs", "", "AW", "#5755d9"),
		),
	)
}

// CardPage
type CardPage struct {
	app.Compo
}

func (h *CardPage) Render() app.UI {
	return pages.Page("Cards", "Cards hold succinct information related to A specific topic", "web/documents/cards.md",
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
}

// BadgePage
type BadgePage struct {
	app.Compo
}

func (h *BadgePage) Render() app.UI {
	return pages.Page("Badges", "Badges are often used as unread number indicators", "/web/documents/badges.md",
		components.NotificationBadge("notification", 7),
		components.ButtonBadge("button", 7, nil),
		components.FigureBadge("https://picturepan2.github.io/spectre/img/avatar-3.png", 7, nil),
	)
}

// AccordionPage
type AccordionPage struct {
	app.Compo
}

func (h *AccordionPage) Render() app.UI {
	return pages.Page("Accordions", "Accordions are used to toggle sections of content", "/web/documents/accordions.md",
		components.Accordion(
			components.AccordionMenu("getting-started", "Getting Started", []string{"Installation", "Custom Version", "Browser-Support"}),
			components.AccordionMenu("elements", "Elements", []string{"Typography", "Tables", "Buttons", "Forms", "Icons.css", "Labels", "Code", "Media"}),
			components.AccordionMenu("layouts", "Layouts", []string{"Flexbox grid", "Responsive", "Hero", "Navbar"}),
			components.AccordionMenu("components", "Components",
				[]string{
					"Accordions",
					"Avatars",
					"Badges",
					"Bars",
					"Breadcrumbs",
					"Cards",
					"Chips",
					"Empty States",
					"Menu",
					"Modals",
					"Nav",
					"Pagination",
					"Panels",
					"Popovers",
					"Steps",
					"Tabs",
					"Tiles",
					"Toasts",
					"Tooltips",
				}),
		),
	)
}

// ModalPage
type ModalPage struct {
	app.Compo
	largeModalOpen  bool
	mediumModalOpen bool
	smallModalOpen  bool
}

// gettingStarted render method
func (h *ModalPage) Render() app.UI {
	lM := components.Modal{
		Size:  "modal-lg",
		ID:    "modal-id",
		Title: "A large modal",
		Body: []app.UI{app.Raw(
			markdown.InlineMarkdown(
				`This modal represents A large modal defined by the size attribute 
<code>modal-lg</code>
`))},
		Footer:  []app.UI{app.Span().Text("large modal footer text")},
		OnClose: nil,
	}
	lM.OnClose = func() {
		h.largeModalOpen = false
		h.Update()
	}
	mM := components.Modal{
		ID:      "modal-id",
		Title:   "A medium modal",
		Body:    []app.UI{app.Span().Text("This modal represents A medium modal defined by applying no size attribute")},
		Footer:  []app.UI{app.Span().Text("medium modal footer text")},
		OnClose: nil,
	}
	mM.OnClose = func() {
		h.mediumModalOpen = false
		h.Update()
	}
	sM := components.Modal{
		Size:  "modal-sm",
		ID:    "modal-id",
		Title: "A small modal",
		Body: []app.UI{app.Raw(
			markdown.InlineMarkdown(
				`This modal represents A large modal defined by the size attribute 
<code>modal-sm</code>
`))}, Footer: []app.UI{app.Span().Text("small modal footer text")},
		OnClose: nil,
	}
	sM.OnClose = func() {
		h.smallModalOpen = false
		h.Update()
	}
	return app.Div().Class("container").Body(
		app.If(h.largeModalOpen,
			&lM),
		app.If(h.mediumModalOpen,
			&mM),
		app.If(h.smallModalOpen,
			&sM),
		layouts.NavBar(),
		layouts.SideBar(
			components.Accordion(
				components.AccordionMenu("getting-started", "Getting Started", []string{"Installation", "Custom Version", "Browser-Support"}),
				components.AccordionMenu("elements", "Elements", []string{"Typography", "Tables", "Buttons", "Forms", "Icons.css", "Labels", "Code", "Media"}),
				components.AccordionMenu("layouts", "Layouts", []string{"Flexbox grid", "Responsive", "Hero", "Navbar"}),
				components.AccordionMenu("components", "Components", []string{"Accordions", "Avatars", "Badges", "Bars", "Breadcrumbs", "Cards", "Chips", "Empty States", "Menu", "Modals", "Nav", "Pagination", "Panels", "Popovers", "Steps", "Tabs", "Tiles", "Toasts", "Tooltips"}),
			), app.Div().Class("docs-content").Class("content").Body(
				layouts.Hero("Modals", "Using Modals (webasm)"),

				app.Button().Class("btn").Text("Large Modal").OnClick(func(ctx app.Context, e app.Event) {
					h.largeModalOpen = true
					h.Update()
				}).Text("click me to open an example large modal"),
				app.Button().Class("btn").Text("Medium Modal").OnClick(func(ctx app.Context, e app.Event) {
					h.mediumModalOpen = true
					h.Update()
				}).Text("click me to open an example medium modal"),
				app.Button().Class("btn").Text("Small Modal").OnClick(func(ctx app.Context, e app.Event) {
					h.smallModalOpen = true
					h.Update()
				}).Text("click me to open an example small modal"),
				markdown.NewRemoteMarkdownDoc().Src("/web/documents/modals.md"),
			),
		),
	)
}

type TilePage struct {
	app.Compo
}

func (h *TilePage) Render() app.UI {
	return pages.Page("Tiles", "Tiles are repeatable or embeddable information blocks.", "/web/documents/tiles.md",
		app.Div().Class("columns").Body(
			app.Div().Class("column col-9 col-sm-12").Body(
				components.Tile("The Avengers", "Earth's Mightiest Heroes joined forces to take on threats that were too big for any one hero to tackle...", "Join", "https://picturepan2.github.io/spectre/img/avatar-1.png"),
				components.Tile("The Avengers", "Earth's Mightiest Heroes joined forces to take on threats that were too big for any one hero to tackle...", "Join", "https://picturepan2.github.io/spectre/img/avatar-1.png"),
			),
		),
	)
}

type TabsPage struct {
	app.Compo
	activeTab string
}

func (h *TabsPage) OnMount(ctx app.Context) {
	h.activeTab = "profile-tab"
}
func (h *TabsPage) Render() app.UI {
	fmt.Println("h.activeTab", h.activeTab)
	return pages.Page("Tabs", "Tabs group information together onto different tabs based on content or context", "web/documents/tabs.md",
		app.Div().Class("column col-12").Body(
			components.Tabs(
				components.Tab("Profile", h.activeTab == "profile-tab", func(ctx app.Context, e app.Event) {
					h.activeTab = "profile-tab"
					h.Update()
				}),
				components.Tab("Files", h.activeTab == "files-tab", func(ctx app.Context, e app.Event) {
					h.activeTab = "files-tab"
					h.Update()
				}),
				components.Tab("Tasks", h.activeTab == "tasks-tab", func(ctx app.Context, e app.Event) {
					h.activeTab = "tasks-tab"
					h.Update()
				}),
			),
			app.If(h.activeTab == "profile-tab",
				components.Card("Profile", "Profile of user", "https://picturepan2.github.io/spectre/img/osx-yosemite.jpg",
					app.Span().Body(
						app.Span().Text("Profile of a user to deal with"),
					),
					app.Span().Body(
						app.Button().Class("btn btn-primary").Text("Card Button"),
					),
				),
			),
			app.If(h.activeTab == "files-tab",
				components.Card("Files", "Files that need handling", "https://picturepan2.github.io/spectre/img/osx-yosemite.jpg",
					app.Span().Body(
						app.Span().Text("Files that the user hasn't processed yet"),
					),
					app.Span().Body(
						app.Button().Class("btn btn-primary").Text("Card Button"),
					),
				),
			),
			app.If(h.activeTab == "tasks-tab",
				components.Card("Tasks", "Tasks that need doing", "https://picturepan2.github.io/spectre/img/osx-yosemite.jpg",
					app.Span().Body(
						app.Span().Text("Tasks the user is yet to deal with"),
					),
					app.Span().Body(
						app.Button().Class("btn btn-primary").Text("Card Button"),
					),
				),
			),
		),
	)
}

type PanelPage struct {
	app.Compo
}

func (h *PanelPage) Render() app.UI {
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
	return pages.Page("Panels", "Panels are flexible view container with auto-expand content section.", "/web/documents/panels.md",
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

type ChipsPage struct {
	app.Compo
}

func (h *ChipsPage) Render() app.UI {
	return pages.Page("Chips", "Chips are complex entities in small blocks.", "/web/documents/chips.md",
		layouts.FlexBox("80%",
			components.Chip("Crime", "", "", false),
			components.Chip("Crime", "avatar-sm", "https://picturepan2.github.io/spectre/img/avatar-1.png", false),
			components.Chip("Crime", "avatar-sm", "https://picturepan2.github.io/spectre/img/avatar-1.png", true),
		),
	)
}

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

type EmptyStatePage struct {
	app.Compo
}

func (h *EmptyStatePage) Render() app.UI {
	return pages.Page("Empty States", "Empty states/blank slates are commonly used as placeholders for first time use, empty data and error screens.", "/web/documents/empty.md",
		app.Br(),
		components.EmptyState("icon-people", "An Empty State", "An example empty state", "btn-primary", "does nothing", nil),
		app.Br(),
		components.EmptyState("icon-3x icon-mail", "You have no new messages", "Click the button to start a conversation", "btn-primary", "Send a message", nil),
	)
}
