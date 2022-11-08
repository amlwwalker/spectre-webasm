package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Checkbox struct {
	app.Compo
	selected bool
	Inline bool
	Indeterminate bool
	ID string
	Description string
	OnClick func(ctx app.Context, e app.Event)
}

func NewInlineCheckbox(description, id string, selected bool, onClick func(ctx app.Context, e app.Event)) app.UI {
	t := Checkbox{
		selected:      selected,
		Inline:      true,
		ID: id,
		Description: description,
	}
	t.OnClick = func(ctx app.Context, e app.Event) {
		t.selected = !t.selected
		onClick(ctx, e)
	}
	return app.UI(&t)
}

func (h *Checkbox) Render() app.UI {
	checkbox := app.Input().Type("checkbox").ID(h.ID).Attr("indeterminate", h.Indeterminate).Checked(h.selected).OnClick(h.OnClick)
	label := app.Label().Class("form-checkbox").Body(
		checkbox,
		app.I().Class("form-icon"),
		app.Span().Text(h.Description),
	)
	if h.Inline {
		label.Class("form-inline")
	}
	return label
}
