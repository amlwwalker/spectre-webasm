# Buttons

Buttons include simple button styles for actions in different types and sizes.

You can find the original Spectre [documentation here](https://picturepan2.github.io/spectre/elements/buttons.html)

**Note** with most elements including buttons you can add styles after the fact. I.e

```go
layouts.FlexBox("100%",
  app.Div().Class("column").Body(components.Button("block button", "btn btn-block", nil)),
  app.Div().Class("column").Body(components.Button("primary button", "btn  btn-primary", nil)),
  app.Div().Class("column").Body(components.Button("link button", "btn  btn-link", nil)),
  app.Div().Class("column").Body(components.Button("succes button", "btn  btn-success", nil)),
  app.Div().Class("column").Body(components.Button("error button", "btn  btn-error", nil)),
  app.Div().Class("column").Body(components.Button("large error button", "btn  btn-error btn-lg", nil)),
  app.Div().Class("column").Body(components.Button("small success button", "btn  btn-success btn-sm", nil)),
)
```

## State

Although off topic this is a good opportunity to mention state. State can be set based off the context off an `onClick`. State allows data to be stored and picked up elsewhere in the application. For instance when you click one of the buttons on this page, text will appear in the paragraph object underneath.

The onClick routine attached to each of the buttons looks like

```go
	onClick := func(ctx app.Context, e app.Event) {
		app.Window().Call("alert", "setting state now...")
		ctx.SetState(stateName, "button was clicked. State was obsererved and monitored.")
	}
```

To make this 'observable', the paragraph has been set to display the text stored on the page object.

Then the pageObject has been configured to 'observe' and update the value when the state changes

```go
var stateName = "button-click"
// ButtonPage
type ButtonPage struct {
	app.Compo
	stateValue string
}

func (h *ButtonPage) OnMount(ctx app.Context) {
	ctx.ObserveState(stateName).Value(&h.stateValue)
}
```

and 

```go
app.P().Text(h.stateValue)
```

When a button is clicked the state is updated which in turn updates the `ButtonPage member variable and the paragraph updates to the new text.`
