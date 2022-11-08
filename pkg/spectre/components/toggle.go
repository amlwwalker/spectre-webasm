package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Toggle struct {
	app.Compo
	Select bool
	Inline bool
	Description string
	OnClick func(ctx app.Context, e app.Event)
}

func (h *Toggle) Render() app.UI {
	checkbox := app.Input().Type("checkbox").OnClick(h.OnClick)
	if h.Select {
		checkbox.Checked(true)
	} else {
		checkbox.Checked(false)
	}
	label := app.Label().Class("form-checkbox").Body(
		checkbox,
		app.I().Class("form-icon"),
		app.Span().Text(h.Description),
	)
	if h.Inline {
		label.Class("form-inline")
	}
	return app.Label().Class("form-switch").Body(
		label,
	)
}
