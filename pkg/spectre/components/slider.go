package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"strconv"
)

type Slider struct {
	app.Compo
	Description string
	Min, Max, Value int
	OnInput func(ctx app.Context, e app.Event)
}

func NewSlider(description string, onClick func(ctx app.Context, e app.Event)) app.UI {
	slider := Slider{
		Description: description,
		Min:         1,
		Max:         10,
		Value:       1,
	}
	slider.OnInput = func(ctx app.Context, e app.Event) {
		jsSrc := ctx.JSSrc()
		value := jsSrc.Get("value").String()
		jsSrc.Call("setAttribute", "value", value) //set the value back on the DOM object
		slider.Value, _ = strconv.Atoi(value) //retrieve the value so that we have a reference to it
		onClick(ctx, e)
	}
	return app.UI(&slider)
}

func (h *Slider) Render() app.UI {
	slider := app.Input().Class("slider").Type("range")
	slider.Attr("min", h.Min).Attr("max", h.Max).Attr("value", h.Value)
	slider.Class("tooltip")
	slider.OnInput(h.OnInput)

	label := app.Label().Body(
		slider,
		app.I().Class("form-icon"),
		app.Span().Text(h.Description))
	return label
}
