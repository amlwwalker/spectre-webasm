package components

import (
	"encoding/base64"
	"encoding/json"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type ReactPlaceholder struct {
	app.Compo

	ComponentName string
	Props         map[string]interface{}
}

func (p *ReactPlaceholder) OnMount(ctx app.Context) {
	// Use JS to invoke the JSX mounting code.
	// This will run every time this component is mounted to the DOM.
	app.Window().Call("mountReactComponents")
}
func (p *ReactPlaceholder) Render() app.UI {
	// We'll use data attributes to pass information to our React entry script.
	// This helps determine which React component to mount and any props/data it needs.
	propsJSON, err := json.Marshal(p.Props)
	if err != nil {
		// handle the error, maybe just set propsJSON to "{}"
		propsJSON = []byte("{}")
	}
	encodedProps := base64.StdEncoding.EncodeToString(propsJSON)

	return app.Div().DataSet("react-component", p.ComponentName).DataSet("react-props", encodedProps)
}
