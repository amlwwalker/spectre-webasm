package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func Input(description, id, placeholder, value string, onChange func(ctx app.Context, e app.Event)) app.UI {
	return app.Div().Class("form-group").Body(
		app.Label().Class("form-label").For(id).Text(description),
		app.Input().Class("form-input").Type("text").ID(id).Placeholder(placeholder).Value(value).OnInput(func(ctx app.Context, e app.Event) {
			onChange(ctx, e)
		}),
	)
}

func TextArea(description, id, placeholder, value string, rows int, onChange func(ctx app.Context, e app.Event)) app.UI {
	return app.Div().Class("form-group").Body(
		app.Label().Class("form-label").For(id).Text(description),
		app.Textarea().Class("form-input").ID(id).Placeholder(placeholder).Text(value).Rows(rows).OnInput(func(ctx app.Context, e app.Event) {
			onChange(ctx, e)
		}),
	)
}

func FormSelect(options []app.HTMLOption) app.UI {
	return app.Div().Class("form-group").Body(
		app.Select().Class("form-select").Body(
			app.Range(options).Slice(func(i int) app.UI {
				return options[i]
			}),
		),
	)
}
