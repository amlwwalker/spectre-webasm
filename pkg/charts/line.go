package charts

import (
"encoding/json"
"github.com/go-echarts/go-echarts/v2/charts"
"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Line struct {
	app.Compo
	Class           string
	Options         *charts.Line
	//Opt *charts.Graph
	eChartsInstance app.Value
}


// must inject the script tag "https://cdnjs.cloudflare.com/ajax/libs/echarts/5.4.0/echarts.js"
func (c *Line) OnMount(ctx app.Context) {
	ctx.Defer(func(context app.Context) {
		c.eChartsInstance = app.Window().Get("echarts").
			Call("init", c.JSValue(), c.Options.Theme)
		c.UpdateConfig(context, c.Options)
	})
}

func (c *Line) OnDismount() {
	if c.eChartsInstance != nil {
		c.eChartsInstance.Call("dispose")
	}
}

func (c *Line) UpdateConfig(ctx app.Context, config *charts.Line) {
	config.Validate()
	c.Options = config

	if c.eChartsInstance != nil {
		c.eChartsInstance.Call("dispose")
	}
	c.eChartsInstance = app.Window().Get("echarts").
		Call("init", c.JSValue(), c.Options.Theme)

	ctx.Async(func() {
		jsonString, _ := json.Marshal(c.Options.JSON())
		options := app.Window().Get("JSON").Call("parse", string(jsonString))
		c.eChartsInstance.Call("setOption", options)
		c.Update()
	})
}

func (c *Line) Render() app.UI {
	if c.Options == nil {
		c.Options = charts.NewLine()
		c.Options.Validate()
	}
	return app.Div().Class(c.Class).ID(c.Options.ID).
		Style("width", c.Options.Initialization.Width).//c.Options.Initialization.Width).
		Style("height", c.Options.Initialization.Height)
}

