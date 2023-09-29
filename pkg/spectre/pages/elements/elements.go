package elements

import (
	"github.com/amlwwalker/spectre-webasm/pkg/spectre/components"
	"github.com/amlwwalker/spectre-webasm/pkg/spectre/layouts"
	pages "github.com/amlwwalker/spectre-webasm/pkg/spectre/pages"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

var stateName = "button-click"

// ButtonPage
type ButtonPage struct {
	app.Compo
	stateValue string
}

func (h *ButtonPage) OnMount(ctx app.Context) {
	ctx.ObserveState(stateName).Value(&h.stateValue)
}
func (h *ButtonPage) Render() app.UI {
	onClick := func(ctx app.Context, e app.Event) {
		app.Window().Call("alert", "setting state now...")
		ctx.SetState(stateName, "button was clicked. State was obsererved and monitored.")
	}
	return pages.NewPage("Buttons", "Buttons submit forms, or execute actions", "/documents/buttons.md",
		layouts.FlexBox("75%",
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
				app.P().Text(h.stateValue),
			),
		),
	)
}
