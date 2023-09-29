---
roles: [internal]
group: components
---
# Modals

Modals are flexible dialog prompts

You can find the original Spectre [documentation here](https://picturepan2.github.io/spectre/components/modals.html)

Create a page structure object to hold the state of the modal

```go
// ModalPage
type ModalPage struct {
  app.Compo
  largeModalOpen bool
  mediumModalOpen bool
  smallModalOpen bool
}
```

Create a button as the action to open a modal, in this case a 'large modal'

```go
app.Button().Class("btn").Text("Large Modal").OnClick(func(ctx app.Context, e app.Event) {
  h.largeModalOpen = true
  h.Update()
  }).Text("click me to open an example large modal"),
```

Create a modal object

```go
lM := components.Modal{
  Size:    "modal-lg", //class from spectre documentation to define size
  ID:      "modal-id",
  Title:   "A large modal",
  Body:    []app.UI{
    	app.P().Text("some modal content"),
  },
  Footer:  []app.UI{app.Span().Text("large modal footer text")},
  OnClose: nil,
}
```

Specify the `OnClose` of the modal. Calling `h.Update()` will update the UI hiding the modal.

```go
lM.OnClose = func() {
  h.largeModalOpen = false
  h.Update()
}
```

Finally when rendering the modal on the page it should appear on

```go
return app.Div().Class("container").Body(
  app.If(h.largeModalOpen,
    &lM),
  ...
```

That way when the `OnClose` is fired, or the button is clicked and the UI is updated the modal will appear based on the state of the boolean value
