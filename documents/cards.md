---
roles: [maintainer, internal, user]
group: components
---
# Cards

Cards are flexible content containers.

You can find the original Spectre  [documentation here](https://picturepan2.github.io/spectre/components/cards.html)

You can use cards within layouts however to define a card

Using a card is quite simple

```go
components.Card("Apple", "Apple desktop background", "https://picturepan2.github.io/spectre/img/osx-yosemite.jpg",
  app.Span().Body(
    app.Span().Text("To make a contribution to the world by making tools for the mind that advance humankind."),
  ),
  app.Span().Body(
    app.Button().Class("btn btn-primary").Text("Card Button"),
  ),
)
```

The card component takes 

1. A title
2. A subtitle
3. A src for the image
4. The body content of the card
5. The footer content of the card


## Layouts
Cards will take up the available space, however they can be placed inside layouts. In the example above, they are laid out like

```go
layouts.Columns(
  components.Card("Apple", "Apple desktop background", "https://picturepan2.github.io/spectre/img/osx-yosemite.jpg",
    app.Span().Body(
      app.Span().Text("To make a contribution to the world by making tools for the mind that advance humankind."),
    ),
    app.Span().Body(
      app.Button().Class("btn btn-primary").Text("Card Button"),
    ),
  ),
  components.Card("Microsoft", "Microsoft desktop background", "https://picturepan2.github.io/spectre/img/osx-el-capitan.jpg",
    app.Span().Body(
      app.Span().Text("Empower every person and every organization on the planet to achieve more."),
    ),
    app.Span().Body(
      app.Button().Class("btn btn-primary").Text("Card Button"),
    ),
  ),
)
```

See layouts/columns for more information on this specific layout
