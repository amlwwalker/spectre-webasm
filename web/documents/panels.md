# Panels

Panels are useful when you need content to be grouped and auto expanded, for instance a list or any other grouped content

You can find the original Spectre [documentation here](https://picturepan2.github.io/spectre/components/panels.html)

### A basic panel

```go
layouts.Columns(
  components.Panel(app.Span().Text("a panel"), app.Span().Text("some panel subtitle"), app.Span().Text("navigation area"), app.Span().Text("some body content"), app.Span().Text("footer content")).Style("max-height", "40vh"),
  components.Panel(app.Span().Text("a panel"), app.Span().Text("some panel subtitle"), app.Span().Text("navigation area"), app.Span().Text("some body content"), app.Span().Text("footer content")).Style("max-height", "40vh"),
)
```
_Note this shows how you can bolt styles on wherever you like_

## Panels with layouts and more complex content

The following demonstrates constructing more complex panels made up of Tabs and Tiles. This is the example shown above.

```go
	avengersTiles := []app.UI {
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

	bannerProfile := []app.UI {
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
```
...
```go
layouts.Columns(
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
    components.ButtonBlock("Save", nil),
  ).Style("max-height", "40vh"),
),
```
